package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	flog, err := os.OpenFile(`server.log`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer flog.Close()

	serverLogger := log.New(flog, "serv ", log.LstdFlags|log.Lshortfile)

	// Получаем текущую директорию
	currentDirectory, err := os.Getwd()
	if err != nil {
		serverLogger.Printf("Ошибка получения директории: %v\n", err)
		return
	}

	// если текущая дериктория cmd, то поднимаемся выше к html
	if filepath.Base(currentDirectory) == "cmd" {
		err = os.Chdir("..")
		if err != nil {
			serverLogger.Printf("Ошибка смены директории: %v\n", err)
			return
		}
	}

	s := server.Init(serverLogger)
	serverLogger.Println("Запуск сервера")
	err = s.Httpserver.ListenAndServe()
	if err != nil {
		serverLogger.Fatalf("Ошибка на стороне сервера: %v\n", err)
	}
}
