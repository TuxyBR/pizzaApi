package handler

import (
	"strconv"

	"github.com/TuxyBR/pizzaApi/internal/data"
	"github.com/TuxyBR/pizzaApi/internal/models"
	"github.com/gin-gonic/gin"
)

func GetPizzas(c *gin.Context) {
	c.JSON(200, gin.H{"pizzas": data.Pizzas})
}

func GetPizzaId(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}
	for _, piz := range data.Pizzas {
		if piz.ID == id {
			c.JSON(200, gin.H{"pizza": piz})
			return
		}
	}
	c.JSON(404, gin.H{"message": "pizza not found"})
}

func PostPizzas(c *gin.Context) {
	var newPizza = models.Pizza{}
	err := c.ShouldBindJSON(&newPizza)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}
	newPizza.ID = len(data.Pizzas) + 1
	data.Pizzas = append(data.Pizzas, newPizza)

	err = data.SavePizza()
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}
	c.JSON(201, newPizza)
}

func DeletePizzaId(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}
	for index, piz := range data.Pizzas {
		if piz.ID == id {
			data.Pizzas = append(data.Pizzas[:index], data.Pizzas[index+1:]...)
			err := data.SavePizza()
			if err != nil {
				c.JSON(400, gin.H{
					"erro": err.Error(),
				})
				return
			}
			c.JSON(200, gin.H{"message": "Pizza deleted"})
			return
		}
	}
	c.JSON(404, gin.H{"message": "Pizza not found"})
}

func UpdatePizza(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}

	var updatedPizza = models.Pizza{}
	err = c.ShouldBindJSON(&updatedPizza)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}

	for index, piz := range data.Pizzas {
		if piz.ID == id {
			oldPizza := data.Pizzas[index]
			data.Pizzas[index] = updatedPizza
			data.Pizzas[index].ID = id
			err := data.SavePizza()
			if err != nil {
				data.Pizzas[index] = oldPizza
				c.JSON(400, gin.H{
					"erro": err.Error(),
				})
				return
			}
			c.JSON(200, gin.H{"Pizza": data.Pizzas[index]})
			return
		}
	}
	c.JSON(404, gin.H{"message": "Pizza not found"})
}
