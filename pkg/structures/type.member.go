package structures

import "go.mongodb.org/mongo-driver/bson/primitive"

type Member struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	TwitchID    string             `json:"twitch_id" bson:"twitch_id"`
	Login       string             `json:"login" bson:"login"`
	DisplayName string             `json:"display_name" bson:"display_name"`
	Color       string             `json:"color" bson:"color"`
	Pfp         string             `json:"pfp" bson:"pfp"`
	PfpSevenTV  string             `json:"pfp_seventv" bson:"pfp_seventv"`
	Links       []string           `json:"links" bson:"links"`
}

type MemberPost struct {
	TwitchID string   `json:"twitch_id" bson:"twitch_id"`
	Links    []string `json:"links" bson:"links"`
}
