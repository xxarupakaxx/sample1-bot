package model

import (
	"encoding/json"
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/xxarupakaxx/sample1-bot/domain"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
func SendWeather(bot *linebot.Client, event *linebot.Event,cityName string) {
	data:=GetWeather(PrefCode(cityName))
	/*result := 0
	if  r:=db.QueryRow("SELECT exists(SELECT code from City where code=$1)",code).Scan(&result);r!=nil{
		log.Fatalf("Couldnot queryRow : %v",r)
	}
	if result==0 {
		if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("存在しない都市名です")).Do(); err != nil {
			log.Fatalf("Could Not sending :%v",err)
		}
		return
	}*/
	description:=strings.NewReplacer("\n","").Replace(data.Description.Text)
	var text string
	for s, _ := range domain.CodeCity {
		text+=s+"、"
	}
	resp:=linebot.NewFlexMessage(
		"Weather Information",
		&linebot.CarouselContainer{
			Type: linebot.FlexContainerTypeCarousel,
			Contents: []*linebot.BubbleContainer{
				{
					Type:      linebot.FlexContainerTypeBubble,
					Direction: linebot.FlexBubbleDirectionTypeLTR,
					Header: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeBaseline,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type:   linebot.FlexComponentTypeText,
								Text:   "今日の天気",
								Size:   linebot.FlexTextSizeTypeLg,
								Align:  linebot.FlexComponentAlignTypeCenter,
								Weight: linebot.FlexTextWeightTypeBold,
								//Color:      "",
								//Action:     nil,
							},
						},
						CornerRadius: linebot.FlexComponentCornerRadiusTypeXxl,
						BorderColor:  "#00bfff",
						//Action: nil,
					},
					Hero: &linebot.ImageComponent{
						Type:        linebot.FlexComponentTypeImage,
						URL:         ConvertTelop(data.Forecasts[0].Telop),
						Size:        linebot.FlexImageSizeTypeXxl,
						AspectRatio: linebot.FlexImageAspectRatioType1to1,
						AspectMode:  linebot.FlexImageAspectModeTypeFit,
						//Action:          nil,
					},
					Body: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeVertical,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type: linebot.FlexComponentTypeText,
								Text: "最高気温 : " + data.Forecasts[0].Temperature.Max.Celsius + "℃\n",
								Flex: linebot.IntPtr(1),
								Size: linebot.FlexTextSizeTypeXl,
								Wrap: true,
								//Action:     nil,
								MaxLines: linebot.IntPtr(2),
							},
							&linebot.TextComponent{
								Type: linebot.FlexComponentTypeText,
								Text: "最低気温 : " + data.Forecasts[0].Temperature.Min.Celsius + "℃\n",
								Flex: linebot.IntPtr(1),
								Size: linebot.FlexTextSizeTypeXl,
								Wrap: true,
								//Action:     nil,
								MaxLines: linebot.IntPtr(2),
							},
							&linebot.TextComponent{
								Type: linebot.FlexComponentTypeText,
								Text: description,
								//Contents:   nil,
								Flex: linebot.IntPtr(6),
								Size: linebot.FlexTextSizeTypeSm,
								Wrap: true,
								//Color:      "",
								//Action:     nil,
								MaxLines: linebot.IntPtr(10),
							},
						},
						BorderColor: "#5cd8f7",
						//Action:          nil,
					},
					Footer: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeBaseline,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type:   linebot.FlexComponentTypeText,
								Text:   "天気予報",
								Align:  linebot.FlexComponentAlignTypeCenter,
								Wrap:   true,
								Color:  "#2196F3",
								Action: linebot.NewURIAction("天気予報", data.Link),
								Style:  linebot.FlexTextStyleTypeItalic,
								Weight: linebot.FlexTextWeightTypeBold,
							},
						},
						Action:      linebot.NewURIAction("天気予報", data.Link),
						BorderColor: "#90CAF9",
					},

					Styles: &linebot.BubbleStyle{
						Header: &linebot.BlockStyle{
							Separator:      true,
							SeparatorColor: "#2196F3",
						},
						Hero: &linebot.BlockStyle{
							Separator:      true,
							SeparatorColor: "#2196F3",
						},
						Body: &linebot.BlockStyle{
							Separator:      true,
							SeparatorColor: "#37474F",
						},
						Footer: &linebot.BlockStyle{
							Separator:      true,
							SeparatorColor: "#2196F3",
						},
					},
				},
				{
					Type:      linebot.FlexContainerTypeBubble,
					Direction: linebot.FlexBubbleDirectionTypeLTR,
					Header: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeBaseline,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type:   linebot.FlexComponentTypeText,
								Text:   "明日の天気",
								Size:   linebot.FlexTextSizeTypeLg,
								Align:  linebot.FlexComponentAlignTypeCenter,
								Weight: linebot.FlexTextWeightTypeBold,
								//Color:      "",
								//Action:     nil,
							},
						},
						CornerRadius: linebot.FlexComponentCornerRadiusTypeXxl,
						BorderColor:  "#00bfff",
						//Action: nil,
					},
					Hero: &linebot.ImageComponent{
						Type:        linebot.FlexComponentTypeImage,
						URL:         ConvertTelop(data.Forecasts[1].Telop),
						Size:        linebot.FlexImageSizeTypeXxl,
						AspectRatio: linebot.FlexImageAspectRatioType1to1,
						AspectMode:  linebot.FlexImageAspectModeTypeFit,
						//Action:          nil,
					},
					Body: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeVertical,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type: linebot.FlexComponentTypeText,
								Text: "最高気温 : " + data.Forecasts[1].Temperature.Max.Celsius + "℃\n",
								Flex: linebot.IntPtr(1),
								Size: linebot.FlexTextSizeTypeXl,
								Wrap: true,
								//Action:     nil,
								MaxLines: linebot.IntPtr(2),
							},
							&linebot.TextComponent{
								Type: linebot.FlexComponentTypeText,
								Text: "最低気温 : " + data.Forecasts[1].Temperature.Min.Celsius + "℃\n",
								Flex: linebot.IntPtr(1),
								Size: linebot.FlexTextSizeTypeXl,
								Wrap: true,
								//Action:     nil,
								MaxLines: linebot.IntPtr(2),
							},
							&linebot.TextComponent{
								Type: linebot.FlexComponentTypeText,
								Text: description,
								//Contents:   nil,
								Flex: linebot.IntPtr(6),
								Size: linebot.FlexTextSizeTypeSm,
								Wrap: true,
								//Color:      "",
								//Action:     nil,
								MaxLines: linebot.IntPtr(10),
							},
						},
						BorderColor: "#5cd8f7",
						//Action:          nil,
					},
					Footer: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeBaseline,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type:   linebot.FlexComponentTypeText,
								Text:   "天気予報",
								Align:  linebot.FlexComponentAlignTypeCenter,
								Wrap:   true,
								Color:  "#2196F3",
								Action: linebot.NewURIAction("天気予報", data.Link),
								Style:  linebot.FlexTextStyleTypeItalic,
								Weight: linebot.FlexTextWeightTypeBold,
							},
						},
						Action:      linebot.NewURIAction("天気予報", data.Link),
						BorderColor: "#90CAF9",
					},

					Styles: &linebot.BubbleStyle{
						Header: &linebot.BlockStyle{
							Separator:      true,
							SeparatorColor: "#2196F3",
						},
						Hero: &linebot.BlockStyle{
							Separator:      true,
							SeparatorColor: "#2196F3",
						},
						Body: &linebot.BlockStyle{
							Separator:      true,
							SeparatorColor: "#37474F",
						},
						Footer: &linebot.BlockStyle{
							Separator:      true,
							SeparatorColor: "#2196F3",
						},
					},
				},
				{
					Type:      linebot.FlexContainerTypeBubble,
					Direction: linebot.FlexBubbleDirectionTypeLTR,
					Header: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeBaseline,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type:   linebot.FlexComponentTypeText,
								Text:   "明後日の天気",
								Size:   linebot.FlexTextSizeTypeLg,
								Align:  linebot.FlexComponentAlignTypeCenter,
								Weight: linebot.FlexTextWeightTypeBold,
								//Color:      "",
								//Action:     nil,
							},
						},
						CornerRadius: linebot.FlexComponentCornerRadiusTypeXxl,
						BorderColor:  "#00bfff",
						//Action: nil,
					},
					Hero: &linebot.ImageComponent{
						Type:        linebot.FlexComponentTypeImage,
						URL:         ConvertTelop(data.Forecasts[2].Telop),
						Size:        linebot.FlexImageSizeTypeXxl,
						AspectRatio: linebot.FlexImageAspectRatioType1to1,
						AspectMode:  linebot.FlexImageAspectModeTypeFit,
						//Action:          nil,
					},
					Body: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeVertical,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type: linebot.FlexComponentTypeText,
								Text: "最高気温 : " + data.Forecasts[2].Temperature.Max.Celsius + "℃\n",
								Flex: linebot.IntPtr(1),
								Size: linebot.FlexTextSizeTypeXl,
								Wrap: true,
								//Action:     nil,
								MaxLines: linebot.IntPtr(2),
							},
							&linebot.TextComponent{
								Type: linebot.FlexComponentTypeText,
								Text: "最低気温 : " + data.Forecasts[2].Temperature.Min.Celsius + "℃\n",
								Flex: linebot.IntPtr(1),
								Size: linebot.FlexTextSizeTypeXl,
								Wrap: true,
								//Action:     nil,
								MaxLines: linebot.IntPtr(2),
							},
							&linebot.TextComponent{
								Type: linebot.FlexComponentTypeText,
								Text: description,
								//Contents:   nil,
								Flex: linebot.IntPtr(6),
								Size: linebot.FlexTextSizeTypeSm,
								Wrap: true,
								//Color:      "",
								//Action:     nil,
								MaxLines: linebot.IntPtr(10),
							},
						},
						BorderColor: "#5cd8f7",
						//Action:          nil,
					},
					Footer: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeBaseline,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type:   linebot.FlexComponentTypeText,
								Text:   "天気予報",
								Align:  linebot.FlexComponentAlignTypeCenter,
								Wrap:   true,
								Color:  "#2196F3",
								Action: linebot.NewURIAction("天気予報", data.Link),
								Style:  linebot.FlexTextStyleTypeItalic,
								Weight: linebot.FlexTextWeightTypeBold,
							},
						},
						Action:      linebot.NewURIAction("天気予報", data.Link),
						BorderColor: "#90CAF9",
					},

					Styles: &linebot.BubbleStyle{
						Header: &linebot.BlockStyle{
							Separator:      true,
							SeparatorColor: "#2196F3",
						},
						Hero: &linebot.BlockStyle{
							Separator:      true,
							SeparatorColor: "#2196F3",
						},
						Body: &linebot.BlockStyle{
							Separator:      true,
							SeparatorColor: "#37474F",
						},
						Footer: &linebot.BlockStyle{
							Separator:      true,
							SeparatorColor: "#2196F3",
						},
					},
				},
				{
					Type:      linebot.FlexContainerTypeBubble,
					Direction: linebot.FlexBubbleDirectionTypeLTR,
					Header: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeBaseline,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type:   linebot.FlexComponentTypeText,
								Text:   "取得できる地域一覧",
								Size:   linebot.FlexTextSizeTypeLg,
								Align:  linebot.FlexComponentAlignTypeCenter,
								Weight: linebot.FlexTextWeightTypeBold,
								//Color:      "",
								//Action:     nil,
							},
						},
						CornerRadius: linebot.FlexComponentCornerRadiusTypeXxl,
						BorderColor:  "#00bfff",
						//Action: nil,
					},
					/*Hero:      &linebot.ImageComponent{
						Type:            linebot.FlexComponentTypeImage,
						URL:             ConvertTelop(data.Forecasts[2].Telop),
						Size:            linebot.FlexImageSizeTypeXxl,
						AspectRatio:     linebot.FlexImageAspectRatioType1to1,
						AspectMode:      linebot.FlexImageAspectModeTypeFit,
						//Action:          nil,
					},*/
					Body: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeVertical,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type: linebot.FlexComponentTypeText,
								Text: text,
								Size: linebot.FlexTextSizeTypeSm,
								Wrap: true,
								//Action:     nil,
								MaxLines: linebot.IntPtr(30),
							},
						},
						BorderColor: "#5cd8f7",
						//Action:          nil,
					},
					Footer: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeBaseline,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type:   linebot.FlexComponentTypeText,
								Text:   "天気一覧表",
								Align:  linebot.FlexComponentAlignTypeCenter,
								Wrap:   true,
								Color:  "#2196F3",
								Action: linebot.NewURIAction("天気一覧表", "https://www.jma.go.jp/bosai/map.html#5/33.164/137.413/&contents=forecast"),
								Style:  linebot.FlexTextStyleTypeItalic,
								Weight: linebot.FlexTextWeightTypeBold,
							},
						},
						Action:      linebot.NewURIAction("天気一覧表", "https://www.jma.go.jp/bosai/map.html#5/33.164/137.413/&contents=forecast"),
						BorderColor: "#90CAF9",
					},

					Styles: &linebot.BubbleStyle{
						Header: &linebot.BlockStyle{
							Separator:      true,
							SeparatorColor: "#2196F3",
						},
						Hero: &linebot.BlockStyle{
							Separator:      true,
							SeparatorColor: "#2196F3",
						},
						Body: &linebot.BlockStyle{
							Separator:      true,
							SeparatorColor: "#37474F",
						},
						Footer: &linebot.BlockStyle{
							Separator:      true,
							SeparatorColor: "#2196F3",
						},
					},
				},
				/*{
					Type:   linebot.FlexContainerTypeBubble,
					Direction: linebot.FlexBubbleDirectionTypeLTR,
					Header: &linebot.BoxComponent{
						Type:        linebot.FlexComponentTypeBox,
						Layout:          linebot.FlexBoxLayoutTypeBaseline,
						Contents:    []linebot.FlexComponent{
							&linebot.TextComponent{
								Type:  linebot.FlexComponentTypeText,
								Text:  "天気の取得できる地域一覧",
								Size:  linebot.FlexTextSizeTypeLg,
								Align: linebot.FlexComponentAlignTypeCenter,
								Wrap:  true,
								Color: "#6AB7FF",
							},
						},
						BorderColor: "#6AB7FF",
					},
					Body:   &linebot.BoxComponent{
						Type:     linebot.FlexComponentTypeBox,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type:     linebot.FlexComponentTypeText,
								Text:     text,
								Size:     linebot.FlexTextSizeTypeSm,
								Wrap:     true,
								Weight: linebot.FlexTextWeightTypeBold,
								MaxLines: linebot.IntPtr(8),
							},
						},
					},
				},*/
			},
		},
	)
	if _,err:=bot.ReplyMessage(event.ReplyToken,resp).Do();err != nil {
		log.Fatalf("weather response error :%v",err)
	}

}
func ConvertTelop(telop string) string {
	if strings.Contains(telop,"晴") && strings.Contains(telop,"曇") {
		return "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQo2GqJ-kmQi-OOw2O5YgHIT9ATmffsvMA0Rpjh7TbYy-7nqB7NZHEGgH2zisO3l6IQC7A&usqp=CAU"
	}
	if strings.Contains(telop, "晴") {
		return "https://illust8.com/wp-content/uploads/2018/08/weather_sun_solar_illust_1080.png"
	}
	if strings.Contains(telop, "曇") {
		return "https://www.jalan.net/jalan/images/pictLL/Y5/L336655/L3366550005036729.jpg"
	}
	if strings.Contains(telop, "雨") {
		return "https://frame-illust.com/fi/wp-content/uploads/2016/05/7749.png"
	}
	if strings.Contains(telop, "雨") && strings.Contains(telop, "曇") {
		return "https://marinchu.com/wp/wp-content/uploads/kumori-300x300.png"
	}
	if strings.Contains(telop, "雨") || strings.Contains(telop, "晴") {
		return "https://lh3.googleusercontent.com/proxy/TKLgewsO3vnHPkGeTRCiKtoz3Jj0IU-rito3tV39LL3JalhdrwuQ34xSBM-xLxUF9m3brN4hg2nyCVPqBbUUga3tupgtQig"
	}
	return "https://pbs.twimg.com/profile_images/1414880257631416321/s0pDGoih_400x400.jpg"
}

func PrefCode(cityName string) string {
	/*var data *domain.City
	if err:=db.QueryRow("SELECT * FROM City WHERE City.cityName=$1",cityName).Scan(&data.CityName,&data.ID);err!=nil{
		log.Fatalf("Couldnot queryRow:%v",err)
	}
	return data*/
	return domain.CodeCity[cityName]
}