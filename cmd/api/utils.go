package api

import (
	"math/rand"
	"net/url"
	"time"

	"github.com/speps/go-hashids/v2"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func isValidURL(testURL string) bool {
	_, err := url.ParseRequestURI(testURL)
	return err == nil
}

func (api *API) GenerateUniqueShortCode(url string) (string, error) {
	i := rand.Intn(len(url))
	hd := hashids.NewData()
	hd.Salt = url
	hd.MinLength = 6
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}
	e, err := h.Encode([]int{i})
	if err != nil {
		return "", err
	}

	return e, nil
}
