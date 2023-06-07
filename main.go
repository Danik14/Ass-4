package main

import (
	"ass_4/pkg/store/postgres"
	"ass_4/services/contact/internal"
	"ass_4/services/contact/internal/repository"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	db, err := postgres.OpenDB(port, user, password, dbname)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	repo := repository.NewRepository()

	useCase := internal.NewUseCase(repo)

	delivery := internal.NewDelivery(useCase, repo)

	fmt.Println(delivery)

	log.Fatal(http.ListenAndServe(":8080", nil))
}