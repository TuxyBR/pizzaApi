package main

import (
	"github.com/TuxyBR/pizzaApi/internal/data"
	"github.com/TuxyBR/pizzaApi/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	data.LoadPizzas()
	data.LoadTeste()
	r := gin.Default()
	r.GET("/pizzas", handler.GetPizzas)
	r.GET("/pizzas/:id", handler.GetPizzaId)
	r.POST("/pizzas", handler.PostPizzas)
	r.DELETE("/pizzas/:id", handler.DeletePizzaId)
	r.PUT("/pizzas/:id", handler.UpdatePizza)

	r.Run()
}
