package main

import (
	"ApiGolang/configs"
	"ApiGolang/internal/entity"
	"ApiGolang/internal/infra/database"
	"ApiGolang/internal/webserver/handlers"
	"net/http"

	"github.com/go-chi/chi"

	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter
	r.Use(Middlewares.Logger)
	r.POST("/product", productHandler.CreateProduct)
	r.GET("/product/{id}", productHandler.GetProduct)
	r.PUT("/product/{id}", productHandler.UpdateProduct)

	http.ListenAndServe("8000", r)

}
