package utils

import (
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func SplitCamelCase(s string) string {
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	split := re.ReplaceAllString(s, `$1 $2`)
	return ToLowercaseFirstLetter(split)
}

func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func ToLowercaseFirstLetter(s string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		if i == 0 {
			continue
		}
		words[i] = strings.ToLower(word[:1]) + word[1:]
	}
	return strings.Join(words, " ")
}
