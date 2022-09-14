package vrchat

import (
	"fmt"
	"net/http"
)

const (
	BaseEndpoint  = "https://api.vrchat.cloud/api/1"
	LoginEndpoint = BaseEndpoint + "/auth/user"
)

type client struct {
	apiKey    string
	userName  string
	password  string
	authToken string
}

func (c *client) AuthenticateUser() error {
	fmt.Println("[VRC] Logging in...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", LoginEndpoint, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.userName, c.password)

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	auth := ParseCookieValue("auth", res)
	if auth == "" {
		return fmt.Errorf("unable to obtain authentication key, check provided credentials")
	}
	c.authToken = auth

	fmt.Println("[VRC] Logged in.")
	return nil
}

func ParseCookieValue(cookieName string, res *http.Response) string {
	var value string
	for _, cookie := range res.Cookies() {
		if cookie.Name == cookieName {
			value = cookie.Value
		}
	}

	return value
}
