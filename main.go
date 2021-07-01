package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"os"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("oioio")
	}
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	//youtebeAPI:=os.Getenv("YOUTUBE_APIKEY")

	e:=echo.New()
	e.Use(middleware.Logger())
	e.POST("/callback", func(c echo.Context) error {
		t:=time.Now().Format("2006-1-02-3-4")
		message:=linebot.NewTextMessage(t)
		if _, err := bot.BroadcastMessage(message).Do(); err != nil {
			log.Fatal(err)
			return err
		}
		return err
	})
	e.Start(":"+port)
	/*router := gin.New()
	router.Use(gin.Logger())

	// LINE Messaging API ルーティング
	router.POST("/callback", func(c *gin.Context) {
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				log.Print(err)
			}
			return
		}

		// "可愛い" 単語を含む場合、返信される
		var replyText string
		replyText = "可愛い"

		// チャットの回答
		var response string
		response = "ありがとう！！"

		// "おはよう" 単語を含む場合、返信される
		var replySticker string
		replySticker = "おはよう"

		// スタンプで回答が来る
		responseSticker := linebot.NewStickerMessage("11537", "52002757")

		// "猫" 単語を含む場合、返信される
		var replyImage string
		replyImage = "猫"

		// 猫の画像が表示される
		responseImage := linebot.NewImageMessage("https://i.gyazo.com/2db8f85c496dd8f21a91eccc62ceee05.jpg", "https://i.gyazo.com/2db8f85c496dd8f21a91eccc62ceee05.jpg")

		// "ディズニー" 単語を含む場合、返信される
		var replyLocation string
		replyLocation = "ディズニー"

		// ディズニーが地図表示される
		responseLocation := linebot.NewLocationMessage("東京ディズニーランド", "千葉県浦安市舞浜", 35.632896, 139.880394)

		for _, event := range events {
			// イベントがメッセージの受信だった場合
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				// メッセージがテキスト形式の場合
				case *linebot.TextMessage:
					replyMessage := message.Text
					// テキストで返信されるケース
					if strings.Contains(replyMessage, replyText) {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(response)).Do()
						// スタンプで返信されるケース
					} else if strings.Contains(replyMessage, replySticker) {
						bot.ReplyMessage(event.ReplyToken, responseSticker).Do()
						// 画像で返信されるケース
					} else if strings.Contains(replyMessage, replyImage) {
						bot.ReplyMessage(event.ReplyToken, responseImage).Do()
						// 地図表示されるケース
					} else if strings.Contains(replyMessage, replyLocation) {
						bot.ReplyMessage(event.ReplyToken, responseLocation).Do()
					}
					// 上記意外は、おうむ返しで返信
					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
					if err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	router.Run(":" + port)*/
}

