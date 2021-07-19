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

func GetWeather(code string) *domain.Weather{
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
	data:=GetWeather(code)
	log.Println("\n",data,"\n")
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
			URL:             data.Forecasts[0].Image.URL,
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
					Text: "最高気温 : " + data.Forecasts[0].Temperature.Max.Celsius + "℃\n",
					Flex:       linebot.IntPtr(1),
					Size:       linebot.FlexTextSizeTypeSm,
					Wrap:       true,
					//Action:     nil,
					MaxLines:   linebot.IntPtr(2),
				},
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: "最低気温 : " + data.Forecasts[0].Temperature.Min.Celsius + "℃\n",
					Flex:       linebot.IntPtr(1),
					Size:       linebot.FlexTextSizeTypeSm,
					Wrap:       true,
					//Action:     nil,
					MaxLines:   linebot.IntPtr(2),
				},
				&linebot.TextComponent{
					Type:       linebot.FlexComponentTypeText,
					Text:       data.Description.Text,
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
					Action:  linebot.NewURIAction("天気予報", data.Link),
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
	if _,err:=bot.ReplyMessage(event.ReplyToken,resp).Do();err != nil {
		log.Fatalf("weather response error :%v",err)
	}

}