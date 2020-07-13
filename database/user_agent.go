package database

import (
	"fmt"

	"github.com/jdxj/study-web/models"
)

const (
	tnUserAgent = "user_agent"
)

func AddUserAgent(ua *models.UserAgent) error {
	query := fmt.Sprintf("select count(*) from %s where value=?", tnUserAgent)
	row := db.QueryRow(query, ua.Value)
	var count int
	if err := row.Scan(&count); err != nil {
		return err
	}
	if count != 0 {
		return nil
	}

	query = fmt.Sprintf("insert into %s (value) values (?)", tnUserAgent)
	_, err := db.Exec(query, ua.Value)
	return err
}
