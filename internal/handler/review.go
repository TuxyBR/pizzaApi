package handler

import (
	"net/http"
	"strconv"

	"github.com/TuxyBR/pizzaApi/internal/data"
	"github.com/TuxyBR/pizzaApi/internal/models"
	"github.com/TuxyBR/pizzaApi/internal/service"
	"github.com/gin-gonic/gin"
)

func PostReview(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	var newReview = models.Review{}
	err = c.ShouldBindJSON(&newReview)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	err = service.ValidadeReview(&newReview)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	var selectedPizza = models.Pizza{}
	found := false
	for index, pizza := range data.Pizzas {
		if pizza.ID == id {
			newReview.ID = len(pizza.Review) + 1
			data.Pizzas[index].Review = append(data.Pizzas[index].Review, newReview)
			selectedPizza = data.Pizzas[index]
			found = true
			break
		}
	}

	if found {
		err = data.SavePizza()
		if err != nil {
			c.JSON(http.StatusNotModified, gin.H{
				"erro": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, selectedPizza)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"erro": "no pizza found",
		})
	}
}
