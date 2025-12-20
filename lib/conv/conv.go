package conv

import (
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}


func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func GenerateSlug(title string) string {
	// Convert to lowercase
	slug := strings.ToLower(title)	
	// Replace spaces and special characters with hyphens
	slug = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(slug, "-")
	// Trim hyphens from the start and end
	slug = strings.Trim(slug, "-")
	
	return slug
}

func StringToInt64(s string) (int64, error) {
	newData, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		return 0, err
	}

	return newData, nil
}

func StringToInt(s string) (int, error) {
	number, err := strconv.Atoi(s)

	if err != nil {
		return 0, err
	}

	return number, nil
}







