package structures

import "go.mongodb.org/mongo-driver/bson/primitive"

type Stream struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	TwitchID  string             `json:"twitch_id" bson:"twitch_id"`
	Login     string             `json:"login" bson:"login"`
	Status    StreamStatus       `json:"status" bson:"status"`
	Title     string             `json:"title" bson:"title"`
	GameID    string             `json:"game_id" bson:"game_id"`
	Game      string             `json:"game" bson:"game"`
	Viewers   int                `json:"viewers" bson:"viewers"`
	Thumbnail string             `json:"thumbnail" bson:"thumbnail"`
	StartedAt string             `json:"started_at" bson:"started_at"`
}

type StreamStatus string

const (
	StreamStatusOnline  StreamStatus = "ONLINE"
	StreamStatusOffline StreamStatus = "OFFLINE"
	StreamStatusAll     StreamStatus = "ALL"
)

var AllStreamStatus = []StreamStatus{
	StreamStatusOnline,
	StreamStatusOffline,
}

func (s StreamStatus) String() string {
	return string(s)
}
