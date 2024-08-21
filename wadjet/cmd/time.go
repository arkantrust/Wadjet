package cmd

import (
	"time"
)

// Time returns the server's local time.
func Time() (string, error) {
	return time.Now().Format(time.RFC3339), nil
}
