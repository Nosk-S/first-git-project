package main

import (
	"database/sql"
	"fmt"
	"gin-project/internal/domain"
	"gin-project/internal/models"
	"gin-project/web"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {

	// env config
	envConfig, err := godotenv.Read(".env")
	if err != nil {
		panic(err)
	}

	// db config
	cfg := mysql.NewConfig()
	cfg.User = envConfig["DBUSER"]
	cfg.Passwd = envConfig["DBPASS"]
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "gin-project"

	// Get a database handle.
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	router := gin.Default()

	public := router.Group("/")
	{
		public.GET("/", func(c *gin.Context) {
			files := []string{
				"templates/index.html",
			}

			tmpl, err := template.ParseFS(web.GetEmbeddedFiles(), files...)
			if err != nil {
				panic(err)
			}

			data := map[string]string{
				"title": "Page d'accueil",
			}

			err = tmpl.Execute(c.Writer, data)
			if err != nil {
				panic(err)
			}
		})
		public.GET("/static/*filepath", gin.WrapH(http.FileServer(http.FS(web.GetEmbeddedFiles()))))
		admin := public.Group("/admin")
		{
			admin.Use(gin.BasicAuth(gin.Accounts{
				"admin": "admin",
			}))
			admin.GET("/", func(c *gin.Context) {
				files := []string{
					"templates/admin/index.html",
				}

				tmpl, err := template.ParseFS(web.GetEmbeddedFiles(), files...)
				if err != nil {
					panic(err)
				}

				data := map[string]string{
					"title": "Page réservé à l'admin",
				}

				err = tmpl.Execute(c.Writer, data)
				if err != nil {
					panic(err)
				}
			})
		}
		api := public.Group("/api")
		{
			api.GET("/card", func(c *gin.Context) {
				var crt domain.Card
				id := c.Query("id")
				err = db.QueryRow("SELECT id, nom, mana, effets FROM cartehs WHERE id = ?", id).Scan(&crt.ID, &crt.Name, &crt.Mana, &crt.Effects)
				if err != nil {
					panic(err.Error()) // proper error handling instead of panic in your app
				}
				c.JSON(200, crt)
			})

			api.GET("/cards", func(c *gin.Context) {
				rows, err := db.Query("SELECT id, nom, mana, effets FROM cartehs ")
				if err != nil {
					panic(err)
				}

				defer rows.Close()

				var cards []domain.Card
				for rows.Next() {
					var crt domain.Card

					if err := rows.Scan(&crt.ID, &crt.Name, &crt.Mana, &crt.Effects); err != nil {
						panic(err)
					}

					cards = append(cards, crt)
				}

				c.JSON(200, cards)
			})

			api.POST("/research", func(c *gin.Context) {
				var crt models.CardResearch

				// 1) Binder le JSON dans crt
				if err := c.ShouldBindJSON(&crt); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"erreur": "JSON invalide ou mal formé"})
					return
				}

				// 2) Vérifier que Research n'est pas vide
				if strings.TrimSpace(crt.Research) == "" {
					c.JSON(http.StatusBadRequest, gin.H{"erreur": "Le champ 'research' est requis pour rechercher une carte."})
					return
				}

				// 3) Faire le SELECT correctement
				rows, err := db.Query("SELECT id, nom, mana, effets FROM `cartehs` WHERE `nom` LIKE ? OR `effets` LIKE ?", "%"+crt.Research+"%", "%"+crt.Research+"%")

				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"erreur": "Échec de la recherche", "details": err.Error()})
					return
				}
				defer rows.Close()

				// 4) Parcourir les résultats
				var cartes []domain.Card
				for rows.Next() {
					var carte domain.Card
					if err := rows.Scan(&carte.ID, &carte.Name, &carte.Mana, &carte.Effects); err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"erreur": "Erreur de lecture des données", "details": err.Error()})
						return
					}
					cartes = append(cartes, carte)
				}

				c.JSON(http.StatusOK, gin.H{
					"cartes": cartes,
				})
			})

			api.Use(gin.BasicAuth(gin.Accounts{
				"admin": "admin",
			}))
			api.POST("/card", func(c *gin.Context) {
				var crt models.CardInsertRequest

				// DEBUG
				// jsonData, err := io.ReadAll(c.Request.Body)
				// if err != nil {
				// 	// Handle error
				// }
				// fmt.Println(string(jsonData))

				if err := c.BindJSON(&crt); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"erreur": "Format JSON invalide"})
					return
				}

				if crt.Name == nil {
					c.JSON(http.StatusBadRequest, gin.H{"erreur": "Le champ 'Name' est requis pour ajouter une carte."})
				}
				if crt.Mana == nil {
					c.JSON(http.StatusBadRequest, gin.H{"erreur": "Le champ 'Mana' est requis pour ajouter une carte."})
				}
				if crt.Effects == nil {
					c.JSON(http.StatusBadRequest, gin.H{"erreur": "Le champ 'Effects' est requis pour ajouter une carte."})
				}

				res, err := db.Exec("INSERT INTO cartehs VALUES (?, ?, ?, ?)", crt.ID, crt.Name, crt.Mana, crt.Effects)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"erreur": "Échec d'ajout", "details": err.Error()})
					return
				}

				rowsAffected, _ := res.RowsAffected()

				c.JSON(http.StatusOK, gin.H{
					"message":         "Carte ajouter avec succès",
					"carte":           crt,
					"lignesModifiees": rowsAffected,
				})
			})
			api.DELETE("/card", func(c *gin.Context) {
				var crt models.CardDeleteRequest

				if err := c.BindJSON(&crt); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"erreur": "Format JSON invalide"})
					return
				}

				if crt.ID == nil {
					c.JSON(http.StatusBadRequest, gin.H{"erreur": "Le champ 'id' est requis pour supprimer une carte."})
					return
				}

				res, err := db.Exec("DELETE FROM cartehs WHERE id = ?", crt.ID)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"erreur": "Échec de mise à jour", "details": err.Error()})
					return
				}

				rowsAffected, _ := res.RowsAffected()

				c.JSON(http.StatusOK, gin.H{
					"message":         "Carte supprimer avec succès",
					"lignesModifiees": rowsAffected,
				})
			})
			api.PATCH("/card", func(c *gin.Context) {

				var crt models.CardUpdateRequest

				if err := c.BindJSON(&crt); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"erreur": "Format JSON invalide"})
					return
				}

				if crt.ID == nil {
					c.JSON(http.StatusBadRequest, gin.H{"erreur": "Le champ 'id' est requis pour mettre à jour une carte."})
					return
				}

				var current models.CardUpdateRequest
				row := db.QueryRow("SELECT id, nom, mana, effets FROM cartehs WHERE id = ?", crt.ID)
				var nom string
				var mana int8
				var effets string
				if err := row.Scan(&current.ID, &nom, &mana, &effets); err != nil {
					c.JSON(http.StatusNotFound, gin.H{"erreur": "Carte non trouvée avec cet ID"})
					return
				}

				current.Name = &nom
				current.Mana = &mana
				current.Effects = &effets

				if crt.Name == nil {
					crt.Name = current.Name
				}
				if crt.Mana == nil {
					crt.Mana = current.Mana
				}
				if crt.Effects == nil {
					crt.Effects = current.Effects
				}

				res, err := db.Exec("UPDATE cartehs SET nom = ?, mana = ?, effets = ? WHERE id = ?", crt.Name, crt.Mana, crt.Effects, crt.ID)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"erreur": "Échec de mise à jour", "details": err.Error()})
					return
				}

				rowsAffected, _ := res.RowsAffected()

				c.JSON(http.StatusOK, gin.H{
					"message":         "Carte mise à jour avec succès",
					"carte":           crt,
					"lignesModifiees": rowsAffected,
				})
			})
		}
	}

	router.Run(":8080")

}
