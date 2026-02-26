package bot

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	re "regexp"
	"time"

	tele "gopkg.in/telebot.v4"
)

type ReelInfo struct {
	Video   *tele.Video
	Caption string
}

func InitBot(token *string) (*tele.Bot, error) {
	client := initClient()
	pref := tele.Settings{
		Token:  *token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		Client: client,
	}
	bot, err := tele.NewBot(pref)
	return bot, err
}

func initClient() *http.Client {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	client := &http.Client{
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				// force IPv4
				return dialer.DialContext(ctx, "tcp4", addr)
			},
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
			DisableKeepAlives: false,
		},
	}
	return client
}

func MakeVideo(videoUrl string) *tele.Video {
	teleVideo := &tele.Video{File: tele.FromURL(videoUrl), MIME: "video/mp4"}
	return teleVideo
}

func IsValidInstagramReelURL(url string) bool {
	pattern := "\\.*instagram.com/reel\\.*/"
	is_valid, _ := re.MatchString(pattern, url)
	return is_valid
}
