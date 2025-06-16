package handler

import (
	"gin-project/internal/domain"
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

	getCardsRequest := new(domain.Card)

	if err := c.ShouldBindJSON(&getCardsRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error parsing request body"})
		panic(err)
	}

	cards, err := h.CardUsecase.GetCards(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error GetCard"})
		panic(err)
	}

	c.JSON(200, cards)

}
