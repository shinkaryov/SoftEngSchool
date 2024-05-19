package get_emails

import (
	"database/sql"
	"fmt"
)

// SelectEmailsFromDB executes a SELECT query and returns a slice of emails
func SelectEmailsFromDB(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT email FROM subscriptions")
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	var emails []string
	for rows.Next() {
		var email string
		err = rows.Scan(&email)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		emails = append(emails, email)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return emails, nil
}
