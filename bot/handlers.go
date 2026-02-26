package bot

import (
	"fmt"

	"gopkg.in/telebot.v4"

	"github.com/amarseillaise/instareels_to_telegram/services"
)

func OnTextHandler(c telebot.Context) error {
	_url := c.Text()
	if IsValidInstagramReelURL(_url) {
		c.Notify(telebot.UploadingVideo)
		shortcode := services.ParseShortcode(_url)
		videoUrl, caption, err := services.GetReelData(shortcode)
		if err == nil {
			teleVideo := MakeVideo(videoUrl)
			teleVideo.Caption = caption
			return c.Reply(teleVideo)
		} else {
			return c.Reply(err.Error())
		}
	}
	return nil
}

func OnQueryHandler(c telebot.Context) error {
	_url := c.Query().Text
	fmt.Println(_url)
	if IsValidInstagramReelURL(_url) {
		vr := telebot.VideoResult{
			URL:      "https://www.youtube.com/watch?v=wEc82Yq1uwo",
			MIME:     "video/mp4",
			ThumbURL: "https://upload.wikimedia.org/wikipedia/commons/f/f2/Felis_silvestris_silvestris_small_gradual_decrease_of_quality_-_JPEG_compression.jpg",
			Title:    "YouTube Video",
		}
		rs := telebot.Results{&vr}
		qr := &telebot.QueryResponse{
			Results: rs,
		}
		return c.Answer(qr)
	}
	return nil
}
