package main

import (
	"database/sql"
	v2 "github.com/dariuszdroba/go-from-template/internal/controller/http/v2"
	"github.com/dariuszdroba/go-from-template/internal/usecase"
	"github.com/dariuszdroba/go-from-template/internal/usecase/repository"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	productRepo := repository.NewProductRepository(db)
	productUC := usecase.NewProductUseCase(productRepo)
	productHandler := v2.NewProductHandler(productUC)

	r := gin.Default()

	// Register product routes
	productHandler.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
