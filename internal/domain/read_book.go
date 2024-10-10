package domain

import (
	"time"
)

type ReadBook struct {
	ID              string     `json:"id,omitempty"`
	BookID          string     `json:"book_id" validate:"required"`
	StartDate       time.Time  `json:"start_date" validate:"required"`
	ExpectedEndDate time.Time  `json:"expected_end_date"`
	ActualEndDate   *time.Time `json:"actual_end_date,omitempty"`
	Comments        []string   `json:"comments,omitempty"`
	Rating          *int       `json:"rating,omitempty"`
}
