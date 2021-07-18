package main

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

func debugFlex(bot *linebot.Client, event *linebot.Event)  {
	/*var (
		layoutV = "Vertical"
		layoutH = "Horizontal"
	)

	msg:=event.Message.(*linebot.TextMessage)
	msgSplit:=strings.Split(msg.Text," ")*/
	
	
	resp:=linebot.NewFlexMessage(
		"debug",
		&linebot.BubbleContainer{
			Type:      linebot.FlexContainerTypeBubble,
			Size:      linebot.FlexBubbleSizeTypeGiga,
			Direction: linebot.FlexBubbleDirectionTypeLTR,
			Header:    &linebot.BoxComponent{
				Type:            linebot.FlexComponentTypeBox,
				Layout:          linebot.FlexBoxLayoutTypeBaseline,
				Contents:        []linebot.FlexComponent{
					&linebot.TextComponent{
						Type:       linebot.FlexComponentTypeText,
						Text:       "aaa",
						Gravity:    linebot.FlexComponentGravityTypeTop,
						Color:      "#473232",
						Style:      linebot.FlexTextStyleTypeItalic,
						Decoration: linebot.FlexTextDecorationTypeUnderline,
					},
					&linebot.TextComponent{
						Type:       linebot.FlexComponentTypeText,
						Text:       "text",
						Contents:   []*linebot.SpanComponent{
							{
								Type:   linebot.FlexComponentTypeSpan,
								Text:   "weather",
								Size:   linebot.FlexTextSizeTypeSm,
								Weight: linebot.FlexTextWeightTypeBold,
								Color:  "#91f89d",
								Style:  linebot.FlexTextStyleTypeItalic,
								Decoration: linebot.FlexTextDecorationTypeUnderline,
							},
							{
								Type:   linebot.FlexComponentTypeSpan,
								Text:   "140010",
								Size:   linebot.FlexTextSizeTypeLg,
								Weight: linebot.FlexTextWeightTypeRegular,
								Color:  "#91f89d",
								Style:  linebot.FlexTextStyleTypeNormal,
								Decoration: linebot.FlexTextDecorationTypeNone,
							},
						},
						Flex:       linebot.IntPtr(3),
						Margin:     linebot.FlexComponentMarginTypeSm,
						Size:       linebot.FlexTextSizeTypeSm,
						Align:      linebot.FlexComponentAlignTypeEnd,
						Gravity:    linebot.FlexComponentGravityTypeBottom,
						Wrap:       true,
						Weight:     linebot.FlexTextWeightTypeBold,
						Color:      "#6443d9",
						Action:     linebot.NewMessageAction("weather","weather 140010"),
						Style:      linebot.FlexTextStyleTypeItalic,
						Decoration: linebot.FlexTextDecorationTypeNone,
						MaxLines:   linebot.IntPtr(5),
					},
				},
				Flex:            nil,
				Spacing:         linebot.FlexComponentSpacingTypeMd,
				Margin:          linebot.FlexComponentMarginTypeSm,
				CornerRadius:    linebot.FlexComponentCornerRadiusTypeXxl,
				BackgroundColor: "#65e7a2",
				BorderColor:     "#de65e7",
				Action:          linebot.NewMessageAction("tameda","weather 140010"),
			},
			Hero:      &linebot.ImageComponent{
				Type:            linebot.FlexComponentTypeImage,
				URL:             "https://pbs.twimg.com/profile_images/1364204318790836232/dgyZpvy6_400x400.jpg",
				Flex:            nil,
				Margin:          linebot.FlexComponentMarginTypeSm,
				Align:           linebot.FlexComponentAlignTypeEnd,
				Gravity:         linebot.FlexComponentGravityTypeTop,
				Size:            linebot.FlexImageSizeTypeXl,
				AspectRatio:     linebot.FlexImageAspectRatioType1to1,
				AspectMode:      linebot.FlexImageAspectModeTypeFit,
				BackgroundColor: "#91c4fa",
				Action:          linebot.NewURIAction("tamedaAction", "https://twitter.com/TamerNazeda"),
			},
			Body:      &linebot.BoxComponent{
				Type:            linebot.FlexComponentTypeBox,
				Layout:          linebot.FlexBoxLayoutTypeHorizontal,
				Contents:        []linebot.FlexComponent{
					&linebot.TextComponent{
						Type:       linebot.FlexComponentTypeText,
						Text:       "BODY",
						Gravity:    linebot.FlexComponentGravityTypeTop,
						Color:      "#473232",
						Style:      linebot.FlexTextStyleTypeItalic,
						Decoration: linebot.FlexTextDecorationTypeUnderline,
					},
					&linebot.TextComponent{
						Type:       linebot.FlexComponentTypeText,
						Text:       "text",
						Contents:   []*linebot.SpanComponent{
							{
								Type:   linebot.FlexComponentTypeSpan,
								Text:   "weather",
								Size:   linebot.FlexTextSizeTypeSm,
								Weight: linebot.FlexTextWeightTypeBold,
								Color:  "#91f89d",
								Style:  linebot.FlexTextStyleTypeItalic,
								Decoration: linebot.FlexTextDecorationTypeUnderline,
							},
							{
								Type:   linebot.FlexComponentTypeSpan,
								Text:   "140010",
								Size:   linebot.FlexTextSizeTypeLg,
								Weight: linebot.FlexTextWeightTypeRegular,
								Color:  "#91f89d",
								Style:  linebot.FlexTextStyleTypeNormal,
								Decoration: linebot.FlexTextDecorationTypeNone,
							},
						},
						Flex:       linebot.IntPtr(3),
						Margin:     linebot.FlexComponentMarginTypeSm,
						Size:       linebot.FlexTextSizeTypeSm,
						Align:      linebot.FlexComponentAlignTypeEnd,
						Gravity:    linebot.FlexComponentGravityTypeBottom,
						Wrap:       true,
						Weight:     linebot.FlexTextWeightTypeBold,
						Color:      "#6443d9",
						Action:     linebot.NewMessageAction("weather","weather 140010"),
						Style:      linebot.FlexTextStyleTypeItalic,
						Decoration: linebot.FlexTextDecorationTypeNone,
						MaxLines:   linebot.IntPtr(5),
			},
				},
			},
			Footer:    &linebot.BoxComponent{
				Type:            linebot.FlexComponentTypeBox,
				Layout:          linebot.FlexBoxLayoutTypeBaseline,
				Contents:        []linebot.FlexComponent{
					&linebot.TextComponent{
						Type:       linebot.FlexComponentTypeText,
						Text:       "aaa",
						Gravity:    linebot.FlexComponentGravityTypeTop,
						Color:      "#473232",
						Style:      linebot.FlexTextStyleTypeItalic,
						Decoration: linebot.FlexTextDecorationTypeUnderline,
					},
					&linebot.TextComponent{
						Type:       linebot.FlexComponentTypeText,
						Text:       "text",
						Contents:   []*linebot.SpanComponent{
							{
								Type:   linebot.FlexComponentTypeSpan,
								Text:   "weather",
								Size:   linebot.FlexTextSizeTypeSm,
								Weight: linebot.FlexTextWeightTypeBold,
								Color:  "#91f89d",
								Style:  linebot.FlexTextStyleTypeItalic,
								Decoration: linebot.FlexTextDecorationTypeUnderline,
							},
							{
								Type:   linebot.FlexComponentTypeSpan,
								Text:   "140010",
								Size:   linebot.FlexTextSizeTypeLg,
								Weight: linebot.FlexTextWeightTypeRegular,
								Color:  "#91f89d",
								Style:  linebot.FlexTextStyleTypeNormal,
								Decoration: linebot.FlexTextDecorationTypeNone,
							},
						},
						Flex:       linebot.IntPtr(3),
						Margin:     linebot.FlexComponentMarginTypeSm,
						Size:       linebot.FlexTextSizeTypeSm,
						Align:      linebot.FlexComponentAlignTypeEnd,
						Gravity:    linebot.FlexComponentGravityTypeBottom,
						Wrap:       true,
						Weight:     linebot.FlexTextWeightTypeBold,
						Color:      "#6443d9",
						Action:     linebot.NewMessageAction("weather","weather 140010"),
						Style:      linebot.FlexTextStyleTypeItalic,
						Decoration: linebot.FlexTextDecorationTypeNone,
						MaxLines:   linebot.IntPtr(5),
					},
				},
				Flex:            nil,
				Spacing:         linebot.FlexComponentSpacingTypeMd,
				Margin:          linebot.FlexComponentMarginTypeSm,
				CornerRadius:    linebot.FlexComponentCornerRadiusTypeXxl,
				BackgroundColor: "#65e7a2",
				BorderColor:     "#de65e7",
				Action:          linebot.NewMessageAction("tameda","weather 140010"),
			},
			Styles:    &linebot.BubbleStyle{
				Header: &linebot.BlockStyle{
					BackgroundColor: "#edb12f",
					Separator:       true,
					SeparatorColor:  "#ff6347",
				},
				Hero:   &linebot.BlockStyle{
					BackgroundColor: "#6647ff",
					Separator:       false,
					SeparatorColor:  "",
				},
				Body:   &linebot.BlockStyle{
					BackgroundColor: "#c01dcf",
					Separator:       false,
					SeparatorColor:  "",
				},
				Footer: &linebot.BlockStyle{
					BackgroundColor: "#d11959",
					Separator:       false,
					SeparatorColor:  "",
				},
			},
		},
	)
	if _, err := bot.ReplyMessage(event.ReplyToken, resp).Do(); err != nil {
		log.Fatalf("debug.go error:%v",err)
	}
}
