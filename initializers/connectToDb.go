package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectToDb() {

	var err error
	dsn := os.Getenv("DB_URL")
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connnect to DB")
	}
	fmt.Println("Connected to DB")
}
