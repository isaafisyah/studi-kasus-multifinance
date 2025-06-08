package database

import (
	"fmt"
	"log"
	"os"

	"github.com/isaafisyah/studi-kasus-multifinance/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabaseConnection(cnf *config.Config) (*gorm.DB, error) {
	fmt.Println("Welcome to " + cnf.Server.Name)
	//connect to database mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Connected to MySQL successfully!")
	return db, nil
}