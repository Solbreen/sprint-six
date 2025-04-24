package service

import (
	"regexp"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Translator(text string) string {
	if isText(text) {
		return morse.ToMorse(text)
	}

	return morse.ToText(text)
}

func isText(text string) bool {
	russianRegex := regexp.MustCompile(`[а-яА-ЯёЁ]`)
	return russianRegex.MatchString(text)
}
