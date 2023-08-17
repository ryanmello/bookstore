package routes

import (
	"github.com/gorilla/mux"
	"github.com/ryanmello/bookstore-management-api/pkg/controllers"
)

var RegisterBookStoreRouter = func(router *mux.Router){
	router.HandleFunc("/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}