package database

import (
	"testing"

	"github.com/jdxj/study-web/models"
)

func TestDB(t *testing.T) {
	err := db.Ping()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	db.Close()
}

func TestAddUserAgent(t *testing.T) {
	defer Close()
	ua := &models.UserAgent{
		Value: "test",
	}
	if err := AddUserAgent(ua); err != nil {
		t.Fatalf("%s\n", err)
	}
}
