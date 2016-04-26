package soundcloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "net/http"s
)

type TrackService interface {
	Get() (*[]Track, error)
}

type TrackOp struct {
	client *Client
}

type Track struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Label       string `json:"label"`
}

func (s *TrackOp) Get() (*[]Track, error) {
	path := "tracks"

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tracks = &[]Track{}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading authentication response bytes: %v", err)
	}
	var marshalErr error
	marshalErr = json.Unmarshal(respBytes, tracks)
	if marshalErr != nil {
		return nil, fmt.Errorf("Error reading authentication response bytes: %v", marshalErr)
	}

	return tracks, nil
}
