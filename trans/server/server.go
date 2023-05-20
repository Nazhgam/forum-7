package server

import (
	"fmt"
	"forum/repo"
	"forum/service"
	"forum/trans/handler"
	"forum/trans/middleware"
	"log"
	"net/http"

	router "forum/trans/custom_router"
)

type Server struct {
	log     *log.Logger
	handler handler.Handler
}

func NewServer(l *log.Logger) {
	route := router.NewRouter()

	repo, err := repo.New(l)
	if err != nil {
		log.Fatalln("error on repo: %v", err)
	}

	svc := service.New(repo, l)

	handler := handler.New(svc, l)

	server := &Server{
		log:     &log.Logger{},
		handler: handler,
	}

	server.Routers(route)

	http.Handle("/", middleware.VerifyUser(route))
	fmt.Println("http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}
