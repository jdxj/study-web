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
