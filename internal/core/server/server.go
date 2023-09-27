package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	host = "127.0.0.1"
	port = 8080
)

type HTTPServer interface {
	Start()
	Stop()
}

type httpServer struct {
	server *http.Server
}

func NewHTTPServer(router *gin.Engine) HTTPServer {
	return httpServer{
		server: &http.Server{
			Addr:           fmt.Sprintf("%s:%d", host, port),
			Handler:        router,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (sv httpServer) Start() {
	go func() {
		if err := sv.server.ListenAndServe(); err != nil {
			log.Printf("failed to stater HttpServer listen port %d, err=%s", port, err.Error())
		}
	}()
	log.Printf("Start Service with port %d", port)
}

func (sv httpServer) Stop() {
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Duration(3)*time.Second,
	)
	defer cancel()

	if err := sv.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown err=%s", err.Error())
	}
}
