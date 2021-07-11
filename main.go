package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/xxarupakaxx/sample1-bot/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

const HELPMESSAGE = `コマンド一覧
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

		if event.Type == linebot.EventTypeMessage {
			switch message:=event.Message.(type){
			case *linebot.LocationMessage:
				sendRestoInfo(bot,event)
			case *linebot.TextMessage:
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
			}

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

// response APIレスポンス
type response struct {
	Results results `json:"results"`
}

// results APIレスポンスの内容
type results struct {
	Shop []shop `json:"shop"`
}

// shop レストラン一覧
type shop struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Photo   photo  `json:"photo"`
	URLS    urls   `json:"urls"`
}

// photo 写真URL一覧
type photo struct {
	Mobile mobile `json:"mobile"`
}

// mobile モバイル用の写真URL
type mobile struct {
	L string `json:"l"`
}

// urls URL一覧
type urls struct {
	PC string `json:"pc"`
}

func getRestoInfo(lat string, lng string) []*linebot.CarouselColumn {
	apikey := "(自分のAPIKEYを入力)"
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

	var data response
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