package routines

import (
	"context"
	"fmt"

	"github.com/thegoldengator/APIv2/internal/apis"
	"github.com/thegoldengator/APIv2/internal/database"
	"github.com/thegoldengator/APIv2/internal/sse"
	"github.com/thegoldengator/APIv2/pkg/structures"
	"go.mongodb.org/mongo-driver/bson"
)

func ViewCount() error {
	ctx := context.Background()

	cursor, err := database.Mongo.Stream.Find(ctx, bson.M{"status": "ONLINE"})
	if err != nil {
		return err
	}

	var streams []structures.Stream
	if err = cursor.All(ctx, &streams); err != nil {
		return err
	}

	var ids = []string{}
	for _, s := range streams {
		ids = append(ids, s.TwitchID)
	}

	helixStreams, err := apis.Twitch.GetStreams(ids)
	if err != nil {
		return err
	}

	var toSend map[string]int = make(map[string]int)
	for _, s := range helixStreams {
		_, err := database.Mongo.Stream.UpdateOne(
			ctx,
			bson.M{"twitch_id": s.ID},
			bson.D{
				{Key: "$set", Value: bson.D{{Key: "viewers", Value: s.ViewerCount}}},
			},
		)

		if err != nil {
			return err
		}

		toSend[s.UserID] = s.ViewerCount
		fmt.Printf("Updated viewer count for %v - %v\n", s.UserLogin, s.ViewerCount)
	}

	// Alert SSE
	var doc interface{} = toSend
	fmt.Println(doc)
	sse.PublishMessage(sse.SSEChannelEvents, sse.SSEMessage{
		Event: sse.SSEMessageEventViewers,
		Data:  doc,
	})

	return nil
}
