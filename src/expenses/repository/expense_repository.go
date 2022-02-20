package repository

import (
	"financial-tracker/src/expenses/entity"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

func Save(expense entity.Expense) entity.Expense {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	//err = db.AutoMigrate(&Expense{})
	//
	//if err != nil {
	//	panic("migration error")
	//}

	expense.Uuid = uuid.NewV4().String()
	expense.CreatedAt = time.Now()
	db.Create(expense)
	log.Printf("expense saved: %+v\n", expense)

	var expenses []entity.Expense
	fromDate, _ := time.Parse(time.RFC3339, "2022-02-17T15:04:05-03:00")
	toDate, _ := time.Parse(time.RFC3339, "2022-02-19T15:04:05-03:00")
	db.Where("created_at BETWEEN ? AND ?", fromDate, toDate).Find(&expense)

	log.Printf("result: %+v\n", expenses)
	return expense
}

func List(expenses []entity.Expense) []entity.Expense {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	db.Find(&expenses)

	log.Printf("result: %+v\n", expenses)
	return expenses
}

func FilterByDate(expenses []entity.Expense) []entity.Expense {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	fromDate, _ := time.Parse(time.RFC3339, "2022-02-17T19:04:05-03:00")
	toDate, _ := time.Parse(time.RFC3339, "2022-02-19T20:00:05-03:00")
	db.Where("created_at BETWEEN ? AND ?", fromDate, toDate).Find(&expenses)

	log.Printf("result: %+v\n", expenses)
	return expenses
}
