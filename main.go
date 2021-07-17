package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/xxarupakaxx/sample1-bot/domain"
	"github.com/xxarupakaxx/sample1-bot/model"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	HELPMESSAGE                              = `コマンド一覧
help でコマンド一覧を表示できます
これから追加予定
makePlayList 指定したチャンネルのプレイリストを作るor 指定した動画たちをプレイリストにする 
upload 動画をアップロード
delete 動画を削除
mychannel 自分のアカウント情報
googleDriveに動画を保存する
二つの動画をつなげて返す
できたら画像加工もできたらいいね
weather code で天気予報取得
位置情報を取得して近くのお店を表示`
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

	events,err:=bot.ParseRequest(c.Request())
	if err != nil {
		log.Fatal(err)
	}


	for _,event:=range events{
		if event.Type==linebot.EventTypeFollow {
			userId:=event.Source.UserID
			user,err:=bot.GetProfile(userId).Do()
			if err!=nil{
				log.Fatalf("Failed in getting user :%v",err)
			}
			db:=model.DBConnect()
			defer db.Close()

			userData:=domain.User{
				Id:         event.Source.UserID,
				DisplayName: user.DisplayName,
				IdType:     string(event.Source.Type),
				Timestamp:  event.Timestamp,
				ReplyToken: event.ReplyToken,
			}
			err=db.QueryRow("SELECT * from user where user.id=$1",userData.Id).Scan(&userData.Id, &userData.DisplayName,&userData.IdType,&userData.Timestamp,&userData.ReplyToken)
			if err != nil {
				_,err=db.Exec("INSERT INTO user VALUES (?,?,?,?,?)",userData.Id,userData.DisplayName,userData.IdType,userData.Timestamp,userData.ReplyToken)
				if err != nil {
					log.Fatalf("Couldnot add user:%v",err)
				}
			}

			text:=user.DisplayName+"さん\n"+HELPMESSAGE
			if _,err:=bot.PushMessage(userId,linebot.NewTextMessage(text),linebot.NewStickerMessage("8522","16581267")).Do();err!=nil{
				log.Fatalf("Failed in Pushing message:%v",err)
			}

		}
		switch message := event.Message.(type) {
		case *linebot.LocationMessage:
			model.SendRestoInfo(bot,event)
		case *linebot.TextMessage:
			user,_:=bot.GetProfile(event.Source.UserID).Do()
			if message.Text == "help" {
				text:=user.DisplayName+"さん\n"+HELPMESSAGE
				if _,err:=bot.ReplyMessage(event.ReplyToken,linebot.NewTextMessage(text)).Do();err!=nil{
					log.Fatalf("Failed in Replying message:%v",err)
				}

			}
			if message.Text=="tameda" {
				//text:=user.DisplayName+"がtamedaと送信しました"
				bot.ReplyMessage(event.ReplyToken,linebot.NewFlexMessage("メンテナンス終了",
					&linebot.BubbleContainer{
						Type: linebot.FlexContainerTypeBubble,
						Body: &linebot.BoxComponent{
							Type:            linebot.FlexComponentTypeBox,
							Layout:          linebot.FlexBoxLayoutTypeBaseline,
							Contents:        []linebot.FlexComponent{
								&linebot.ImageComponent{
									Type:            linebot.FlexComponentTypeImage,
									URL:             "https://pbs.twimg.com/profile_images/1364204318790836232/dgyZpvy6_400x400.jpg",
									Gravity:         linebot.FlexComponentGravityTypeBottom,
									Size:            linebot.FlexImageSizeTypeFull,
									AspectRatio:     linebot.FlexImageAspectRatioType1to3,
									AspectMode:      linebot.FlexImageAspectModeTypeFit,
									BackgroundColor: "#f891df",
								},
							},
							CornerRadius:  linebot.FlexComponentCornerRadiusTypeMd  ,
							BackgroundColor: "#6de765",
							BorderColor:     "#3BAF75",
							Action:         linebot.NewURIAction("tameda","https://twitter.com/TamerNazeda") ,
						},
					},
				)).Do()
			}
			if strings.Contains(message.Text, "weather") {
				msg:=message.Text
				code:=msg[len("weather "):]
				model.SendWeather(bot,event,code)
			}

		case *linebot.VideoMessage:
			if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("未実装")).Do(); err != nil {
				log.Fatalf("Failed in getting url:%v",err)
			}

			/*result,err:=http.Get(url)
			if err != nil {
				log.Fatalf("Failed in Getting url:%v",err)
			}
			defer result.Body.Close()
			body,err:=ioutil.ReadAll(result.Body)
			if err != nil {
				log.Fatalf("Failed in Reading url:%v",err)
			}
			if _,err:=bot.ReplyMessage(event.ReplyToken,linebot.NewTextMessage(string(body))).Do();err!=nil{
				log.Fatalf("Failed in Replying message:%v",err)
			}*/
		}
	}
	return c.String(http.StatusOK,"OK")

}

