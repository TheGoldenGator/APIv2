package twitch

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/thegoldengator/APIv2/internal/database"
	"github.com/thegoldengator/APIv2/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ManyUserColors struct {
	Data []UserColorResponse `json:"data"`
}

type UserColorResponse struct {
	UserID    string `json:"user_id"`
	UserLogin string `json:"user_login"`
	UserName  string `json:"user_name"`
	Color     string `json:"color"`
}

type UserColor struct {
	UserID string `json:"user_id"`
	Color  string `json:"color"`
}

// This is only temp to init the color for all members
func (t Twitch) SetColors() error {
	// Get all IDs from members
	ctx := context.Background()
	cursor, err := database.Mongo.Members.Find(ctx, bson.M{}, options.Find().SetProjection(bson.D{{Key: "twitch_id", Value: 1}}))
	if err != nil {
		return err
	}

	var results []bson.D
	if err = cursor.All(ctx, &results); err != nil {
		return err
	}

	var ids []string
	for _, d := range results {
		ids = append(ids, d.Map()["twitch_id"].(string))
	}

	// Get colors from the array of IDs
	colors, err := t.GetColors(ids)
	if err != nil {
		return err
	}

	// Insert color field right before pfp
	for _, uc := range colors {
		database.Mongo.Members.UpdateOne(ctx, bson.M{"twitch_id": uc.UserID}, bson.D{{Key: "$set", Value: bson.D{{Key: "color", Value: uc.Color}}}})
	}

	fmt.Println(colors)
	return nil
}

func (t Twitch) GetColors(ids []string) ([]UserColor, error) {
	// Split up IDs to split up requests
	chunked := utils.Chunk(ids, 100)

	var toSend []UserColor
	for _, v := range chunked {
		body, err := t.makeRequest("GET", "/chat/color?user_id="+strings.Join(v, "&user_id="))
		if err != nil {
			return nil, err
		}

		var colors ManyUserColors
		if err := json.Unmarshal(body, &colors); err != nil {
			if string(body) == `""` {
				return nil, nil
			}

			return nil, err
		}

		for _, ucr := range colors.Data {
			toSend = append(toSend, UserColor{
				UserID: ucr.UserID,
				Color:  ucr.Color,
			})
		}
	}
	return toSend, nil
}
