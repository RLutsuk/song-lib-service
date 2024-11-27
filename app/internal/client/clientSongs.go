package clientSongs

import (
	"effective_project/app/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client interface {
	FetchSongInfo(group, title string) (*models.Song, error)
}

type songClient struct {
	baseURL string
}

func NewSongClient(baseURL string) Client {
	return &songClient{
		baseURL: baseURL,
	}
}

func (c *songClient) FetchSongInfo(group, title string) (*models.Song, error) {
	params := url.Values{}
	params.Add("group", group)
	params.Add("song", title)

	fullURL := fmt.Sprintf("%s/info?%s", c.baseURL, params.Encode())
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to /info")
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return nil, fmt.Errorf("bad request")
	}

	if resp.StatusCode == http.StatusInternalServerError {
		return nil, fmt.Errorf("internal server error")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var songInfo models.Song
	if err := json.Unmarshal(body, &songInfo); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return &songInfo, nil
}
