package models //ici que je place toute mes structure

type CardResearch struct {
	Research string  `json:"research"`
	ID       *int    `json:"id"`
	Name     *string `json:"name"`
	Mana     *int8   `json:"mana"`
	Effects  *string `json:"effects"`
}

type CardInsertRequest struct {
	ID      *int    `json:"id"`
	Name    *string `json:"name"`
	Mana    *int8   `json:"mana"`
	Effects *string `json:"effects"`
}

type CardDeleteRequest struct {
	ID *int `json:"id"`
}

type CardUpdateRequest struct {
	ID      *int    `json:"id"`
	Name    *string `json:"name"`
	Mana    *int8   `json:"mana"`
	Effects *string `json:"effects"`
}
