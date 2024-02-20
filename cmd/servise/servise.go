package main

import (
	"fmt"
	"log"
	"os"
	"postgres/pkg/storage"
	"postgres/pkg/storage/postgres"
)

// Интерфейс БД.
var db storage.Interface

func main() {
	var err error
	pwd := os.Getenv("DBPASS")
	if pwd == "" {
		os.Exit(1)
	}
	connstr :=
		"postgres://postgres:" +
			pwd + "@127.0.0.1/Task"
	// присвоение переменной типа интерфейс конкретной реализации БД
	log.Println("1") // метка 1
	db, err = postgres.New(connstr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("2") // метка 2
	tasks, err := db.Tasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("3") // метка 3
	fmt.Println(tasks)
}
