package repository

import (
	"database/sql"
	"gin-project/internal/domain"
)

type CardRepository interface {
}

type CardRepositoryImpl struct {
	DB *sql.DB
}

func NewCardRepository(db *sql.DB) CardRepository {
	return &CardRepositoryImpl{DB: db}
}

func (r *CardRepositoryImpl) SelectCards() *[]domain.Card {
	rows, err := r.DB.Query("SELECT id, nom, mana, effets FROM cartehs ")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var cards []domain.Card
	for rows.Next() {
		var crt domain.Card

		if err := rows.Scan(&crt.ID, &crt.Name, &crt.Mana, &crt.Effects); err != nil {
			panic(err)
		}

		cards = append(cards, crt)
	}

	return &cards
}
