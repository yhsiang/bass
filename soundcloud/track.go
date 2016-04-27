package soundcloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
)

type TrackService interface {
	Get(string) (*Track, error)
	GetAll() ([]Track, error)
	Create(map[string]string, string) (*Track, error)
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

func (s *TrackOp) Create(params map[string]string, filePath string) (*Track, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("track[asset_data]", filepath.Base(filePath))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewFormRequest(path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())

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
