package urlshort

import (
	"testing"
)

func CreateDatabase() {
	DATABASE.Connect("urlshort.test")
	DATABASE.CreateTable()
}

func TestInsert(t *testing.T) {
	CreateDatabase()
	result, err := DATABASE.Insert("https://www.twitch.tv/l4k3", "revolução")
	if err != nil {
		t.Errorf("error inserting record: %v", err)
	}
	t.Logf("result: %v", result)
}

func TestGetCode(t *testing.T) {
	CreateDatabase()
	code, err := DATABASE.GetCode("https://www.twitch.tv/l4k3")
	if err != nil {
		t.Errorf("error getting code: %v", err)
	}
	t.Logf("code: %v", code)
}

func TestGetURL(t *testing.T) {
	CreateDatabase()
	url, err := DATABASE.GetURL("revolução")
	if err != nil {
		t.Errorf("error getting url: %v", err)
	}
	t.Logf("url: %v", url)
}

func TestCheckExistURL(t *testing.T) {
	CreateDatabase()
	exists, err := DATABASE.CheckExistURL("https://www.twitch.tv/l4k3")
	if err != nil {
		t.Errorf("error checking url: %v", err)
	}

	if !exists {
		t.Errorf("url does not exist")
	}

	t.Logf("exists: %v", exists)
}

func TestCheckNoExistURL(t *testing.T) {
	CreateDatabase()
	exists, err := DATABASE.CheckExistURL("http://www.google.com/not-found")
	if err != nil {
		t.Errorf("error checking url: %v", err)
	}

	if exists {
		t.Errorf("url should not exist")
	}

	t.Logf("exists: %v", exists)
}
