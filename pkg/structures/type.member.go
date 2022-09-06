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
	Links       *MemberLink        `json:"links" bson:"links"`
}

type MemberLink struct {
	Twitch        string `json:"twitch" bson:"twitch"`
	Reddit        string `json:"reddit" bson:"reddit"`
	Instagram     string `json:"instagram" bson:"instagram"`
	Twitter       string `json:"twitter" bson:"twitter"`
	Discord       string `json:"discord" bson:"discord"`
	Youtube       string `json:"youtube" bson:"youtube"`
	Tiktok        string `json:"tiktok" bson:"tiktok"`
	VrchatLegends string `json:"vrchat_legends" bson:"vrchat_legends"`
}

type MemberPost struct {
	TwitchID string     `json:"twitch_id" bson:"twitch_id"`
	Links    MemberLink `json:"links" bson:"links"`
}
