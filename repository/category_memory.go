package repository

import (
	"errors"

	"github.com/erikprogramador/repository-patter-go-lang/entity"
)

type CategoriesMemoryDb struct {
	Categories []entity.Category
}

type CategoryRepositoryMemory struct {
	db CategoriesMemoryDb
}

func NewCategoryRepositoryMemory(db CategoriesMemoryDb) *CategoryRepositoryMemory {
	return &CategoryRepositoryMemory{db: db}
}

func (c *CategoryRepositoryMemory) Get(id string) (entity.Category, error) {
	for _, category := range c.db.Categories {
		if category.ID == id {
			return category, nil
		}
	}
	return entity.Category{}, errors.New("No category found with this id")
}

func (c *CategoryRepositoryMemory) Create(category entity.Category) (entity.Category, error) {
	c.db.Categories = append(c.db.Categories, category)
	return category, nil
}
