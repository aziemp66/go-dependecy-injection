package main

import (
	"fmt"
	"net/http"

	"github.com/aziemp66/go-dependecy-injection/app"
	"github.com/aziemp66/go-dependecy-injection/controller"
	"github.com/aziemp66/go-dependecy-injection/helper"
	"github.com/aziemp66/go-dependecy-injection/middleware"
	"github.com/aziemp66/go-dependecy-injection/repository"
	"github.com/aziemp66/go-dependecy-injection/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("Server is running on port 3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
