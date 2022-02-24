package main

import (
	expenseController "financial-tracker/src/expenses/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	router.GET("/expenses", expenseController.List)
	router.GET("/expenses/total-value", expenseController.GetTotalValue)
	router.POST("/expenses", expenseController.Store)

	err := router.Run("127.0.0.1:8080")
	if err != nil {
		panic("router error")
	}
	log.Println("starting server")
}
