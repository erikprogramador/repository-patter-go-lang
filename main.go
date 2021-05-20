package main

import (
	"database/sql"
	"fmt"

	"github.com/erikprogramador/repository-patter-go-lang/repository"
	"github.com/erikprogramador/repository-patter-go-lang/service"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// db := repository.CategoriesMemoryDb{Categories: []entity.Category{}}
	// repositoryMemory := repository.NewCategoryRepositoryMemory(db)

	db, _ := sql.Open("sqlite3", "./sqlite.repository")
	repositorySqlite := repository.NewCategorySqlite(db)

	service := service.NewCategoryService(repositorySqlite)
	cat, _ := service.Create("Minha categoria")
	fmt.Println(cat)
}
