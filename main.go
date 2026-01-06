package main

import (
	"github.com/TuxyBR/pizzaApi/models"
	"github.com/gin-gonic/gin"
)

var pizzas = []models.Pizza{
	{ID: 1, Name: "Toscana", Price: 49.5},
	{ID: 2, Name: "Marguerita", Price: 79.5},
	{ID: 3, Name: "Atum com queijo", Price: 69.5},
}

func main() {
	r := gin.Default()
	r.GET("/pizzas", getPizzas)
	r.POST("/pizzas", postPizzas)
	r.Run()
}

func getPizzas(c *gin.Context) {
	c.JSON(200, gin.H{"pizzas": pizzas})
}

func postPizzas(c *gin.Context) {
	var newPizza = models.Pizza{}
	err := c.ShouldBindJSON(&newPizza)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}
	pizzas = append(pizzas, newPizza)
}
