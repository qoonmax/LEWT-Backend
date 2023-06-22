package translateService

import (
	"github.com/joho/godotenv"
	"os"
	"regexp"
	"strings"
)

var blacklistWords []string

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading env variables")
	}

	blacklistString := os.Getenv("BLACKLIST_WORDS")
	blacklistWords = strings.Split(blacklistString, ",")
}

func isNotEmptyString(inputText string) bool {
	return inputText != "" || len(inputText) > 3
}

func isRussianString(inputText string) bool {
	reg := regexp.MustCompile("[^а-яА-Я]")
	russianSymbols := reg.ReplaceAllString(inputText, "")

	percentage := float64(len(russianSymbols)) / float64(len(inputText))

	return percentage >= 0.7
}

func isWhiteString(inputText string) bool {
	inputText = strings.ToLower(inputText)

	for _, blackWord := range blacklistWords {
		if strings.Contains(inputText, blackWord) {
			return false
		}
	}
	return true
}
