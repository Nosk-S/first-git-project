package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gin-project/internal/domain"
)

type CardRepository interface {
	SelectCards(ctx context.Context) ([]domain.Card, error)
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

	fmt.Println(cards)

	return cards, nil
}
