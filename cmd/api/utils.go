package api

import (
	"math/rand"
	"net/url"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func isValidURL(testURL string) bool {
	_, err := url.ParseRequestURI(testURL)
	return err == nil
}

func (api *API) GenerateUniqueShortCode() (string, error) {
	const length = 6

	for i := 0; i < 5; i++ {
		code := generateRandomCode(length)
		var exists bool
		err := api.DB.Connection.QueryRow("SELECT EXISTS(SELECT 1 FROM urls WHERE short_code = $1)", code).
			Scan(&exists)
		if err != nil {
			return "", err
		}
		if !exists {
			return code, nil
		}
	}
	return "", nil
}

func generateRandomCode(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
