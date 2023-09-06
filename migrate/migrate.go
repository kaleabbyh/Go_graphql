package main

import (
	"fmt"
	"os"

	models "github.com/kaleabbyh/Food_Recipie/model"
	"github.com/kaleabbyh/Food_Recipie/utils"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)




func ConnectDB() (*gorm.DB, error) {

	err := godotenv.Load()
	utils.CheckErr(err)

		DB_HOST     := os.Getenv("DB_HOST")
		DB_PORT     := os.Getenv("DB_PORT")
		DB_USER     := os.Getenv("DB_USER")
		DB_PASSWORD := os.Getenv("DB_PASSWORD")
		DB_NAME     := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	utils.CheckErr(err)
	return db, nil
}




func main() {

	db, err := ConnectDB()
	utils.CheckErr(err)
	sqlDB, err := db.DB()
	utils.CheckErr(err)
	defer sqlDB.Close()


	// migrating User table
	err = db.AutoMigrate(&models.User{})
	utils.CheckErr(err)
	fmt.Println("User table successfully migrated")


	// migrating Category table
	err = db.AutoMigrate(&models.Category{})
	utils.CheckErr(err)
	fmt.Println("Category table successfully migrated")

	// migrating Recipe table
	err = db.AutoMigrate(&models.Recipe{})
	utils.CheckErr(err)
	fmt.Println("Recipe table successfully migrated")

	// migrating Ingredient table
	err = db.AutoMigrate(&models.Ingredient{})
	utils.CheckErr(err)
	fmt.Println("Ingredient table successfully migrated")
	
	// migrating Recipe_ingredient table
	err = db.AutoMigrate(&models.Recipe_ingredient{})
	utils.CheckErr(err)
	fmt.Println("Recipe_ingredient table successfully migrated")

	// migrating Step table
	err = db.AutoMigrate(&models.Step{})
	utils.CheckErr(err)
	fmt.Println("Step table successfully migrated")	
	

	// Perform database query
	var users []models.User
	db.Find(&users)


}