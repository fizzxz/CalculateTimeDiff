package router

import (
	"../database"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/task", database.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", database.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/{id}", database.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", database.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", database.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTask", database.DeleteAllTask).Methods("DELETE", "OPTIONS")
	return router
}
