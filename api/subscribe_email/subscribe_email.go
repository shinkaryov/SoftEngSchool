package subscribe_email

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
)

// Insert Email into the database
func InsertEmail(db *sql.DB, email string) error {
	query := "INSERT INTO subscriptions  VALUES ($1, 1)"
	_, err := db.Exec(query, email)
	if err != nil {
		// Check if the error is a pq.Error
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return fmt.Errorf("already subscribed")
			}
		}
		return fmt.Errorf("error executing insert: %w", err)
	}
	return nil
}
