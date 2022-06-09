package urlshort

import (
	"testing"
)

func TestRandomString(t *testing.T) {
	for i := 0; i < 100; i++ {
		s := randomString(10)
		if len(s) != 10 {
			t.Errorf("RandomString(10) = %s, want 10", s)
		}
		t.Log(s)
	}
}

func TestCreateURLShort(t *testing.T) {
	var code string
	CreateDatabase()
	url := "https://www.twitch.tv/l4k3"
	for i := 0; i < 10; i++ {
		shortener, err := CreateShortURL(url)
		if err != nil {
			t.Errorf("CreateShortURL(%s) = %s, want nil", url, err)
		}
		if code == "" {
			code = shortener.Code
		}
		if shortener.Code != code {
			t.Errorf("CreateShortURL(%s) = %s, want %s", url, shortener.Code, code)
		}
		t.Log(shortener)
	}
}
