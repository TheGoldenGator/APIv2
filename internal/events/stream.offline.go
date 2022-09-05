package events

import (
	"context"
	"fmt"

	"github.com/thegoldengator/APIv2/internal/apis/twitch"
	"github.com/thegoldengator/APIv2/internal/database"
	"github.com/thegoldengator/APIv2/internal/sse"
	"github.com/thegoldengator/APIv2/pkg/structures"
	"go.mongodb.org/mongo-driver/bson"
)

// Changes MongoDB status for streamer to offline.
func StreamOffline(event twitch.EventSubStreamOfflineEvent) error {
	_, err := database.Mongo.Stream.UpdateOne(
		context.Background(),
		bson.M{"twitch_id": event.BroadcasterUserID},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: "status", Value: "OFFLINE"}, {Key: "viewers", Value: "1"}}},
		},
	)

	if err != nil {
		return err
	}

	fmt.Printf("Stream went offline for %v", event.BroadcasterUserLogin)

	var member structures.Member
	if err = database.Mongo.Members.FindOne(context.Background(), bson.M{"twitch_id": event.BroadcasterUserID}).Decode(&member); err != nil {
		return err
	}

	// Alert SSE
	var doc interface{} = event
	sse.PublishMessage(sse.SSEChannelEvents, sse.SSEMessage{Event: sse.SSEMessageEventStreamOffline, Member: &member, Data: doc})
	return nil
}
