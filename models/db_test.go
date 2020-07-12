package models

import (
	"testing"
)

func TestDB(t *testing.T) {
	err := database.Ping()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	database.Close()
}

func TestAddUserAgent(t *testing.T) {
	defer CloseDB()
	ua := &UserAgent{
		Value: "test",
	}
	if err := AddUserAgent(ua); err != nil {
		t.Fatalf("%s\n", err)
	}
}
