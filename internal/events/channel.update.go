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

func ChannelUpdate(event twitch.EventSubChannelUpdateEvent) error {
	result, err := database.Mongo.Stream.UpdateOne(
		context.Background(),
		bson.M{"twitch_id": event.BroadcasterUserID},
		bson.M{"$set": bson.M{"title": event.Title, "game": event.CategoryName, "game_id": event.CategoryID}},
	)

	if err != nil {
		return err
	}

	var member structures.Member
	if err = database.Mongo.Members.FindOne(context.Background(), bson.M{"twitch_id": event.BroadcasterUserID}).Decode(&member); err != nil {
		return err
	}

	// Alert SSE
	var doc interface{} = event
	sse.PublishMessage(sse.SSEChannelEvents, sse.SSEMessage{Event: sse.SSEMessageEventChannelUpdate, Member: &member, Data: doc})
	fmt.Printf("[CHANNEL.UPDATE] Stream changed for %v: %v [%v:%v] changed: %v", event.BroadcasterUserLogin, event.Title, event.CategoryName, event.CategoryID, result.ModifiedCount)
	return nil
}
