package sse

import (
	"encoding/json"

	"github.com/r3labs/sse/v2"
)

// Publishes message to SSE
func PublishMessage(channel SSEChannel, message SSEMessage) {
	jsonStr, _ := json.Marshal(message)
	Server.Publish(channel.String(), &sse.Event{
		Data: jsonStr,
	})
}

func PublishPing(channel SSEChannel) {
	msg := SSEMessage{
		Event:  "ping",
		Member: nil,
		Stream: nil,
		Data:   nil,
	}
	jsonStr, _ := json.Marshal(msg)
	Server.Publish(channel.String(), &sse.Event{
		Data: jsonStr,
	})
}
