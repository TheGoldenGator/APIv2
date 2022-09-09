package twitch

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Fetches a Twitch user
func (t Twitch) GetUser(id string) (*User, error) {
	body, err := t.makeRequest("GET", fmt.Sprintf("/users?id=%v", id))
	if err != nil {
		return nil, err
	}

	var users ManyUsers
	if err := json.Unmarshal(body, &users); err != nil {
		return nil, err
	}

	if len(users.Users) == 0 {
		return nil, errors.New("user is banned or their name changed")
	}

	return &users.Users[0], nil
}

func (t Twitch) GetUsers(ids []string) ([]User, error) {
	fmt.Println(ids)
	body, err := t.makeRequest("GET", "/users?id="+strings.Join(ids, "&id="))
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	var members ManyUsers
	if err := json.Unmarshal(body, &members); err != nil {
		if string(body) == `""` {
			return nil, errors.New("no body returned from Twitch")
		}
	}

	return members.Users, nil
}
