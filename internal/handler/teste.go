package handler

import (
	"strconv"

	"github.com/TuxyBR/pizzaApi/internal/data"
	"github.com/TuxyBR/pizzaApi/internal/models"
	"github.com/gin-gonic/gin"
)

func GetTeste(c *gin.Context) {
	c.JSON(200, gin.H{"teste": data.Teste})
}

func GetTesteId(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}
	for _, test := range data.Teste {
		if test.IDOc == id {
			c.JSON(200, gin.H{"teste": test})
			return
		}
	}
	c.JSON(404, gin.H{"message": "teste not found"})
}

func PostTeste(c *gin.Context) {
	var newTeste = models.Teste{}
	err := c.ShouldBindJSON(&newTeste)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}
	data.Teste = append(data.Teste, newTeste)
	err = data.SaveTeste()
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}
	c.JSON(201, newTeste)
}

func DeleteTesteId(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}
	for index, test := range data.Teste {
		if test.IDObra == id {
			data.Teste = append(data.Teste[:index], data.Teste[index+1:]...)
			err := data.SaveTeste()
			if err != nil {
				c.JSON(400, gin.H{
					"erro": err.Error(),
				})
				return
			}
			c.JSON(200, gin.H{"message": "teste deleted"})
			return
		}
	}
	c.JSON(404, gin.H{"message": "teste not found"})
}

func UpdateTeste(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}

	var updatedTeste = models.Teste{}
	err = c.ShouldBindJSON(&updatedTeste)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}

	for index, test := range data.Teste {
		if test.IDOc == id {
			oldTeste := data.Teste[index]
			data.Teste[index] = updatedTeste
			data.Teste[index].IDOc = id
			err := data.SaveTeste()
			if err != nil {
				data.Teste[index] = oldTeste
				c.JSON(400, gin.H{
					"erro": err.Error(),
				})
				return
			}
			c.JSON(200, gin.H{"Teste": data.Teste[index]})
			return
		}
	}
	c.JSON(404, gin.H{"message": "teste not found"})
}
