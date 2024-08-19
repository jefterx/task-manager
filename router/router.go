package router

import (
	"taskmanager/handlers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	return router
}
