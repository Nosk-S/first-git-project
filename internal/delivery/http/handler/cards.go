package handler

import (
	"gin-project/internal/models"
	"gin-project/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CardsHandler struct {
	CardUsecase usecase.CardUsecase
}

func NewCardsHandler(cardUsecase usecase.CardUsecase) *CardsHandler {
	return &CardsHandler{
		CardUsecase: cardUsecase,
	}
}

func (h *CardsHandler) GetCards(c *gin.Context) {

	cards, err := h.CardUsecase.GetCards(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, cards)

}

func (h *CardsHandler) GetCard(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"erreur": "id requis"})
		return
	}

	card, err := h.CardUsecase.GetCard(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erreur": err.Error()})
		return
	}

	c.JSON(200, card)

}

func (h *CardsHandler) CardWhere(c *gin.Context) {

	var research models.CardResearch

	if err := c.ShouldBindJSON(&research); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erreur": "JSON invalide ou mal formé"})
		return
	}

	if research.Research == "" {
		c.JSON(http.StatusBadRequest, gin.H{"erreur": "Le champ 'research' est requis"})
		return
	}

	cards, err := h.CardUsecase.CardWhere(c.Request.Context(), research)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erreur": err.Error()})
		return
	}

	c.JSON(200, cards)

}
