package handlers

import (
	"bufio"
	"fmt"

	"log"
	"net/http"
	"os"
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
	var str string
	for scanner.Scan() {
		str = str + service.Translator(scanner.Text())
		fmt.Fprint(fileNew, service.Translator(scanner.Text()))
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(str))
}
