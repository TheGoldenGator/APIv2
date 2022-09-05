package apis

import (
	"github.com/thegoldengator/APIv2/internal/apis/twitch"
)

var Twitch *twitch.Twitch

func init() {
	Twitch = new(twitch.Twitch)
}
