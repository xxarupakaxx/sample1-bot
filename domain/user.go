package domain

import "time"

type User struct {
	Id         string
	IdType     string
	Timestamp  time.Time
	ReplyToken string
	Status     string
}
