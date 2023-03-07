package repository

import (
	"fmt"
	"grocerylist/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type itemRepo struct {
	db *gorm.DB
}

func NewItemRepository() itemRepo {
	log.Println(models.Item{})
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "test-user:password123@tcp(127.0.0.1:3306)/books?charset=utf8&parseTime=True&loc=Local",
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		db.AutoMigrate(&models.Item{})
		newRepo := itemRepo{db}
		return newRepo
	}
}

func (repo *itemRepo) CreateItem(item string) (bool, error) {
	row := models.Item{Name: item}
	res := repo.db.Create(&row)
	if res.RowsAffected == 0 {
		return false, fmt.Errorf("failed to insert")
	}
	return true, nil
}

func (repo *itemRepo) DeleteItem(item string) (bool, error) {
	row := models.Item{Name: item}
	res := repo.db.Where("name = ?", item).Delete(&row)
	if res.RowsAffected == 0 {
		return false, fmt.Errorf("failed to delete")
	}
	return true, nil
}

func (repo *itemRepo) UpdateItem(item string) (bool, error) {
	row := models.Item{Name: item}
	res := repo.db.Find(row)
	if res.RowsAffected == 0 {
		return false, fmt.Errorf("not found")
	}
	repo.db.Save(&row)
	repo.db.Model(&row).Update("name", item)
	return true, nil
}

func (repo *itemRepo) GetItem(item string) *models.Item {
	row := models.Item{Name: item}
	repo.db.Where("name = ?", item).First(row)
	return &row
}

func (repo *itemRepo) GetList(list *[]models.Item) {
	repo.db.Find(list)
}
