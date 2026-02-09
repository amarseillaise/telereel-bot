package services

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	tempDir = "./.temp"
)

func GetReelPath(shortcode string) (string, string, error) {
	videoPath, captionPath, err := downloadRemote(shortcode)
	return videoPath, captionPath, err
}

func downloadRemote(shortcode string) (string, string, error) {
	err := DownloadReel(shortcode)
	if err != nil {
		return "", "", err
	}
	captionPath := findFile(shortcode, []string{".description", ".txt"})
	videoPath := findFile(shortcode, []string{".mp4", ".avi", ".mkv", ".mov"})
	return videoPath, captionPath, nil
}

func findFile(shortcode string, extensions []string) string {
	res := ""
	for _, ext := range extensions {
		mathces, err := filepath.Glob(fmt.Sprintf("%s/%s/*%s", tempDir, shortcode, ext))
		if err == nil && len(mathces) > 0 {
			res = mathces[0]
			break
		}
	}
	return res
}

func ParseShortcode(_url string) string {
	pattern := "reel/.+/"
	re := regexp.MustCompile(pattern)
	match := re.FindString(_url)
	resultsSlice := strings.Split(match, "/")
	shortcode := resultsSlice[1]
	return shortcode
}
