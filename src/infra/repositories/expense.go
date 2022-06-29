package repositories

import (
	"financial-tracker/src/entities"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

var dbAddress = "test.db"

func Save(expense entities.Expense) entities.Expense {
	db, err := gorm.Open(sqlite.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	//err = db.AutoMigrate(&entities.Expense{})
	//
	//if err != nil {
	//	panic("migration error")
	//}

	expense.Uuid = uuid.NewV4().String()
	expense.CreatedAt = time.Now()
	db.Create(expense)
	log.Printf("expense saved: %+v\n", expense)
	return expense
}

func List(expenses []entities.Expense) []entities.Expense {
	db, err := gorm.Open(sqlite.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.Find(&expenses)
	return expenses
}

func GetTotalValue() float32 {
	db, err := gorm.Open(sqlite.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	var totalValue int64
	db.Table("expenses").Select("sum(value)").Find(&totalValue)
	return float32(totalValue) / 100
}

//func FilterByDate(expenses []entities.Expense) []entities.Expense {
//	db, err := gorm.Open(sqlite.Open(dbAddress), &gorm.Config{})
//
//	if err != nil {
//		panic("failed to connect to database")
//	}
//
//	fromDate, _ := time.Parse(time.RFC3339, "2022-02-17T19:04:05-03:00")
//	toDate, _ := time.Parse(time.RFC3339, "2022-02-19T20:00:05-03:00")
//	db.Where("created_at BETWEEN ? AND ?", fromDate, toDate).Find(&expenses)
//
//	log.Printf("result: %+v\n", expenses)
//	return expenses
//}
