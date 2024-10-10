package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rfulgencio3/go-personal-library/configs"
	"github.com/rfulgencio3/go-personal-library/internal/handler"
	"github.com/rfulgencio3/go-personal-library/internal/repository/mongodb"
	"github.com/rfulgencio3/go-personal-library/internal/usecase"

	_ "github.com/rfulgencio3/go-personal-library/docs"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/rfulgencio3/go-personal-library/internal/middleware"
)

// @title Go Personal Library API
// @version 1.0
// @description API for managing personal library of books.

// @contact.name Ricardo Fulgencio
// @contact.url https://github.com/rfulgencio3
// @contact.email rfulgencio3@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

// @schemes http
func main() {
	// Carregar a configuração do arquivo .env
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar a configuração: %v", err)
	}

	// Conectar ao MongoDB
	client, err := mongodb.NewMongoClient(config)
	if err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	}

	// Inicializar Repositório, UseCase e Handler
	bookRepo := mongodb.NewBookRepository(client, config)
	bookUseCase := usecase.NewBookUseCase(bookRepo)
	bookHandler := handler.NewBookHandler(bookUseCase)

	readBookRepo := mongodb.NewReadBookRepository(client, config)
	readBookUC := usecase.NewReadBookUseCase(readBookRepo)
	readBookHandler := handler.NewReadBookHandler(readBookUC)

	// Configurar Rotas
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)
	bookHandler.RegisterRoutes(router)
	readBookHandler.RegisterRoutes(router)

	// Registrar o handler do Swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Iniciar Servidor HTTP
	log.Printf("Servidor está executando na porta %s", config.ServerPort)
	if err := http.ListenAndServe(":"+config.ServerPort, router); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
