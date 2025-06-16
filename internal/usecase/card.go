package usecase

import (
	"context"
	"gin-project/internal/domain"
	"gin-project/internal/repository"
)

type CardUsecase interface {
	GetCards(ctx context.Context) ([]domain.Card, error)
}

type cardUsecase struct {
	repo repository.CardRepository
}

func NewCardUsecase(repo repository.CardRepository) CardUsecase {
	return &cardUsecase{repo: repo}
}

func (u *cardUsecase) GetCards(ctx context.Context) ([]domain.Card, error) {
	return u.repo.SelectCards(ctx)
}
