package models

import (
	"fmt"
)

const (
	tnUserAgent = "user_agent"
)

type UserAgent struct {
	ID    int
	Value string
}

func AddUserAgent(ua *UserAgent) error {
	query := fmt.Sprintf("select count(*) from %s where value=?", tnUserAgent)
	row := database.QueryRow(query, ua.Value)
	var count int
	if err := row.Scan(&count); err != nil {
		return err
	}
	if count != 0 {
		return nil
	}

	query = fmt.Sprintf("insert into %s (value) values (?)", tnUserAgent)
	_, err := database.Exec(query, ua.Value)
	return err
}
