package handlers

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

type Handler struct {
	logger *log.Logger
}

func NewHandlers(logger *log.Logger) *Handler {
	return &Handler{logger: logger}
}

func (h *Handler) MainHandler(w http.ResponseWriter, r *http.Request) {

	// Получаем текущую директорию
	s, err := os.Getwd()
	if err != nil {
		h.logger.Printf("Ошибка получения директории: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// если текущая дериктория cmd, то поднимаемся выше к html
	if filepath.Base(s) == "cmd" {
		err = os.Chdir("..")
		if err != nil {
			h.logger.Printf("Ошибка смены директории: %v\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	// грузим html
	http.ServeFile(w, r, "index.html")
}

func (h *Handler) UploadHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		h.logger.Printf("Ошибка парсинга формы: %v\n", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		h.logger.Printf("Ошибка получения файла: %v\n", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileNew, err := os.Create(time.Now().UTC().String() + ".txt")
	if err != nil {
		h.logger.Printf("Ошибка при создании файла: %v\n", err)
		return
	}
	defer fileNew.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Fprint(fileNew, service.Translator(scanner.Text()))
	}
}
