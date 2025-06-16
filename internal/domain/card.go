package domain

type Card struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Mana    int8   `json:"mana"`
	Effects string `json:"effects"`
}
