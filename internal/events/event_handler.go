package events

import (
	"context"
	"time"

	"github.com/thegoldengator/APIv2/internal/database"
	"github.com/thegoldengator/APIv2/pkg/structures"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleEvent(event structures.HelixSub) error {
	ctx := context.Background()
	result, err := database.Mongo.Events.UpdateOne(
		ctx,
		bson.M{"twitch_id": event.TwitchID, "event": event.Event},
		bson.M{"$set": bson.M{"status": event.Status.String(), "updated_at": time.Now().Format(time.RFC3339)}},
	)

	if err != nil {
		return err
	}

	// Create if not found
	if result.MatchedCount == 0 {
		var doc interface{} = structures.HelixSub{
			ID:        primitive.NewObjectID(),
			UUID:      event.UUID,
			TwitchID:  event.TwitchID,
			Event:     event.Event,
			Status:    event.Status,
			UpdatedAt: event.CreatedAt,
			CreatedAt: event.CreatedAt,
		}
		if _, err := database.Mongo.Events.InsertOne(ctx, doc); err != nil {
			return err
		}
	}

	return nil
}
