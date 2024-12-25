package main

import (
	"database/sql"
	v2 "github.com/dariuszdroba/go-from-template/internal/controller/http/v2"
	"github.com/dariuszdroba/go-from-template/internal/usecase"
	"github.com/dariuszdroba/go-from-template/internal/usecase/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test") // will be replaced with .env var
	if err != nil {
		log.Fatal(err)
	}
	productRepo := repository.NewProductRepository(db)
	productUC := usecase.NewProductUseCase(productRepo)
	productHandler := v2.NewProductHandler(productUC)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// Register product routes
	productHandler.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
