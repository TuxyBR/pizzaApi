package handler

import (
	"net/http"
	"strconv"

	"github.com/TuxyBR/pizzaApi/internal/data"
	"github.com/TuxyBR/pizzaApi/internal/models"
	"github.com/TuxyBR/pizzaApi/internal/service"
	"github.com/gin-gonic/gin"
)

func GetPizzas(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"pizzas": data.Pizzas})
}

func GetPizzaId(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}
	for _, piz := range data.Pizzas {
		if piz.ID == id {
			c.JSON(http.StatusOK, gin.H{"pizza": piz})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "pizza not found"})
}

func PostPizzas(c *gin.Context) {
	var newPizza = models.Pizza{}
	err := c.ShouldBindJSON(&newPizza)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}
	newPizza.ID = len(data.Pizzas) + 1
	err = service.ValidadePizza(&newPizza)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"erro": err.Error(),
		})
		return
	}
	data.Pizzas = append(data.Pizzas, newPizza)

	err = data.SavePizza()
	if err != nil {
		c.JSON(http.StatusNotModified, gin.H{
			"erro": err.Error(),
		})
		return
	}
	c.JSON(http.StatusFound, newPizza)
}

func DeletePizzaId(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}
	for index, piz := range data.Pizzas {
		if piz.ID == id {
			data.Pizzas = append(data.Pizzas[:index], data.Pizzas[index+1:]...)
			err := data.SavePizza()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"erro": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Pizza deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Pizza not found"})
}

func UpdatePizza(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	var updatedPizza = models.Pizza{}
	err = c.ShouldBindJSON(&updatedPizza)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
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
				c.JSON(http.StatusBadRequest, gin.H{
					"erro": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{"Pizza": data.Pizzas[index]})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Pizza not found"})
}
