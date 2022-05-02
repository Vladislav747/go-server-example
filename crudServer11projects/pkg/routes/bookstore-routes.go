package routes

import (
	"github.com/akhil/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Method("POST")
	router.HandleFunc("/book/", controllers.CreateBook).Method("GET")
	router
}
