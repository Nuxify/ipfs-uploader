package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type IPFSRequest struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

var (
	client *http.Client = &http.Client{Timeout: 60 * time.Second}
)

// Post is a http post request helper
func Post(url string, payload []IPFSRequest, response interface{}) error {
	// convert to reader
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		return err
	}

	// add custom headers
	req.Header.Add("X-API-KEY", os.Getenv("MORALIS_API_KEY"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}
	return nil
}
