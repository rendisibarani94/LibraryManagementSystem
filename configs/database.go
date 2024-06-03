package configs

import (
	"first-jwt/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// init style for gorm (go orm)
var DB *gorm.DB

// connectDB function
func ConnectDB() {
	//Open function to gain connection to db

	if errEnv := godotenv.Load(); errEnv != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("Failed To Connect To Database")
	}

	// Enable logging for debugging (print SQL queries)
    db = db.Debug()

	// database migration function
	db.AutoMigrate(
		&models.User{},
		&models.Book{},
		&models.BorrowedBook{},
	)

	//refering the DB(gorm.DB) variable to db(configured) variable
	DB = db
	log.Println("Database Connected")	
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close a connection to database")
	}
	dbSQL.Close()
}