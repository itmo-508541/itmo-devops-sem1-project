package main

import (
	"fmt"
	"log"
	"project_sem/internal/database"

	_ "github.com/gorilla/mux"
)

func main() {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			log.Fatal(panicErr)
		}
	}()

	fmt.Println(database.Connection().Config().Database)
}
