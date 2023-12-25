package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"user-service/internal/controller"
	"user-service/internal/core/server"
	"user-service/internal/core/service"
	"user-service/internal/infra/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := repository.NewDB()

	if err != nil {
		log.Printf("Error creating database, err = %s\n", err.Error())
		return
	}

	user_repo := repository.NewUserRepository(db)

	user_service := service.NewUserService(user_repo)

	router := gin.Default()
	user_controller := controller.NewUserController(router, user_service)
	user_controller.InitRouter()

	server := server.NewHTTPServer(router)
	server.Start()
	defer server.Stop()

	// Listen for OS signals to perform a graceful shutdown
	log.Println("listening signals...")
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	<-c
	log.Println("graceful shutdown...")
}
