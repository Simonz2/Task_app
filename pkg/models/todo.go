package models

import (
	"log"

	"github.com/Simonz2/Task_app/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Todo struct {
	gorm.Model `gorm:""`
	Completed  bool   `json:"completed"`
	Body       string `json:"body"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Todo{})
	db.AutoMigrate(&User{})
}
func (t *Todo) CreateTodo() *Todo {
	db.Create(&t)
	return t
}
func GetTodos() []Todo {
	var todos []Todo
	db.Find(&todos)
	return todos
}
func GetTodoById(id int64) *Todo {
	var getTodo Todo
	if err := db.Where("id=?", id).First(getTodo).Error; err != nil {
		return nil //not found or error
	}
	db.Where("id=?", id).First(&getTodo)
	return &getTodo
}
func DeleteTodo(id int64) Todo {
	var todo Todo
	if err := db.Where("id=?", id).First(&todo).Error; err != nil {
		return Todo{} //not found or error
	}
	db.Where("id=?", id).Delete(&todo)
	return todo
}

func PatchTodo(id int64) *Todo {
	var todo Todo
	// find todo by id
	log.Println("patching todo with id:", id)
	if err := db.Where("ID = ?", id).First(&todo).Error; err != nil {
		return nil // Not found or error
	}
	// update todo.Completed to true
	todo.Completed = true
	db.Save(&todo)
	return &todo
}
