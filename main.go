package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/line/line-bot-sdk-go/linebot"
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
	/*db:=model.DBConnect()
	defer db.Close()*/
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
		/*if event.Type==linebot.EventTypeFollow {
			userId:=event.Source.UserID
			user,err:=bot.GetProfile(userId).Do()
			if err!=nil{
				log.Fatalf("Failed in getting user :%v",err)
			}



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

		}*/
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
				debug(bot,event)
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
		}
	}
	return c.String(http.StatusOK,"OK")

}

func debug(bot *linebot.Client, event *linebot.Event) {
	resp:=linebot.NewFlexMessage("Weather Information",&linebot.BubbleContainer{
		Type:      linebot.FlexContainerTypeBubble,
		Direction: linebot.FlexBubbleDirectionTypeLTR,
		Header:    &linebot.BoxComponent{
			Type:            linebot.FlexComponentTypeBox,
			Layout:          linebot.FlexBoxLayoutTypeBaseline,
			Contents:        []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:       linebot.FlexComponentTypeText,
					Text:       "今日の天気",
					Size:       linebot.FlexTextSizeTypeLg,
					Align:      linebot.FlexComponentAlignTypeCenter,
					Weight:     linebot.FlexTextWeightTypeBold,
					//Color:      "",
					//Action:     nil,
				},
			},
			CornerRadius:    linebot.FlexComponentCornerRadiusTypeXxl,
			BorderColor:     "#00bfff",
			//Action: nil,
		},
		Hero:      &linebot.ImageComponent{
			Type:            linebot.FlexComponentTypeImage,
			URL:             "https://twitter.com/home",
			Size:            linebot.FlexImageSizeTypeXxl,
			AspectRatio:     linebot.FlexImageAspectRatioType1to1,
			AspectMode:      linebot.FlexImageAspectModeTypeFit,
			//Action:          nil,
		},
		Body:      &linebot.BoxComponent{
			Type:            linebot.FlexComponentTypeBox,
			Layout:          linebot.FlexBoxLayoutTypeVertical,
			Contents:        []linebot.FlexComponent{
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: "最高気温 : " +"33" + "℃\n",
					Flex:       linebot.IntPtr(1),
					Size:       linebot.FlexTextSizeTypeXl,
					Wrap:       true,
					//Action:     nil,
					MaxLines:   linebot.IntPtr(2),
				},
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: "最低気温 : " + "10" + "℃\n",
					Flex:       linebot.IntPtr(1),
					Size:       linebot.FlexTextSizeTypeXl,
					Wrap:       true,
					//Action:     nil,
					MaxLines:   linebot.IntPtr(2),
				},
				&linebot.TextComponent{
					Type:       linebot.FlexComponentTypeText,
					Text:       "ワクチンの情報を詳しく知るのは重要なことだけれどもワクチンを怖がりすぎたら注射のストレスで病気になるそれが副反応とか言われちゃあおしまいだよね打ちたい打ちたい、わくわく！\nで打ちに行かなきゃ。ワクワク！、ね。ワクチンワクワク！、ね。ヨコハマヨーヨーね。うん。",
					//Contents:   nil,
					Flex:       linebot.IntPtr(3),
					Size:       linebot.FlexTextSizeTypeSm,
					Wrap:       true,
					//Color:      "",
					//Action:     nil,
					MaxLines:   linebot.IntPtr(5),
				},
			},
			BorderColor:     "#5cd8f7",
			//Action:          nil,
		},
		Footer: &linebot.BoxComponent{
			Type:            linebot.FlexComponentTypeBox,
			Layout:          linebot.FlexBoxLayoutTypeBaseline,
			Contents:        []linebot.FlexComponent{
				&linebot.ButtonComponent{
					Type:    linebot.FlexComponentTypeButton,
					Action:  linebot.NewURIAction("天気予報", "https://twitter.com/home"),
					Style:   linebot.FlexButtonStyleTypeLink,
					Color:   "#80DEEA",
				},
			},
			BorderColor:     "#90CAF9",
		},

		Styles:    &linebot.BubbleStyle{
			Header: &linebot.BlockStyle{
				Separator:       true,
				SeparatorColor:  "#2196F3",
			},
			Hero:   &linebot.BlockStyle{
				Separator:      true,
				SeparatorColor: "#2196F3" ,

			},
			Body:   &linebot.BlockStyle{
				Separator:      true,
				SeparatorColor: "#37474F",
			},
			Footer: &linebot.BlockStyle{
				Separator:      true,
				SeparatorColor: "#2196F3",
			},
		},
	})
	if _, err := bot.ReplyMessage(event.ReplyToken, resp).Do(); err != nil {
		log.Printf("debug error:%v",err)
	}
}

