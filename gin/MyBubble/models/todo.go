package models

import (
	"MyBubble/dao"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// Todo 增删改查
func CreateATodo(todo *Todo) (err error) {
	if err = dao.DB.Create(&todo).Error; err != nil {
		return err
	}
	return
}

func GetAllTodo() (todoList []*Todo, err error) {
	if err := dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return
}

func Update(todo *Todo) (err error) {
	if err = dao.DB.Save(todo).Error; err != nil {
		return
	}
	return
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
