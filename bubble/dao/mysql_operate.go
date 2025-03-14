package dao

import (
	"bubble/databases"
	"bubble/models"
)

type Todo models.Todo

func CreateATodo(todo *models.Todo) (err error) {
	if err := databases.DB.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func GetAllTodos() (todoList []*Todo, err error) {
	if err := databases.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetATodoById(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = databases.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func UpdateATodo(todo *Todo) (err error) {
	if err := databases.DB.Save(todo).Error; err != nil {
		return err
	}
	return nil
}

func DeleteATodo(id string) (err error) {
	if err := databases.DB.Where("id=?", id).Delete(&Todo{}).Error; err != nil {
		return err
	}
	return nil
}
