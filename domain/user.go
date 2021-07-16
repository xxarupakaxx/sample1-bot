package domain

import "time"

type User struct {
	Id         string
	DisplayName string
	IdType     string
	Timestamp  time.Time
	ReplyToken string
}
