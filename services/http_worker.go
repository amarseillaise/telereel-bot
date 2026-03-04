package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const (
	reel_url = "https://www.instagram.com/reel/"
)

type ApiResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type ReelData struct {
	VideoUrl string
	Caption  string
	Title    string
}

func GetReelData(shortcode string) (ReelData, error) {
	result := ReelData{}
	host_url := os.Getenv("FILE_SERVER_URL")
	api_endpoint := fmt.Sprintf("%s/%s", host_url, shortcode)
	resp, err := downloadRemote(api_endpoint)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	result.VideoUrl = getReelVideoUrl(api_endpoint)
	result.Caption = getReelCaption(api_endpoint)
	result.Title = cutTitleIfNeeded(result.Caption)
	appendSourceUrlToCaption(&result.Caption, reel_url+shortcode)
	cutCaptionIfNeed(&result.Caption)
	return result, nil
}

func downloadRemote(url_path string) (*http.Response, error) {
	resp, err := http.Post(url_path, "", nil)
	return resp, err
}

func getReelVideoUrl(api_endpoint string) string {
	videoUrl := fmt.Sprintf("%s/video.mp4", api_endpoint)
	return videoUrl
}

func getReelCaption(api_endpoint string) string {
	methodUrl := fmt.Sprintf("%s/description", api_endpoint)
	resp, err := http.Get(methodUrl)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return ""
	}
	captionData := make([]byte, resp.ContentLength)
	_, err = resp.Body.Read(captionData)
	if err != nil {
		return ""
	}
	response := ApiResponse{}
	err = json.Unmarshal(captionData, &response)
	if err != nil || !response.Success {
		return ""
	}
	return response.Message
}

func cutTitleIfNeeded(caption string) string {
	maxTitleLength := 100
	if len(caption) > maxTitleLength {
		return (caption)[:maxTitleLength] + "..."
	}
	return caption
}

func appendSourceUrlToCaption(caption *string, url string) {
	if len(*caption) > 0 {
		url += "\n\n"
	}
	*caption = url + *caption
}

func cutCaptionIfNeed(caption *string) {
	if len(*caption) >= 1024 {
		*caption = (*caption)[:1023]
	}
}

func ParseShortcode(_url string) string {
	pattern := "reel/.+/"
	re := regexp.MustCompile(pattern)
	match := re.FindString(_url)
	if match == "" {
		return ""
	}
	resultsSlice := strings.Split(match, "/")
	shortcode := resultsSlice[1]
	return shortcode
}
