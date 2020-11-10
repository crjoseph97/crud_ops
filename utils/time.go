package utils

import (
	"time"
)

// StringToTime converts time of string in format into to time in toFormat
func StringToTime(timeStr string, format string, toFormat string) (time.Time, error) {
	date, err := time.Parse(format, timeStr)
	if err != nil {
		return time.Time{}, err
	}

	date, err = time.Parse(toFormat, date.Format(toFormat))
	if err != nil {
		return time.Time{}, err
	}

	return date, err
}
