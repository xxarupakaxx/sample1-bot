package main

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/xxarupakaxx/sample1-bot/model"
	"log"
)

func debugFlex(bot *linebot.Client, event *linebot.Event)  {
	data:=model.GetWeather("140010")
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
					Text: nil,
					Contents:   []*linebot.SpanComponent{
						{
							Type:       linebot.FlexComponentTypeSpan,
							Text:       "最高気温 : " + data.Forecasts[0].Temperature.Max.Celsius + "℃",
							Size:       linebot.FlexTextSizeTypeXxl,
							Weight:     linebot.FlexTextWeightTypeBold,
							Color:      "#fc0703",
							Decoration: linebot.FlexTextDecorationTypeUnderline,
						},
						{
							Type:       linebot.FlexComponentTypeSpan,
							Text:       "最低気温 : " + data.Forecasts[0].Temperature.Min.Celsius + "℃",
							Size:       linebot.FlexTextSizeTypeSm,
							Color:      "#03befc",
						},
					},
					Flex:       linebot.IntPtr(2),
					Size:       linebot.FlexTextSizeTypeSm,
					Wrap:       false,
					//Action:     nil,
					MaxLines:   linebot.IntPtr(2),
				},
				&linebot.TextComponent{
					Type:       linebot.FlexComponentTypeText,
					Text:       data.Description.Text,
					//Contents:   nil,
					Flex:       linebot.IntPtr(3),
					Size:       linebot.FlexTextSizeTypeSm,
					Wrap:       false,
					//Color:      "",
					//Action:     nil,
					MaxLines:   linebot.IntPtr(5),
				},
			},
			BorderColor:     "#5cd8f7",
			//Action:          nil,
		},
		Footer:    &linebot.BoxComponent{
			Type:            linebot.FlexComponentTypeBox,
			Layout:          linebot.FlexBoxLayoutTypeBaseline,
			Contents:        []linebot.FlexComponent{
				&linebot.ButtonComponent{
					Type:    linebot.FlexComponentTypeButton,
					Action:  linebot.NewURIAction("Weather URL", data.Link),
					Style:   linebot.FlexButtonStyleTypeLink,
					Color:   "#5cf7ac",
				},
			},
			CornerRadius:    linebot.FlexComponentCornerRadiusTypeMd,
			//BackgroundColor: "",
			BorderColor:     "#5cf7ac",
			//Action:          nil,
		},
		Styles:    &linebot.BubbleStyle{
			Header: &linebot.BlockStyle{
				Separator:      true,
				
			},
			Hero:   &linebot.BlockStyle{
				Separator:      true,

			},
			Body:   &linebot.BlockStyle{
				Separator:      true,

			},
			Footer: &linebot.BlockStyle{
				Separator:      false,

			},
		},
	})
	if _,err:=bot.ReplyMessage(event.ReplyToken,resp).Do();err != nil {
		log.Fatalf("debug.go error :%v",err)
	}
}
