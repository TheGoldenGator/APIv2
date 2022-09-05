package events

import "time"

func GetRFCTimestamp() string {
	now := time.Now()
	return now.Format(time.RFC3339)
}
