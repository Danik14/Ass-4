package main

import (
	"ass_3/pkg/store/postgres"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// type application struct {
// }

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

}
