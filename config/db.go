package conn

import (
	"database/sql"
	"fmt"
	"os"
	"sample/utils"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func ConnectDB() (*sql.DB, error){

	err := godotenv.Load()
	utils.CheckErr(err)

		DB_HOST     := os.Getenv("DB_HOST")
		DB_PORT     := os.Getenv("DB_PORT")
		DB_USER     := os.Getenv("DB_USER")
		DB_PASSWORD := os.Getenv("DB_PASSWORD")
		DB_NAME     := os.Getenv("DB_NAME")
	
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	utils.CheckErr(err)
	return db, nil
}



