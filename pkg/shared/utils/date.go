package utils

import "time"

// GetDateToString func
func GetDateToString(date *time.Time) string {
	if date.IsZero() {
		return ""
	}
	return date.Format(FormatTime)
}
