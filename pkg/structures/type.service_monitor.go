package structures

type SSEStats struct {
	Connections int `json:"connections" bson:"connections"`
}

type RateLimits struct {
	TwitchRatelimit          string `json:"twitch_ratelimit" bson:"twitch_ratelimit"`
	TwitchRatelimitRemaining string `json:"twitch_ratelimit_remaining" bson:"twitch_ratelimit_remaining"`
	TwitchRatelimitReset     string `json:"twitch_ratelimit_reset" bson:"twitch_ratelimit_reset"`
}
