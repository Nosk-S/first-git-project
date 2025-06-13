package handler

import (
	_ "database/sql"
	"gin-project/web"
	"text/template"

	"github.com/gin-gonic/gin"
)

//var db *sql.DB

type PublicHandler struct {
}

func NewPublicHandler() *PublicHandler {
	return &PublicHandler{}
}

func (h *PublicHandler) Index(c *gin.Context) {
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
}

func (h *PublicHandler) Admin(c *gin.Context) {
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

}

// func (h *PublicHandler) Getcards(c *gin.Context) {
// 	c.JSON(200, cards)
// }
