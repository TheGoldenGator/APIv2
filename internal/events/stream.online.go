package events

import (
	"context"
	"fmt"
	"time"

	"github.com/thegoldengator/APIv2/internal/apis/twitch"
	"github.com/thegoldengator/APIv2/internal/database"
	"github.com/thegoldengator/APIv2/internal/sse"
	"github.com/thegoldengator/APIv2/pkg/structures"
	"go.mongodb.org/mongo-driver/bson"
)

func StreamOnline(event twitch.EventSubStreamOnlineEvent) error {
	_, err := database.Mongo.Stream.UpdateOne(
		context.Background(),
		bson.M{"twitch_id": event.BroadcasterUserID},
		bson.M{"$set": bson.M{"status": "ONLINE", "started_at": event.StartedAt.Format(time.RFC3339)}},
	)

	if err != nil {
		return err
	}

	fmt.Printf("Stream went online for %v", event.BroadcasterUserLogin)

	var member structures.Member
	if err = database.Mongo.Members.FindOne(context.Background(), bson.M{"twitch_id": event.BroadcasterUserID}).Decode(&member); err != nil {
		return err
	}

	var stream structures.Stream
	if err = database.Mongo.Stream.FindOne(context.Background(), bson.M{"twitch_id": event.BroadcasterUserID}).Decode(&stream); err != nil {
		return err
	}

	// Alert SSE
	var doc interface{} = event
	sse.PublishMessage(sse.SSEChannelEvents, sse.SSEMessage{Event: sse.SSEMessageEventStreamOnline, Member: &member, Stream: &stream, Data: doc})
	return nil
}
