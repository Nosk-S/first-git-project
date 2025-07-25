package usecase

import (
	"context"
	"fmt"
	"gin-project/internal/models"
	"gin-project/internal/repository"
)

type CardUsecase interface {
	GetCards(ctx context.Context) ([]models.CardSelect, error)
	GetCard(ctx context.Context, id string) (models.CardSelect, error)
	CardWhere(ctx context.Context, research models.CardResearch) ([]models.CardSelect, error)
	InsertCard(ctx context.Context, newcard models.CardInsertRequest) (models.CardInsertRequest, error)
	DeleteCard(ctx context.Context, deletecard models.CardDeleteRequest) (models.CardDeleteRequest, error)
	UpdateCard(ctx context.Context, updatecard models.CardUpdateRequest) (models.CardUpdateRequest, error)
}

type cardUsecase struct {
	repo repository.CardRepository
}

func NewCardUsecase(repo repository.CardRepository) CardUsecase {
	return &cardUsecase{repo: repo}
}

func (u *cardUsecase) GetCards(ctx context.Context) ([]models.CardSelect, error) {
	return u.repo.SelectCards(ctx)
}

func (u *cardUsecase) GetCard(ctx context.Context, id string) (models.CardSelect, error) {
	return u.repo.SelectCard(ctx, id)
}

func (u *cardUsecase) CardWhere(ctx context.Context, research models.CardResearch) ([]models.CardSelect, error) {
	return u.repo.ResearchCards(ctx, research)
}

func (u *cardUsecase) InsertCard(ctx context.Context, newcard models.CardInsertRequest) (models.CardInsertRequest, error) {

	if newcard.Name == nil {
		return models.CardInsertRequest{}, fmt.Errorf("le champ Name est requis")
	}
	if newcard.Mana == nil {
		return models.CardInsertRequest{}, fmt.Errorf("le champ Mana est requis")
	}
	if newcard.Effects == nil {
		return models.CardInsertRequest{}, fmt.Errorf("le champ Effects est requis")
	}

	return u.repo.AddCard(ctx, newcard)

}

func (u *cardUsecase) DeleteCard(ctx context.Context, deletecard models.CardDeleteRequest) (models.CardDeleteRequest, error) {

	if deletecard.ID == nil {
		return models.CardDeleteRequest{}, fmt.Errorf("le champ ID est requis")
	}

	currentID := fmt.Sprint(*deletecard.ID)

	if _, err := u.repo.SelectCard(ctx, currentID); err != nil {
		return models.CardDeleteRequest{}, err
	}

	return u.repo.DeleteCard(ctx, deletecard)
}

func (u *cardUsecase) UpdateCard(ctx context.Context, updatecard models.CardUpdateRequest) (models.CardUpdateRequest, error) {
	if updatecard.ID == nil {
		return models.CardUpdateRequest{}, fmt.Errorf("le champ ID est requis")
	}

	if updatecard.Name == nil && updatecard.Mana == nil && updatecard.Effects == nil {
		return models.CardUpdateRequest{}, fmt.Errorf(
			"au moins un champ (Name, Mana ou Effects) doit être renseigné pour la mise à jour",
		)
	}

	currentID := fmt.Sprint(*updatecard.ID)
	currentcard, err := u.repo.SelectCard(ctx, currentID)
	if err != nil {
		return models.CardUpdateRequest{}, err
	}

	if updatecard.Name == nil {
		updatecard.Name = &currentcard.Name
	}
	if updatecard.Mana == nil {
		updatecard.Mana = &currentcard.Mana
	}
	if updatecard.Effects == nil {
		updatecard.Effects = &currentcard.Effects
	}

	return u.repo.EditCard(ctx, updatecard)
}
