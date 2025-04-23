package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	flog, err := os.OpenFile(`server.log`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer flog.Close()

	serverLogger := log.New(flog, "serv ", log.LstdFlags|log.Lshortfile)

	s := server.Init(serverLogger)
	serverLogger.Println("Запуск сервера")
	err = s.Httpserver.ListenAndServe()
	if err != nil {
		serverLogger.Fatalf("Ошибка на стороне сервера: %v\n", err)
	}
}
