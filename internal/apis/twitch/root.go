package twitch

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/thegoldengator/APIv2/internal/config"
)

type Twitch struct{}

// HTTP client to make requests
func (t Twitch) httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

var baseUrl string = "https://api.twitch.tv/helix"

func (t Twitch) makeRequest(method, endpoint string) ([]byte, error) {
	req, _ := http.NewRequest(method, fmt.Sprintf("%v%v", baseUrl, endpoint), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", config.Config.GetString("twitch_client_token")))
	req.Header.Add("Client-Id", config.Config.GetString("twitch_client_id"))

	c := t.httpClient()
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body, nil
}
