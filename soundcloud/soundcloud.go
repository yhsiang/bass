package soundcloud

import (
	"bytes"
	"encoding/json"
	// "fmt"
	"net/http"
	"net/url"
	// "io/ioutil"s
)

const (
	defaultBaseURL = "https://api.soundcloud.com/"
	mediaType      = "application/json"
)

type Client struct {
	client  *http.Client
	oauth   *Oauth
	BaseURL *url.URL

	//UserAgent
	Track TrackService
}

func NewClient(
	httpClient *http.Client,
	clientId string,
	clientSecret string,
	username string,
	password string,
) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	oauth := &Oauth{
		ClientId:     clientId,
		clientSecret: clientSecret,
		username:     username,
		password:     password,
	}
	err := oauth.Authenticate(baseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client:  httpClient,
		oauth:   oauth,
		BaseURL: baseURL,
	}

	c.Track = &TrackOp{client: c}

	return c, nil
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)
	if method == "GET" {
		q := u.Query()
		q.Set("client_id", c.oauth.ClientId)
		u.RawQuery = q.Encode()
	}

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	// fmt.Println(u.String())
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)

	return req, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
