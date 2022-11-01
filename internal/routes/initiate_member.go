package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/thegoldengator/APIv2/internal/apis"
	"github.com/thegoldengator/APIv2/internal/config"
	"github.com/thegoldengator/APIv2/internal/database"
	"github.com/thegoldengator/APIv2/pkg/structures"
)

type TwitchEventSubBody struct {
	Type      string                      `json:"type"`
	Version   string                      `json:"version"`
	Condition TwitchEventSubBodyCondition `json:"condition"`
	Transport TwitchEventSubBodyTransport `json:"transport"`
}

type TwitchEventSubBodyCondition struct {
	BroadcasterUserID string `json:"broadcaster_user_id"`
}

type TwitchEventSubBodyTransport struct {
	Method   string `json:"method"`
	Callback string `json:"callback"`
	Secret   string `json:"secret"`
}

func InitiateMember(w http.ResponseWriter, r *http.Request) {
	var body structures.MemberPost
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return
	}

	// Fetch user from Twitch
	user, err := apis.Twitch.GetUser(body.TwitchID)
	if err != nil {
		return
	}

	// Create the new member in the database
	var member interface{} = structures.Member{
		TwitchID:    user.ID,
		Login:       user.Login,
		DisplayName: user.DisplayName,
		Color:       "#fff",
		Pfp:         user.ProfileImageURL,
		Links:       body.Links,
	}
	_, errDb := database.Mongo.Members.InsertOne(context.Background(), member)
	if errDb != nil {
		return
	}

	// Create the stream document

	// Initiate the EventSub events `stream.online`, `stream.offline`, and `channel.update`
	var events = []string{"stream.online", "stream.offline", "channel.update"}
	for _, v := range events {
		var postBody TwitchEventSubBody = TwitchEventSubBody{
			Type:    v,
			Version: "1",
			Condition: TwitchEventSubBodyCondition{
				BroadcasterUserID: user.ID,
			},
			Transport: TwitchEventSubBodyTransport{
				Method:   "webhook",
				Callback: config.Config.GetString("twitch_eventsub_callback"),
				Secret:   config.Config.GetString("twitch_eventsub_secret"),
			},
		}

		postBodyStr, _ := json.Marshal(postBody)
		respBody := bytes.NewBuffer(postBodyStr)
		resp, err := http.Post("https://api.twitch.tv/helix/eventsub/subscriptions", "application/json", respBody)
		if err != nil {
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		fmt.Println(string(body))
	}
}
