package main

import (
	"financial-tracker/src/api/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	router.GET("/expenses", handlers.ListExpenses)
	router.GET("/expenses/total-value", handlers.GetExpensesTotalValue)
	router.POST("/expenses", handlers.StoreExpense)

	err := router.Run("127.0.0.1:8080")
	if err != nil {
		panic("router error")
	}
	log.Println("starting server")
}
