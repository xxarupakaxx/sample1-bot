package main

import (
	"github.com/echo"
	"github.com/echo/middleware"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"os"
	"strings"
)

func main() {
	port:=os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	bot,err:=linebot.New(os.Getenv("CHANNELID"),os.Getenv("ACCESSTOKEN"))

	if err != nil {
		log.Fatal(err)
	}

	e:=echo.New()
	e.Use(middleware.Logger())

	e.POST("", func(c echo.Context) error {
		events,err:=bot.ParseRequest(c.Request())
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				log.Print(err)
			}
			return err
		}

		var replyText string
		replyText="cute"

		var res string
		res="ty"

		var replySticker string
		replySticker ="good moring"

		responseSticker :=linebot.NewStickerMessage("11537", "52002757")

		var replyImage string
		replyImage="cat"

		responseImage:=linebot.NewImageMessage("https://i.gyazo.com/2db8f85c496dd8f21a91eccc62ceee05.jpg", "https://i.gyazo.com/2db8f85c496dd8f21a91eccc62ceee05.jpg")

		var replyLocation string
		replyLocation ="Disney"

		responseLocation:=linebot.NewLocationMessage("TokyoDisneyLand", "千葉県浦安市舞浜", 35.632896, 139.880394)

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					replyMessage:=message.Text

					if strings.Contains(replyMessage, replyText) {
						bot.ReplyMessage(event.ReplyToken,linebot.NewTextMessage(res)).Do()
					}else if strings.Contains(replyMessage, replySticker) {
						bot.ReplyMessage(event.ReplyToken,responseSticker).Do()
					}else if strings.Contains(replyMessage, replyLocation) {
						bot.ReplyMessage(event.ReplyToken,responseLocation).Do()
					}else if strings.Contains(replyMessage, replyImage) {
						bot.ReplyMessage(event.ReplyToken, responseImage).Do()
						// 地図表示されるケース
					}

					_,err=bot.ReplyMessage(event.ReplyToken,linebot.NewTextMessage(replyMessage)).Do()
					if err != nil {
						return err
					}
				}
			}
		}
		return err
	})
}
