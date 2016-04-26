package soundcloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	grantType = "password"
	tokenPath = "/oauth2/token"
)

type Oauth struct {
	ClientId     string
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`

	clientSecret string
	username     string
	password     string
}

func (oauth *Oauth) Authenticate(baseURL *url.URL) error {
	payload := url.Values{
		"grant_type":    {grantType},
		"client_id":     {oauth.ClientId},
		"client_secret": {oauth.clientSecret},
		"username":      {oauth.username},
		"password":      {oauth.password},
	}

	rel, err := url.Parse(tokenPath)
	if err != nil {
		return err
	}
	u := baseURL.ResolveReference(rel)

	body := strings.NewReader(payload.Encode())

	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return fmt.Errorf("Error creating authenitcation request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Error sending authentication request: %v", err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading authentication response bytes: %v", err)
	}

	if err := json.Unmarshal(respBytes, oauth); err != nil {
		return fmt.Errorf("Unable to unmarshal authentication response: %v", err)
	}

	return nil
}
