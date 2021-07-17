package model

import (
	"encoding/json"
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/xxarupakaxx/sample1-bot/domain"
	"io/ioutil"
	"log"
	"net/http"
)

func getWeather(code string) *domain.Weather{
	url := fmt.Sprintf("https://weather.tsukumijima.net/api/forecast/city/%s", code)
	res,err:=http.Get(url)
	if err != nil {
		log.Fatalf("Coundnot get https://weather.tsukumijima.net/api/forecast/:%v",err)
	}
	defer res.Body.Close()
	body,err:=ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Failed in Reading https://weather.tsukumijima.net/api/forecast/city/ response :%v",err)
	}
	var data *domain.Weather
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalf("Failed in Changing Json: %v",err)
	}
	return data
}
func SendWeather(bot *linebot.Client, event *linebot.Event,code string) {
	resp:=linebot.NewTemplateMessage(
		"weather",
		linebot.NewButtonsTemplate(
				"https://weather.tsukumijima.net/logo.png",
				"Your Region Code",
				"Please Select your code",
				linebot.NewDatetimePickerAction(
					"code",
					"https://weather.tsukumijima.net/api/forecast/city/:code",
					"datetime",
					"2017-09-01T12:00",
					"",
					"",
					),
			),
		)
	if _,err:=bot.ReplyMessage(event.ReplyToken,resp).Do();err!=nil{
		log.Fatalf("coundnot sending resp:%v",err)
	}
	data:=getWeather(code)
	message:=data.Title+"\n"+data.PublicTimeFormatted + data.Description.Text +"\n今日は"+data.Forecasts[0].Telop+ "で風邪向きが"+data.Forecasts[0].Detail.Wind+"波が"+data.Forecasts[0].Detail.Wave+"です\nまた、最高気温が"+data.Forecasts[0].Temperature.Max.Celsius+"\n最低気温が"+data.Forecasts[0].Temperature.Min.Celsius+"です\n0 時から 6 時までの降水確率は"+data.Forecasts[0].ChanceOfRain.T0006+"\n"+"6 時から 12 時までの降水確率"+data.Forecasts[0].ChanceOfRain.T0612+"\n"+"12 時から 18 時までの降水確率"+data.Forecasts[0].ChanceOfRain.T1218+"\n"+"18 時から 24 時までの降水確率は"+data.Forecasts[0].ChanceOfRain.T1824+"となるでしょう"
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do(); err != nil {
		log.Fatalf("Coundnot posting weather:%v",err)
	}
}