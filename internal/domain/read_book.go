package domain

import "time"

// ReadBook represents a record of a book being read by a user.
type ReadBook struct {
	ID              string    `json:"id"`
	BookID          string    `json:"book_id"`
	StartDate       time.Time `json:"start_date"`
	ExpectedEndDate time.Time `json:"expected_end_date"`
	ActualEndDate   time.Time `json:"actual_end_date"`
	Comments        string    `json:"comments"`
	Rating          int       `json:"rating"`
}
