package domain

import (
	"time"
)

type Video struct {
	VideoID      string    `json:"videoId"`
	PublishedAt  time.Time `json:"publishedAt"`
	ChannelID    string    `json:"channelId"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	ChannelTitle string    `json:"channelTitle"`
}




