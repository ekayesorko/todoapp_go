package conn

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"todolist-assignment/app/domain"
)

var dbGlob *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := "root:localpass@tcp(localhost:3306)/todoapp?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db err", err.Error())
	}
	err2 := db.AutoMigrate(&domain.User{}, &domain.Todo{})
	if err2 != nil {
		fmt.Println("migration failed")
	}
	dbGlob = db
	return dbGlob
}
