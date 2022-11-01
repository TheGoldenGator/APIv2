package routes

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/thegoldengator/APIv2/internal/apis/twitch"
	"github.com/thegoldengator/APIv2/internal/config"
	"github.com/thegoldengator/APIv2/internal/events"
	"github.com/thegoldengator/APIv2/pkg/structures"
)

// Verify message from EventSub
func VerifyEventSubNotification(secret string, header http.Header, message string) bool {
	hmacMessage := []byte(fmt.Sprintf("%s%s%s", header.Get("Twitch-Eventsub-Message-Id"), header.Get("Twitch-Eventsub-Message-Timestamp"), message))
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(hmacMessage)
	hmacsha256 := fmt.Sprintf("sha256=%s", hex.EncodeToString(mac.Sum(nil)))
	return hmacsha256 == header.Get("Twitch-Eventsub-Message-Signature")
}

type eventsubNotification struct {
	Subscription twitch.EventSubSubscription `json:"subscription"`
	Challenge    string                      `json:"challenge"`
	Event        json.RawMessage             `json:"event"`
}

// Route that fetches POSTed eventsub notifications from Twitch
func EventsubRecievedNotification(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	defer r.Body.Close()

	// Verify Twitch sent the event
	if !VerifyEventSubNotification(config.Config.GetString("twitch_eventsub_secret"), r.Header, string(body)) {
		log.Println("No valid signature on subscription")
		return
	} else {
		log.Println("Verified signature on subscription")
	}
	var vals eventsubNotification
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&vals)
	if err != nil {
		log.Println(err)
		return
	}

	twitchEventType := r.Header.Get("Twitch-Eventsub-Message-Type")

	switch twitchEventType {
	case "notification":
		eventType := bytes.NewBuffer([]byte(vals.Subscription.Type)).String()
		switch {
		case eventType == "stream.online":
			var streamOnline twitch.EventSubStreamOnlineEvent
			err := json.NewDecoder(bytes.NewReader(vals.Event)).Decode(&streamOnline)
			if err != nil {
				panic(err.Error())
			}

			errDb := events.StreamOnline(streamOnline)
			if errDb != nil {
				panic(err.Error())
			}

		case eventType == "stream.offline":
			var streamOffline twitch.EventSubStreamOfflineEvent
			err := json.NewDecoder(bytes.NewReader(vals.Event)).Decode(&streamOffline)
			if err != nil {
				panic(err.Error())
			}

			errDb := events.StreamOffline(streamOffline)
			if errDb != nil {
				panic(err.Error())
			}

		case eventType == "channel.update":
			var streamUpdate twitch.EventSubChannelUpdateEvent
			err := json.NewDecoder(bytes.NewReader(vals.Event)).Decode(&streamUpdate)
			if err != nil {
				panic(err.Error())
			}

			errDb := events.ChannelUpdate(streamUpdate)
			if errDb != nil {
				panic(err.Error())
			}
		}

		// if there's a challenge in the request, respond with only the challenge to verify your eventsub.
	case "webhook_callback_verification":
		// Since we're verifiying it, create the subscription if it doesn't already exist
		events.HandleEvent(structures.HelixSub{
			UUID:      vals.Subscription.ID,
			Event:     vals.Subscription.Type,
			TwitchID:  vals.Subscription.Condition.BroadcasterUserID,
			CreatedAt: vals.Subscription.CreatedAt.Format(time.RFC3339),
			Status:    structures.HelixSubStatusEnabled,
		})
		w.Write([]byte(vals.Challenge))
		return

	case "revocation":
		switch vals.Subscription.Status {
		case "user_removed":
			// User mentioned in subscription doesn't exist anymore. The notification status is set to user_removed
			events.HandleEvent(structures.HelixSub{
				UUID:      vals.Subscription.ID,
				Event:     vals.Subscription.Type,
				TwitchID:  vals.Subscription.Condition.BroadcasterUserID,
				CreatedAt: vals.Subscription.CreatedAt.Format(time.RFC3339),
				Status:    structures.HelixSubStatusUserRemoved,
			})

		case "authorization_revoked":
			// The user revoked the authorization token or simply changed their password. The notification’s status is set to authorization_revoked.
			events.HandleEvent(structures.HelixSub{
				UUID:      vals.Subscription.ID,
				Event:     vals.Subscription.Type,
				TwitchID:  vals.Subscription.Condition.BroadcasterUserID,
				CreatedAt: vals.Subscription.CreatedAt.Format(time.RFC3339),
				Status:    structures.HelixSubStatusAuthorizationRevoked,
			})

		case "notification_failures_exceeded":
			// he callback failed to respond in a timely manner too many times.
			events.HandleEvent(structures.HelixSub{
				UUID:      vals.Subscription.ID,
				Event:     vals.Subscription.Type,
				TwitchID:  vals.Subscription.Condition.BroadcasterUserID,
				CreatedAt: vals.Subscription.CreatedAt.Format(time.RFC3339),
				Status:    structures.HelixSubStatusNotificationFailuresExceeded,
			})
		}
	}

}
