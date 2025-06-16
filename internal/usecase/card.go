package usecase

import (
	"context"
	"gin-project/internal/domain"
	"gin-project/internal/models"
	"gin-project/internal/repository"
)

type CardUsecase interface {
	GetCards(ctx context.Context) ([]domain.Card, error)
	GetCard(ctx context.Context, id string) (domain.Card, error)
	CardWhere(ctx context.Context, research models.CardResearch) ([]domain.Card, error)
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

func (u *cardUsecase) GetCard(ctx context.Context, id string) (domain.Card, error) {
	return u.repo.SelectCard(ctx, id)
}

func (u *cardUsecase) CardWhere(ctx context.Context, research models.CardResearch) ([]domain.Card, error) {
	return u.repo.ResearchCards(ctx, research)
}
