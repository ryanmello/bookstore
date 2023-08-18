package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ryanmello/bookstore-management-api/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRouter(r)

	fmt.Printf("Starting server on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}