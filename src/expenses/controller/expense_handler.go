package controller

import (
	"financial-tracker/src/expenses/entity"
	"financial-tracker/src/expenses/repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func List(c *gin.Context) {
	var expenses []entity.Expense
	expenses = repository.List(expenses)
	c.JSON(http.StatusOK, expenses)
	log.Println("GET request processed")
}

func GetTotalValue(c *gin.Context) {
	var totalValue float32
	totalValue = repository.GetTotalValue()
	c.JSON(http.StatusOK, totalValue)
	log.Println("GET request processed")
}

//func FilterByDate(c *gin.Context) {
//	var expenses []entity.Expense
//	expenses = repository.FilterByDate(expenses)
//	c.JSON(http.StatusOK, expenses)
//	log.Println("GET request processed")
//}

func Store(c *gin.Context) {
	var expense entity.Expense

	err := c.BindJSON(&expense)
	if err != nil {
		log.Printf("POST request FAILED with %v\n", err)
		return
	}

	log.Printf("Expense built from request: %+v", expense)

	expense = repository.Save(expense)

	c.JSON(http.StatusOK, expense)
	log.Println("POST request processed")
}
