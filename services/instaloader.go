package services

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const ytdlp = "yt-dlp"

func GetReel(shortcode string) error {
	err := executeCMD(shortcode)
	return err
}

func executeCMD(shortcode string) error {
	args := getScriptArgs(shortcode)
	cmd := exec.Command(ytdlp, args...)

	outputDir := filepath.Join(tempDir, shortcode)
	os.MkdirAll(outputDir, 0755)

	res, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing yt-dlp: %v\nOutput: %s", err, res)
		return err
	}
	return nil
}

func getScriptArgs(shortcode string) []string {
	url := fmt.Sprintf("https://www.instagram.com/reel/%s/", shortcode)
	outputPath := fmt.Sprintf("%s/%s/%%(title)s.%%(ext)s", tempDir, shortcode)

	return []string{
		url,
		"-o", outputPath,
		"--write-description",
		"--no-warnings",
		"--quiet",
		"--no-playlist",
		"--cookies", "./cookies.txt",
		"--format", "bestvideo[height<=1280][width<=720][vcodec^=avc1]+bestaudio/best[height<=1280][width<=720][vcodec^=avc1]/best",
	}
}
