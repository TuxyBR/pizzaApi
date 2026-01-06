package service

import (
	"fmt"

	"github.com/TuxyBR/pizzaApi/internal/models"
)

func ValidadePizza(pizza *models.Pizza) error {
	if pizza.Price < 0 {
		return fmt.Errorf("o preco da pizza nao pode ser negativo")
	}
	return nil
}
