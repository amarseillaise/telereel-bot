package bot

import (
	"gopkg.in/telebot.v4"

	"github.com/amarseillaise/instareels_to_telegram/services"
)

func OnTextHandler(c telebot.Context) error {
	_url := c.Text()
	if IsValidInstagramReelURL(_url) {
		c.Notify(telebot.UploadingVideo)
		shortcode := services.ParseShortcode(_url)
		reelData, err := services.GetReelData(shortcode)
		if err == nil {
			teleVideo := MakeVideo(reelData.VideoUrl)
			teleVideo.Caption = reelData.Caption
			return c.Reply(teleVideo)
		} else {
			return c.Reply(err.Error())
		}
	}
	return nil
}

func OnQueryHandler(c telebot.Context) error {
	_url := c.Query().Text
	if IsValidInstagramReelURL(_url) {
		shortcode := services.ParseShortcode(_url)
		reelData, err := services.GetReelData(shortcode)
		if err != nil {
			return nil
		}
		vr := telebot.VideoResult{
			URL:      reelData.VideoUrl,
			Caption:  reelData.Caption,
			MIME:     "video/mp4",
			ThumbURL: "https://upload.wikimedia.org/wikipedia/commons/f/f2/Felis_silvestris_silvestris_small_gradual_decrease_of_quality_-_JPEG_compression.jpg",
			Title:    reelData.Title,
		}
		rs := telebot.Results{&vr}
		qr := &telebot.QueryResponse{
			Results: rs,
		}
		return c.Answer(qr)
	}
	return nil
}
