package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fadilahonespot/chatbot/repository/cached"
	"github.com/fadilahonespot/chatbot/repository/http/chatgbt"
	"github.com/fadilahonespot/chatbot/repository/mysql"
	"github.com/fadilahonespot/chatbot/server/handler"
	"github.com/fadilahonespot/chatbot/server/router"
	"github.com/fadilahonespot/chatbot/usecase"
	"github.com/fadilahonespot/chatbot/utils/database"
	"github.com/fadilahonespot/chatbot/utils/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Setup Env
	godotenv.Load()

	time.Sleep(30 * time.Second)

	// Setup Logger
	logger.NewLogger()

	// Setup Database
	db := database.InitDB()

	// Setup Mysql
	userRepo := mysql.NewUserRepository(db)
	chatRepo := mysql.NewChatRepository(db)

	// Setup Wrapper
	openAiWrapper := chatgbt.NewWrapper()
	cacheWrapper := cached.NewWrapper()

	// Setup Usecase
	userUsecase := usecase.NewUserUsecase(userRepo)
	chatUsecase := usecase.NewChatUsecase(userRepo, chatRepo, openAiWrapper, cacheWrapper)

	// Setup Handler
	userHandler := handler.NewUserHandler(userUsecase)
	chatHandler := handler.NewChatHandler(chatUsecase)

	// Setup Router
	route := router.NewRouter().
		SetUserHandler(userHandler).
		SetChatHandler(chatHandler).
		Validate()

	route.SetupRouter()

	port := os.Getenv("APP_PORT")
	fmt.Printf("Server running on :%v...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
