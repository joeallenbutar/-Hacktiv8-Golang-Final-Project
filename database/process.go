package database

import (
	"Final-Project/models"

	"github.com/jinzhu/gorm"
)

func GetToDos(db *gorm.DB) ([]models.Todo, error) {
	todos := []models.Todo{}
	query := db.Select("todos.*").Group("todos.id")
	if err := query.Find(&todos).Error; err != nil {
		return todos, err
	}

	return todos, nil
}

func GetToDoByID(id int, db *gorm.DB) (models.Todo, bool, error) {
	b := models.Todo{}

	query := db.Select("todos.*")
	query = query.Group("todos.id")
	err := query.Where("todos.id = ?", id).First(&b).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return b, false, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return b, false, nil
	}

	return b, true, nil
}

func UpdateToDo(db *gorm.DB, b *models.Todo) error {
	if err := db.Save(&b).Error; err != nil {
		return err
	}

	return nil
}

func DeleteToDo(id string, db *gorm.DB) error {
	var b models.Todo
	if err := db.Where("id = ? ", id).Delete(&b).Error; err != nil {
		return err
	}

	return nil
}