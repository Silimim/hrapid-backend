package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type ContextKey string

const UserKey ContextKey = "user"

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

func DebugRequest(r *http.Request) {
	log.Printf("Request Method: %s", r.Method)
	log.Printf("Request URL: %s", r.URL.String())
	log.Printf("Request Headers:")
	for key, values := range r.Header {
		log.Printf("  %s: %v", key, values)
	}

	log.Println("Context values:")
	PrintContextInternals(r.Context(), 0)
}

func PrintContextInternals(ctx context.Context, depth int) {
	indent := fmt.Sprintf("%*s", depth*2, "")

	log.Printf("%sContext type: %T", indent, ctx)

	if val := ctx.Value(UserKey); val != nil {
		log.Printf("%sUser value found: %+v", indent, val)
	} else {
		log.Printf("%sNo user value found with key 'user'", indent)
	}

	if nextCtx := ctx.Value(struct{}{}); nextCtx != nil {
		if nextCtx != ctx {
			PrintContextInternals(nextCtx.(context.Context), depth+1)
		}
	}
}

func IsPointerType(fieldType string) bool {
	return len(fieldType) > 0 && fieldType[0] == '*'
}

func SendJSONResponse(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
