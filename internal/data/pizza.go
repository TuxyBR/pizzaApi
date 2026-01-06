package data

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TuxyBR/pizzaApi/internal/models"
)

var Pizzas []models.Pizza

func LoadPizzas() {
	file, err := os.Open("data/pizzas.json")
	if err != nil {
		fmt.Printf("ocorreu um erro ao tentar carregar o arquivo: %v\n", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Pizzas)
	if err != nil {
		fmt.Printf("ocorreu um erro decodificar o arquivo: %v\n", err)
		return
	}
}

func SavePizza() {
	file, err := os.Create("data/pizzas.json")
	if err != nil {
		fmt.Printf("ocorreu um erro ao tentar carregar o arquivo: %v\n", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(Pizzas)
	if err != nil {
		fmt.Printf("ocorreu um erro ao gerar o arquivo: %v\n", err)
		return
	}
}
