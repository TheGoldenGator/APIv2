package events

import (
	"context"
	"fmt"

	"github.com/thegoldengator/APIv2/internal/apis/twitch"
	"github.com/thegoldengator/APIv2/internal/database"
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

	fmt.Printf("[CHANNEL.UPDATE] Stream changed for %v: %v [%v:%v] changed: %v", event.BroadcasterUserLogin, event.Title, event.CategoryName, event.CategoryID, result.ModifiedCount)
	return nil
}
