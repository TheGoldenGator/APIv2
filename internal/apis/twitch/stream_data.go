package twitch

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (t Twitch) GetStreamData(id string) (*Stream, error) {
	body, err := t.makeRequest("GET", fmt.Sprintf("/streams?user_id=%v", id))
	if err != nil {
		return nil, err
	}

	var streams ManyStreams
	if err := json.Unmarshal(body, &streams); err != nil {
		return nil, err
	}

	if len(streams.Streams) == 0 {
		return nil, errors.New("no streams found")
	}

	return &streams.Streams[0], nil
}
