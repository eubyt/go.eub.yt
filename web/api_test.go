package web

import (
	"testing"
)

func TestIsUrlValid(t *testing.T) {
	listString := []string{
		"https://www.twitch.tv/l4k3",
		"https://www.twitch.tv/l4k3/",
		"https://www.twitch.tv/l4k3/videos",
		"https://www.twitch.tv/l4k3/videos/",
		"https://www.twitch.tv/l4k3/videos/12345",
		"https://www.twitch.tv/l4k3/videos/12345/",
		"https://www.twitch.tv/l4k3/videos/12345/comments?page=1",
		"https://www.twitch.tv/l4k3/videos/12345/comments?page=1#comments",
		"https://www.twitch.tv/l4k3/videos/12345/comments?page=1&comments=12345",

		"https://twitch.tv/l4k3",
		"http://twitch.tv/l4k3",
	}

	for _, url := range listString {
		if !IsUrl(url) {
			t.Errorf("%s is not a valid url", url)
		} else {
			t.Logf("%s is a valid url", url)
		}
	}
}

func TestIsUrlNotValid(t *testing.T) {
	listString := []string{
		"aaaaa",
		"http://",
	}

	for _, url := range listString {
		if isUrl(url) {
			t.Errorf("%s is a valid url", url)
		} else {
			t.Logf("%s is not a valid url", url)
		}
	}
}
