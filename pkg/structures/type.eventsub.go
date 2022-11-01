package structures

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HelixSub struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	UUID      string             `json:"uuid" bson:"uuid"`
	Event     string             `json:"event" bson:"event"`
	TwitchID  string             `json:"twitch_id" bson:"twitch_id"`
	Status    HelixSubStatus     `json:"status" bson:"status"`
	UpdatedAt string             `json:"updated_at" bson:"updated_at"`
	CreatedAt string             `json:"created_at" bson:"created_at"`
}

type HelixSubStatus string

func (hss HelixSubStatus) String() string {
	return string(hss)
}

var (
	HelixSubStatusEnabled                            HelixSubStatus = "enabled"
	HelixSubStatusWebhookCallbackVerificationPending HelixSubStatus = "webhook_callback_verification_pending"
	HelixSubStatusUserRemoved                        HelixSubStatus = "user_removed"
	HelixSubStatusAuthorizationRevoked               HelixSubStatus = "authorization_revoked"
	HelixSubStatusNotificationFailuresExceeded       HelixSubStatus = "notification_failures_exceeded"
)

var HelixSubStatusMap map[string]HelixSubStatus = map[string]HelixSubStatus{
	"enabled":                               HelixSubStatusEnabled,
	"webhook_callback_verification_pending": HelixSubStatusWebhookCallbackVerificationPending,
	"user_removed":                          HelixSubStatusUserRemoved,
	"authorization_revoked":                 HelixSubStatusAuthorizationRevoked,
	"notification_failures_exceeded":        HelixSubStatusNotificationFailuresExceeded,
}

func ResolveHelixSubStatus(str string) (*HelixSubStatus, error) {
	if val, ok := HelixSubStatusMap[str]; ok {
		return &val, nil
	} else {
		return nil, errors.New("invalid helix subscription status")
	}
}
