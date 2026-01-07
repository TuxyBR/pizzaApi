package models

type Review struct {
	ID      int    `json:"id"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}
