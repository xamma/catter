package fetcher

import (
	"net/http"
	"net/url"
	"encoding/json"
	"io"
	"fmt"
	"os"
	"path"
)

const (
	Url = "https://api.thecatapi.com/v1/images/search"
)

type ApiResponse struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Width int `json:"width"`
	Height int `json:"height"`
}

func FetchCatImages(limit int) ([]ApiResponse, error) {
	baseURL, err := url.Parse(Url)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	params := url.Values{}
	params.Add("limit", fmt.Sprintf("%d", limit))
	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		return nil, fmt.Errorf("error fetching cat images: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var apiResponses []ApiResponse
	if err := json.Unmarshal(body, &apiResponses); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	if len(apiResponses) == 0 {
		return nil, fmt.Errorf("no images found in response")
	}

	return apiResponses, nil
}

func SaveCatImage(imageUrl, saveDir string) error {
	err := os.MkdirAll(saveDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create folder %s: %w", saveDir, err)
	}

	fileName := path.Base(imageUrl)
	imagePath := path.Join(saveDir, fileName)

	resp, err := http.Get(imageUrl)
	if err != nil {
		return fmt.Errorf("failed to download image: %w", err)
	}
	defer resp.Body.Close()

	file, err := os.Create(imagePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}

	return nil
}