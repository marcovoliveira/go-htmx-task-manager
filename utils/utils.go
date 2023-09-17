package utils

import (
	"fmt"
	"time"
)

// ParseDate parses a date string in the format "YYYY-MM-DD" and returns a time.Time value.
func ParseDate(dateStr string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func ParseInt(s string) int {
	i := 0
	fmt.Sscanf(s, "%d", &i)
	return i
}

// IsValidDate checks if a date string is in the valid "YYYY-MM-DD" format.
func IsValidDate(dateStr string) bool {
	_, err := ParseDate(dateStr)
	return err == nil
}

// ValidateID checks if an ID is valid (greater than zero).
func ValidateID(id int) error {
	if id <= 0 {
		return fmt.Errorf("ID must be greater than zero")
	}
	return nil
}
