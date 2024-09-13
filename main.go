package main

import (
    "task-api/controllers"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func main() {
    var err error

    DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database")
    }

    DB.AutoMigrate(&controllers.Task{})
    r := mux.NewRouter()

    r.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
    r.HandleFunc("/tasks/{id}", controllers.GetTask).Methods("GET")
    r.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
    r.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
    r.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")

    log.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}