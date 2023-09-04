package main

import (
	"fmt"
	models "sample/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = "Kaleabbyh@2"
	DB_NAME     = "postgres"
)


func ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}




func main() {

	db, err := ConnectDB()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	
	sqlDB, err := db.DB()
	if err != nil {
        fmt.Println("Failed to get underlying *sql.DB:", err)
        return
    }
	defer sqlDB.Close()
	// Auto-migrate the table
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Failed to perform auto migration:", err)
		return
	}

	// Perform database query
	var users []models.User
	db.Find(&users)

	fmt.Println(users)
}