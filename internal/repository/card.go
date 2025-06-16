package repository

import (
	"context"
	"database/sql"
	"gin-project/internal/domain"
	"gin-project/internal/models"
)

type CardRepository interface {
	SelectCards(ctx context.Context) ([]domain.Card, error)
	SelectCard(ctx context.Context, id string) (domain.Card, error)
	ResearchCards(ctx context.Context, research models.CardResearch) ([]domain.Card, error)
}

type SQLCardRepository struct {
	db *sql.DB
}

func NewSQLCardRepository(db *sql.DB) CardRepository {
	return &SQLCardRepository{db: db}
}

func (r *SQLCardRepository) SelectCards(ctx context.Context) ([]domain.Card, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, nom, mana, effets FROM cartehs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cards := make([]domain.Card, 0)
	for rows.Next() {
		var c domain.Card

		if err := rows.Scan(&c.ID, &c.Name, &c.Mana, &c.Effects); err != nil {
			return nil, err
		}
		cards = append(cards, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}

func (r *SQLCardRepository) SelectCard(ctx context.Context, id string) (domain.Card, error) {
	var card domain.Card
	err := r.db.QueryRow("SELECT id, nom, mana, effets FROM cartehs WHERE id = ?", id).Scan(&card.ID, &card.Name, &card.Mana, &card.Effects)
	if err != nil {
		return domain.Card{}, err
	}

	return card, nil

}

func (r *SQLCardRepository) ResearchCards(ctx context.Context, research models.CardResearch) ([]domain.Card, error) {

	pattern := "%" + research.Research + "%"

	rows, err := r.db.QueryContext(ctx, "SELECT id, nom, mana, effets FROM `cartehs` WHERE `nom` LIKE ? OR `effets` LIKE ?", pattern, pattern)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cards := make([]domain.Card, 0)
	for rows.Next() {
		var c domain.Card

		if err := rows.Scan(&c.ID, &c.Name, &c.Mana, &c.Effects); err != nil {
			return nil, err
		}
		cards = append(cards, c)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cards, nil

}

// api.POST("/research", func(c *gin.Context) {
// 	var crt models.CardResearch

// 	// 1) Binder le JSON dans crt
// 	if err := c.ShouldBindJSON(&crt); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"erreur": "JSON invalide ou mal formé"})
// 		return
// 	}

// 	// 2) Vérifier que Research n'est pas vide
// 	if strings.TrimSpace(crt.Research) == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"erreur": "Le champ 'research' est requis pour rechercher une carte."})
// 		return
// 	}

// 	// 3) Faire le SELECT correctement
// 	rows, err := db.Query("SELECT id, nom, mana, effets FROM `cartehs` WHERE `nom` LIKE ? OR `effets` LIKE ?", "%"+crt.Research+"%", "%"+crt.Research+"%")

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"erreur": "Échec de la recherche", "details": err.Error()})
// 		return
// 	}
// 	defer rows.Close()

// 	// 4) Parcourir les résultats
// 	var cartes []domain.Card
// 	for rows.Next() {
// 		var carte domain.Card
// 		if err := rows.Scan(&carte.ID, &carte.Name, &carte.Mana, &carte.Effects); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"erreur": "Erreur de lecture des données", "details": err.Error()})
// 			return
// 		}
// 		cartes = append(cartes, carte)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"cartes": cartes,
// 	})
// })
