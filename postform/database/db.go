package database

import (
	"crud-test/handlers"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB(DB *gorm.DB) handlers.Handler {
	return handlers.Handler{DB}
}

func Init() *gorm.DB {
	err := godotenv.Load(".env")

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao realizar conectar com o postgress")
	}
	fmt.Println("Sucesso ao realizar conexão com o postgres")
	db.AutoMigrate(&handlers.Users{})
	return db

}
