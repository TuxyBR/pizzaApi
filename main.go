package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/TuxyBR/pizzaApi/models"
	"github.com/gin-gonic/gin"
)

var pizzas []models.Pizza

func main() {
	loadPizzas()
	r := gin.Default()
	r.GET("/pizzas", getPizzas)
	r.GET("/pizzas/:id", getPizzaId)
	r.POST("/pizzas", postPizzas)
	r.Run()
}

func getPizzas(c *gin.Context) {
	c.JSON(200, gin.H{"pizzas": pizzas})
}

func getPizzaId(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}
	for _, piz := range pizzas {
		if piz.ID == id {
			c.JSON(200, gin.H{"pizza": piz})
			return
		}
	}
	c.JSON(404, gin.H{"message": "pizza not found"})
}

func postPizzas(c *gin.Context) {
	defer savePizza()
	var newPizza = models.Pizza{}
	err := c.ShouldBindJSON(&newPizza)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}
	newPizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, newPizza)

	c.JSON(201, newPizza)
}

func loadPizzas() {
	file, err := os.Open("data/pizzas.json")
	if err != nil {
		fmt.Printf("ocorreu um erro ao tentar carregar o arquivo: %v\n", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&pizzas)
	if err != nil {
		fmt.Printf("ocorreu um erro decodificar o arquivo: %v\n", err)
		return
	}
}

func savePizza() {
	file, err := os.Create("data/pizzas.json")
	if err != nil {
		fmt.Printf("ocorreu um erro ao tentar carregar o arquivo: %v\n", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(pizzas)
	if err != nil {
		fmt.Printf("ocorreu um erro ao gerar o arquivo: %v\n", err)
		return
	}
}
