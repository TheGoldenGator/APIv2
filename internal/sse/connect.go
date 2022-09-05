package sse

import (
	"time"

	"github.com/r3labs/sse/v2"
	"github.com/thegoldengator/APIv2/pkg/structures"
)

type SSEChannel string
type SSEMessageEvent string

type SSEMessage struct {
	Event  SSEMessageEvent    `json:"event"`
	Member *structures.Member `json:"member"` // Member associated with the event that's in database
	Stream *structures.Stream `json:"stream"` // Stream is only there when `stream.online` and `channel.update` is fired
	Data   interface{}        `json:"data"`   // Data from the event
}

var (
	Server *sse.Server

	// Channels
	SSEChannelEvents SSEChannel = "events"

	// Events
	SSEMessageEventStreamOnline  SSEMessageEvent = "stream.online"
	SSEMessageEventStreamOffline SSEMessageEvent = "stream.offline"
	SSEMessageEventChannelUpdate SSEMessageEvent = "channel.update"
	SSEMessageEventViewers       SSEMessageEvent = "viewers"
)

func (s SSEChannel) String() string {
	return string(s)
}

func Connect() {
	Server = sse.New()
	Server.EventTTL = time.Second * 1
	Server.CreateStream(SSEChannelEvents.String())
}
