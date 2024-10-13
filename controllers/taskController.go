package controllers

import (
    "Task-Management-API-with-Go/models"
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(database *gorm.DB) {
    DB = database
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
    var tasks []models.Task
    if err := DB.Find(&tasks).Error; err != nil {
        http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var task models.Task
    if err := DB.First(&task, id).Error; err != nil {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(task)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    json.NewDecoder(r.Body).Decode(&task)

    DB.Create(&task)
    json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var task models.Task
    if err := DB.First(&task, id).Error; err != nil {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    json.NewDecoder(r.Body).Decode(&task)
    DB.Save(&task)

    json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var task models.Task
    if err := DB.First(&task, id).Error; err != nil {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    DB.Delete(&task)
    json.NewEncoder(w).Encode("Task deleted")
}