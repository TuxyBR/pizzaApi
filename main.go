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

	r.GET("/teste", handler.GetTeste)
	r.GET("/teste/:id", handler.GetTesteId)
	r.POST("/teste", handler.PostTeste)
	r.DELETE("/teste/:id", handler.DeleteTesteId)
	r.PUT("/teste/:id", handler.UpdateTeste)

	r.Run()
}
