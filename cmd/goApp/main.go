package main

import (
	"goapp/config"
	"goapp/internal/database"
	"goapp/internal/handler"
	"goapp/internal/repository"
	"goapp/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	r := mux.NewRouter()
	handler.RegisterUserRoutes(r, userService)

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
