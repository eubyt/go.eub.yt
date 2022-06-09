package urlshort

import (
	"math/rand"
	"time"
)

const random_string_length = 10

type Shortener struct {
	Code string `json:"code"`
	Url  string `json:"url"`
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// eub.yt/HYLNr5GuVf
func CreateShortURL(url string, customUrl string) (Shortener, error) {
	var code = randomString(random_string_length)

	if customUrl != "" {
		code = customUrl
	}

	exist, err := DATABASE.CheckExistURL(url)
	if err != nil {
		return Shortener{}, err
	}

	// Se existir o URL, retorna o código
	if exist {
		code, err = DATABASE.GetCode(url)
		if err != nil {
			return Shortener{}, err
		}

		return Shortener{code, url}, nil
	}

	result, err := DATABASE.Insert(url, code)
	if err != nil {
		return Shortener{}, err
	}

	return Shortener{
		Code: result.code,
		Url:  result.url,
	}, nil
}

func SearchShortURL(code string) (Shortener, error) {
	exist, err := DATABASE.CheckExistCode(code)

	if err != nil || !exist {
		return Shortener{Code: "", Url: ""}, err
	}

	url, err := DATABASE.GetURL(code)
	if err != nil {
		return Shortener{Code: "", Url: ""}, err
	}

	return Shortener{
		code,
		url,
	}, nil
}
