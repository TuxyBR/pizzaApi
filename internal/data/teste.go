package data

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TuxyBR/pizzaApi/internal/models"
)

var Teste []models.Teste

func SaveTeste() error {
	file, err := os.Create("internal/data/teste.json")
	if err != nil {
		return fmt.Errorf("ocorreu um erro ao tentar carregar o arquivo: %v\n", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(Teste)
	if err != nil {
		return fmt.Errorf("ocorreu um erro ao gerar o arquivo: %v\n", err)
	}
	return nil
}

func LoadTeste() {
	file, err := os.Open("internal/data/teste.json")
	if err != nil {
		fmt.Printf("ocorreu um erro ao tentar carregar o arquivo: %v\n", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Teste)
	if err != nil {
		fmt.Printf("ocorreu um erro decodificar o arquivo: %v\n", err)
		return
	}
}
