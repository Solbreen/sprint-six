package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger     *log.Logger
	Httpserver http.Server
}

func Init(logger *log.Logger) *Server {

	server := &Server{Logger: logger}
	handler := handlers.NewHandlers(logger)
	// Так, тут делаем маршрутизатор и хендлеры регаем
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.MainHandler)
	mux.HandleFunc("/upload", handler.UploadHandler)

	// Экземпляр для сервера
	server.Httpserver = http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return server
}
