package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/xxarupakaxx/sample1-bot/domain"
	"github.com/xxarupakaxx/sample1-bot/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"unicode/utf8"
)

const (
	userStatusAvailable, userStatusNotAvailable string= "available", "not_available"
	HELPMESSAGE                              = `コマンド一覧
help でコマンド一覧を表示できます
channel 指定したチャンネルの最新の動画を表示します
VideoGood 指定した動画のいいね数を表示します
これから追加予定
makePlayList 指定したチャンネルのプレイリストを作るor 指定した動画たちをプレイリストにする 
upload 動画をアップロード
delete 動画を削除
mychannel 自分のアカウント情報
例
channel amazrashi Official YouTube Channel`
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("oioio")
	}

	e:=echo.New()
	e.Use(middleware.Logger())
	e.POST("/callback", lineHandler)
	e.Start(":"+port)

}

func lineHandler(c echo.Context) error {

	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.Response().WriteHeader(http.StatusBadRequest)
		}else {
			c.Response().WriteHeader(http.StatusInternalServerError)
		}
		return err
	}

	//youtebeAPI:=os.Getenv("YOUTUBE_APIKEY")
	events,err:=bot.ParseRequest(c.Request())
	if err != nil {
		log.Fatal(err)
	}


	for _,event:=range events{
		if event.Type==linebot.EventTypeFollow {
			bot.ReplyMessage(event.ReplyToken,linebot.NewTextMessage(HELPMESSAGE))
			db:=model.DBConnect()
			defer db.Close()

			userData:=domain.User{
				Id:         event.Source.UserID,
				IdType:     string(event.Source.Type),
				Timestamp:  event.Timestamp,
				ReplyToken: event.ReplyToken,
				Status:     userStatusAvailable,
			}
			_,err:=db.Exec("INSERT INTO user (id, id_type, reply_token, status) VALUES (?,?,?,?,?)",userData.Id,userData.IdType,userData.Timestamp,userData.ReplyToken,userData.Status)
			if err != nil {
				log.Fatal(err)
			}
		}
		if event.Type == linebot.EventTypeMessage {

			switch event.Type {
			case linebot.EventTypeFollow:

			case linebot.EventTypeMessage:
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					switch message.Text {
					case "text":
						resp:=linebot.NewTextMessage(message.Text)

						if _,err:=bot.ReplyMessage(event.ReplyToken,resp).Do();err!=nil{
							log.Fatal(err)
						}
					case "location":
						resp:=linebot.NewLocationMessage("現在地","宮城県多賀城市",38.297807, 141.031)
						if _,err:=bot.ReplyMessage(event.ReplyToken,resp).Do();err!=nil{
							log.Fatal(err)
						}
					case "sticker":
						resp:=linebot.NewStickerMessage("3","230")
						if _, err := bot.ReplyMessage(event.ReplyToken, resp).Do(); err != nil {
							log.Fatal(err)
						}
					case "image":
						resp:=linebot.NewImageMessage("https://farm5.staticflickr.com/4849/45718165635_328355a940_m.jpg", "https://farm5.staticflickr.com/4849/45718165635_328355a940_m.jpg")
						if _, err := bot.ReplyMessage(event.ReplyToken, resp).Do(); err != nil {
							log.Fatal(err)
						}
					case "buttontemplate":
						resp := linebot.NewTemplateMessage(
							"this is a buttons template",
							linebot.NewButtonsTemplate(
								"https://farm5.staticflickr.com/4849/45718165635_328355a940_m.jpg",
								"Menu",
								"Please select",
								linebot.NewPostbackAction("Buy", "action=buy&itemid=123", "", "displayText"),
								linebot.NewPostbackAction("Buy", "action=buy&itemid=123", "text", ""),
								linebot.NewURIAction("View detail", "http://example.com/page/123"),
							),
						)
						if _, err := bot.ReplyMessage(event.ReplyToken, resp).Do(); err != nil {
							log.Fatal(err)
						}
					case "datetimepicker":
						resp := linebot.NewTemplateMessage(
							"this is a buttons template",
							linebot.NewButtonsTemplate(
								"https://farm5.staticflickr.com/4849/45718165635_328355a940_m.jpg",
								"Menu",
								"Please select a date,  time or datetime",
								linebot.NewDatetimePickerAction("Date", "action=sel&only=date", "date", "2017-09-01", "2017-09-03", ""),
								linebot.NewDatetimePickerAction("Time", "action=sel&only=time", "time", "", "23:59", "00:00"),
								linebot.NewDatetimePickerAction("DateTime", "action=sel", "datetime", "2017-09-01T12:00", "", ""),
							),
						)
						if _, err := bot.ReplyMessage(event.ReplyToken, resp).Do(); err != nil {
							log.Fatal(err)
						}
					case"confirm":
						resp := linebot.NewTemplateMessage(
							"this is a confirm template",
							linebot.NewConfirmTemplate(
								"Are you sure?",
								linebot.NewMessageAction("Yes", "yes"),
								linebot.NewMessageAction("No", "no"),
							),
						)
						if _, err := bot.ReplyMessage(event.ReplyToken, resp).Do(); err != nil {
							log.Fatal(err)
						}
					case "carousel":
						resp := linebot.NewTemplateMessage(
							"this is a carousel template with imageAspectRatio,  imageSize and imageBackgroundColor",
							linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									"https://farm5.staticflickr.com/4849/45718165635_328355a940_m.jpg",
									"this is menu",
									"description",
									linebot.NewPostbackAction("Buy", "action=buy&itemid=111", "", ""),
									linebot.NewPostbackAction("Add to cart", "action=add&itemid=111", "", ""),
									linebot.NewURIAction("View detail", "http://example.com/page/111"),
								).WithImageOptions("#FFFFFF"),
								linebot.NewCarouselColumn(
									"https://farm5.staticflickr.com/4849/45718165635_328355a940_m.jpg",
									"this is menu",
									"description",
									linebot.NewPostbackAction("Buy", "action=buy&itemid=111", "", ""),
									linebot.NewPostbackAction("Add to cart", "action=add&itemid=111", "", ""),
									linebot.NewURIAction("View detail", "http://example.com/page/111"),
								).WithImageOptions("#FFFFFF"),
							).WithImageOptions("rectangle", "cover"),
						)
						if _, err := bot.ReplyMessage(event.ReplyToken, resp).Do(); err != nil {
							log.Fatal(err)
						}
					case "flex":
						resp:=linebot.NewFlexMessage(
							"this is a flex message",
							&linebot.BubbleContainer{
								Type:      linebot.FlexContainerTypeBubble,
								Body:      &linebot.BoxComponent{
									Type:            linebot.FlexComponentTypeBox,
									Layout:          linebot.FlexBoxLayoutTypeVertical,
									Contents:        []linebot.FlexComponent{
										&linebot.TextComponent{
											Type:       linebot.FlexComponentTypeText,
											Text:       "Hello",
										},
										&linebot.TextComponent{Type: linebot.FlexComponentTypeText,Text: "World"},
									},
								},
							})
						if _, err := bot.ReplyMessage(event.ReplyToken, resp).Do(); err != nil {
							log.Fatal(err)
						}
					case "quick":
						resp:=linebot.NewTextMessage(
							"select your favorite food category or send me your location",
							).WithQuickReplies(
							linebot.NewQuickReplyItems(
								linebot.NewQuickReplyButton("https://trap.jp/content/images/2020/01/traP_logo_icon-1.png", linebot.NewMessageAction("traP", "traP")),
								linebot.NewQuickReplyButton("https://trap.jp/content/images/2021/04/trap_logo_full.jpg", linebot.NewMessageAction("Tempura", "Tempura")),
								linebot.NewQuickReplyButton("", linebot.NewLocationAction("Send location")),
							))
						if _, err := bot.ReplyMessage(event.ReplyToken, resp).Do(); err != nil {
							log.Fatal(err)
						}
					}
				}
			}
			/*switch message:=event.Message.(type){
			case *linebot.LocationMessage:
				sendRestoInfo(bot,event)
			case *linebot.TextMessage:
				replymessage:=message.Text
				
				user:=bot.GetProfile(event.Source.UserID)
				replymessage:=message.Text
				if replymessage=="help"{
					userProfile,_:=user.Do()
					bot.ReplyMessage(event.ReplyToken,linebot.NewTextMessage(userProfile.DisplayName+"さん\n"+HELPMESSAGE)).Do()
				}else if strings.Contains(replymessage, "channel") {
					channelName:=replymessage[8:len(replymessage)]
					videoes,err:=model.SerarchYoutubeChannel(channelName)
					if err != nil {
						log.Fatal(err)
					}
					for _,video:=range videoes{
						baseurl:="https://www.youtube.com/watch?v="
						if len(video.VideoID)!=0 {
							url:=baseurl+video.VideoID
							replymessage+="\n"
							replymessage+=url
						}
					}
					bot.ReplyMessage(event.ReplyToken,linebot.NewTextMessage(replymessage))
				}
				if _,err :=bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replymessage)).Do();err!=nil {
					log.Print(err)
				}
			}*/

		}
	}
	return c.String(http.StatusOK,"OK")

}

func sendRestoInfo(bot *linebot.Client, e *linebot.Event) {
	msg := e.Message.(*linebot.LocationMessage)

	lat := strconv.FormatFloat(msg.Latitude, 'f', 2, 64)
	lng := strconv.FormatFloat(msg.Longitude, 'f', 2, 64)

	replyMsg := getRestoInfo(lat, lng)

	res := linebot.NewTemplateMessage(
		"レストラン一覧",
		linebot.NewCarouselTemplate(replyMsg...).WithImageOptions("rectangle", "cover"),
	)

	if _, err := bot.ReplyMessage(e.ReplyToken, res).Do(); err != nil {
		log.Print(err)
	}
}


func getRestoInfo(lat string, lng string) []*linebot.CarouselColumn {
	apikey := os.Getenv("HOTPEPPERKEY")
	url := fmt.Sprintf(
		"https://webservice.recruit.co.jp/hotpepper/gourmet/v1/?format=json&key=%s&lat=%s&lng=%s",
		apikey, lat, lng)

	// リクエストしてボディを取得
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data domain.Response
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	var ccs []*linebot.CarouselColumn
	for _, shop := range data.Results.Shop {
		addr := shop.Address
		if 60 < utf8.RuneCountInString(addr) {
			addr = string([]rune(addr)[:60])
		}

		cc := linebot.NewCarouselColumn(
			shop.Photo.Mobile.L,
			shop.Name,
			addr,
			linebot.NewURIAction("ホットペッパーで開く", shop.URLS.PC),
		).WithImageOptions("#FFFFFF")
		ccs = append(ccs, cc)
	}
	return ccs
}