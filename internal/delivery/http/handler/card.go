package handler

import (
	"gin-project/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterCardRoutes(rg *gin.RouterGroup, uc usecase.CardUsecase) {
	rg.GET("/cards", getAllCards(uc))
}

func getAllCards(uc usecase.CardUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		cards, err := uc.GetCards(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, cards)
	}
}
