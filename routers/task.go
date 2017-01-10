package routers

import (
	"github.com/colombia9503/RESTful-MDB/controllers"
	"github.com/gorilla/mux"
)

//task routes
func SetTaskRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/tasks", controllers.Task.Create).Methods("POST")
	router.HandleFunc("/tasks/{id}", controllers.Task.FindOne).Methods("GET")
	router.HandleFunc("/tasks", controllers.Task.FindAll).Methods("GET")
	router.HandleFunc("/tasks/{id}", controllers.Task.Update).Methods("PUT")
	router.HandleFunc("/tasks/{id}", controllers.Task.Delete).Methods("DELETE")
	return router
}
