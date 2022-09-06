package twitch

import (
	"encoding/json"
	"strings"
)

func (t Twitch) GetStreams(ids []string) ([]Stream, error) {
	body, err := t.makeRequest("GET", "/streams?user_id="+strings.Join(ids, "&user_id="))
	if err != nil {
		return nil, err
	}

	var streams ManyStreams
	if err := json.Unmarshal(body, &streams); err != nil {
		if string(body) == `""` {
			return nil, nil
		}

		return nil, err
	}

	return streams.Streams, nil
}
