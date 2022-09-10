package sse

import (
	"encoding/json"
)

// Publishes message to SSE
func PublishMessage(channel SSEChannel, message SSEMessage) {
	jsonStr, _ := json.Marshal(message)
	SSEServer.Notifier <- jsonStr
}

func PublishPing(channel SSEChannel) {
	msg := SSEMessage{
		Event:  "ping",
		Member: nil,
		Stream: nil,
		Data:   nil,
	}
	jsonStr, _ := json.Marshal(msg)
	SSEServer.Notifier <- jsonStr
}
