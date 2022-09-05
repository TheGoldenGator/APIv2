package routines

import (
	"context"
	"errors"
	"fmt"

	"github.com/thegoldengator/APIv2/internal/apis"
	"github.com/thegoldengator/APIv2/internal/database"
	"github.com/thegoldengator/APIv2/pkg/structures"
	"go.mongodb.org/mongo-driver/bson"
)

func Pfp() error {
	ctx := context.Background()

	cursor, err := database.Mongo.Members.Find(ctx, bson.M{})
	if err != nil {
		return err
	}

	var members []structures.Member
	if err = cursor.All(ctx, &members); err != nil {
		return err
	}

	var ids = []string{}
	for _, s := range members {
		ids = append(ids, s.TwitchID)
	}

	helixUsers, err := apis.Twitch.GetUsers(ids)
	if err != nil {
		return err
	}

	if len(helixUsers) == 0 {
		return errors.New("no users")
	}

	for _, s := range helixUsers {
		_, err := database.Mongo.Members.UpdateOne(
			ctx,
			bson.M{"twitch_id": s.ID},
			bson.D{
				{Key: "$set", Value: bson.D{{Key: "pfp", Value: s.ProfileImageURL}}},
			},
		)

		if err != nil {
			return err
		}

		fmt.Printf("Updated pfp for %v\n", s.Login)
	}

	return nil
}
