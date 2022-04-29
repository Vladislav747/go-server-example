package routes

import (
	"github.com/gorilla/mux"
	"github.com/akhil/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Method("POST")
	router.HandleFunc("/book/", controllers.CreateBook).Method("GET")
}