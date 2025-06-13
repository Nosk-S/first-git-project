package usecase

import (
	"gin-project/internal/domain"
	"gin-project/internal/repository"
)

type CardUsecase interface {
}

type CardUsecaseImpl struct {
	CardRepository repository.CardRepository
}

func NewCardUsecase(cardRepo repository.CardRepository) CardUsecase {
	return &CardUsecaseImpl{
		CardRepository: cardRepo,
	}

}

func (uc *CardUsecaseImpl) SelectCards([]domain.Card) (domain.Card, error) {

	return

}
