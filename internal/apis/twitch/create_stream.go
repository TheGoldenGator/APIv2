package twitch

import (
	"context"
	"fmt"
	"time"

	"github.com/thegoldengator/APIv2/internal/database"
	"github.com/thegoldengator/APIv2/pkg/structures"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (t Twitch) CreateStreams() ([]*structures.Member, error) {
	cursor, err := database.Mongo.Members.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var streamers []*structures.Member
	if err = cursor.All(context.Background(), &streamers); err != nil {
		return nil, err
	}

	var docs []interface{}
	for _, m := range streamers {
		var strim structures.Stream
		if err := database.Mongo.Stream.FindOne(context.Background(), bson.M{"twitch_id": m.TwitchID}).Decode(&strim); err != nil {
			if err == mongo.ErrNoDocuments {
				// No document found for member, so create the stream with empty data
				member, err := t.GetUser(m.TwitchID)
				if err != nil {
					member = &User{
						ID:              "525322083",
						Login:           "jouffa",
						DisplayName:     "jouffa",
						Type:            "live",
						BroadcasterType: "affiliate",
						Description:     "",
						ProfileImageURL: "https://static-cdn.jtvnw.net/jtv_user_pictures/cf778454-20ee-440c-96a9-702caa5d0beb-profile_image-600x600.png",
						OfflineImageURL: "",
						ViewCount:       1,
						Email:           "",
						CreatedAt:       time.Now(),
					}
					docs = append(docs, *member)
				}

				stream, err := t.GetStreamData(m.TwitchID)
				fmt.Println(err)
				if err != nil {
					// If there's an error that means streamer is offline
					// No *legal* way for me to get previous data so I put "N/A" for data that can't be fetched currently
					var toInsert interface{} = structures.Stream{
						ID:        primitive.NewObjectID(),
						TwitchID:  member.ID,
						Login:     member.Login,
						Status:    structures.StreamStatusOffline,
						Title:     "N/A",
						GameID:    "0",
						Game:      "N/A",
						Viewers:   1,
						Thumbnail: fmt.Sprintf("https://static-cdn.jtvnw.net/previews-ttv/live_user_%s-{width}x{height}.jpg", member.Login),
						StartedAt: time.Now().Format(time.RFC3339),
					}
					fmt.Println(member.Login + " is offline: inserting empty data")
					docs = append(docs, toInsert)
				} else {
					// Insert current stream data since they are online
					var toInsert interface{} = structures.Stream{
						ID:        primitive.NewObjectID(),
						TwitchID:  member.ID,
						Login:     member.Login,
						Status:    structures.StreamStatusOnline,
						Title:     stream.Title,
						GameID:    stream.GameID,
						Game:      stream.GameName,
						Viewers:   stream.ViewerCount,
						Thumbnail: fmt.Sprintf("https://static-cdn.jtvnw.net/previews-ttv/live_user_%s-{width}x{height}.jpg", member.Login),
						StartedAt: stream.StartedAt.Format(time.RFC3339),
					}
					fmt.Println(member.Login + " is online: inserting current stream data")
					docs = append(docs, toInsert)
				}
			}
		}
	}

	_, errInsert := database.Mongo.Stream.InsertMany(context.Background(), docs)
	if errInsert != nil {
		return nil, errInsert
	}

	return nil, nil
}
