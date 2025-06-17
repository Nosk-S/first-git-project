package repository

import (
	"context"
	"database/sql"
	"gin-project/internal/models"
)

type CardRepository interface {
	SelectCards(ctx context.Context) ([]models.CardSelect, error)
	SelectCard(ctx context.Context, id string) (models.CardSelect, error)
	ResearchCards(ctx context.Context, research models.CardResearch) ([]models.CardSelect, error)
	AddCard(ctx context.Context, newcard models.CardInsertRequest) (models.CardInsertRequest, error)
	DeleteCard(ctx context.Context, deletecard models.CardDeleteRequest) (models.CardDeleteRequest, error)
	EditCard(ctx context.Context, updatecard models.CardUpdateRequest) (models.CardUpdateRequest, error)
}

type SQLCardRepository struct {
	db *sql.DB
}

func NewSQLCardRepository(db *sql.DB) CardRepository {
	return &SQLCardRepository{db: db}
}

func (r *SQLCardRepository) SelectCards(ctx context.Context) ([]models.CardSelect, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, nom, mana, effets FROM cartehs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cards := make([]models.CardSelect, 0)
	for rows.Next() {
		var c models.CardSelect

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

func (r *SQLCardRepository) SelectCard(ctx context.Context, id string) (models.CardSelect, error) {
	var card models.CardSelect
	err := r.db.QueryRow("SELECT id, nom, mana, effets FROM cartehs WHERE id = ?", id).Scan(&card.ID, &card.Name, &card.Mana, &card.Effects)
	if err != nil {
		return models.CardSelect{}, err
	}

	return card, nil

}

func (r *SQLCardRepository) ResearchCards(ctx context.Context, research models.CardResearch) ([]models.CardSelect, error) {

	pattern := "%" + research.Research + "%"

	rows, err := r.db.QueryContext(ctx, "SELECT id, nom, mana, effets FROM `cartehs` WHERE `nom` LIKE ? OR `effets` LIKE ?", pattern, pattern)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cards := make([]models.CardSelect, 0)
	for rows.Next() {
		var c models.CardSelect

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

func (r *SQLCardRepository) AddCard(ctx context.Context, newcard models.CardInsertRequest) (models.CardInsertRequest, error) {

	if _, err := r.db.ExecContext(
		ctx,
		"INSERT INTO cartehs (nom, mana, effets) VALUES (?, ?, ?)",
		*newcard.Name,
		*newcard.Mana,
		*newcard.Effects,
	); err != nil {
		return models.CardInsertRequest{}, err
	}

	return newcard, nil
}

func (r *SQLCardRepository) DeleteCard(ctx context.Context, deletecard models.CardDeleteRequest) (models.CardDeleteRequest, error) {

	if _, err := r.db.ExecContext(
		ctx,
		"DELETE FROM cartehs WHERE id = ?",
		*deletecard.ID,
	); err != nil {
		return models.CardDeleteRequest{}, err
	}

	return deletecard, nil
}

func (r *SQLCardRepository) EditCard(ctx context.Context, updatecard models.CardUpdateRequest) (models.CardUpdateRequest, error) {
	if _, err := r.db.ExecContext(
		ctx,
		"UPDATE cartehs SET nom = ?, mana = ?, effets = ? WHERE id = ?",
		updatecard.Name, updatecard.Mana, updatecard.Effects, updatecard.ID,
	); err != nil {
		return models.CardUpdateRequest{}, err

	}
	return updatecard, nil
}
