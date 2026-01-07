package service

import (
	"fmt"

	"github.com/TuxyBR/pizzaApi/internal/models"
)

func ValidadeReview(pizza *models.Review) error {
	if pizza.Rating < 0 {
		return fmt.Errorf("o review da pizza nao pode ser negativo")
	}
	if pizza.Rating == 0 {
		return fmt.Errorf("o review da pizza deve ser entre 1 e 5")
	}
	if pizza.Rating > 5 {
		return fmt.Errorf("o review da pizza deve ser entre 1 e 5")
	}
	if pizza.Comment == "" {
		return fmt.Errorf("a descrição deve ser preenchida")
	}
	return nil
}
