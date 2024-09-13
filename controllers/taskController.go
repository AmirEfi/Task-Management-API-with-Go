package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    "gorm.io/gorm"
)

var DB *gorm.DB

func GetTasks(w http.ResponseWriter, r *http.Request) {
    var tasks []Task
    DB.Find(&tasks)
    json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var task Task
    if err := DB.First(&task, id).Error; err != nil {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(task)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task Task
    json.NewDecoder(r.Body).Decode(&task)

    DB.Create(&task)
    json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var task Task
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

    var task Task
    if err := DB.First(&task, id).Error; err != nil {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    DB.Delete(&task)
    json.NewEncoder(w).Encode("Task deleted")
}