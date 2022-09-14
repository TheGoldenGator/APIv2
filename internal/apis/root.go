package apis

import (
	"github.com/thegoldengator/APIv2/internal/apis/twitch"
	"github.com/thegoldengator/APIv2/internal/apis/vrchat"
)

var Twitch *twitch.Twitch
var VRChat *vrchat.VRC

func init() {
	Twitch = new(twitch.Twitch)
	VRChat = new(vrchat.VRC)
}
