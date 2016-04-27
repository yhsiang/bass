package soundcloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "net/http"s
)

type TrackService interface {
	Get(string) (*Track, error)
	GetAll() ([]Track, error)
}

type TrackOp struct {
	client *Client
}

type Track struct {
	Id          uint64 `json:"id"`
	UserId      uint64 `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Label       string `json:"label"`
}

var (
	path = "tracks"
)

func (s *TrackOp) Get(id string) (*Track, error) {
	path := path + "/" + id

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var track = &Track{}
	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("Error reading authentication response bytes: %v", err)
	}

	var marshalErr error
	marshalErr = json.Unmarshal(respBytes, track)
	if marshalErr != nil {
		return nil, fmt.Errorf("Error reading authentication response bytes: %v", marshalErr)
	}

	return track, nil
}

func (s *TrackOp) GetAll() ([]Track, error) {
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tracks []Track
	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("Error reading authentication response bytes: %v", err)
	}

	var marshalErr error
	marshalErr = json.Unmarshal(respBytes, &tracks)
	if marshalErr != nil {
		return nil, fmt.Errorf("Error reading authentication response bytes: %v", marshalErr)
	}

	return tracks, nil
}
