package handlers

import (
	"financial-tracker/src/entities"
	"financial-tracker/src/infra/repositories"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ListExpenses(c *gin.Context) {
	var expenses []entities.Expense
	expenses = repositories.List(expenses)
	c.JSON(http.StatusOK, expenses)
}

func GetExpensesTotalValue(c *gin.Context) {
	var totalValue float32
	totalValue = repositories.GetTotalValue()
	c.JSON(http.StatusOK, totalValue)
}

//func FilterByDate(c *gin.Context) {
//	var expenses []entities.Expense
//	expenses = repositories.FilterByDate(expenses)
//	c.JSON(http.StatusOK, expenses)
//	log.Println("GET request processed")
//}

func StoreExpense(c *gin.Context) {
	var expense entities.Expense
	err := c.BindJSON(&expense)
	if err != nil {
		log.Printf("POST request FAILED with %v\n", err)
		return
	}
	log.Printf("Expense built from request: %+v", expense)
	expense = repositories.Save(expense)
	c.JSON(http.StatusOK, expense)
}
