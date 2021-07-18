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
	db:=model.DBConnect()
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
				text:=user.DisplayName+"がtamedaと送信しました"
				bot.ReplyMessage(event.ReplyToken,linebot.NewFlexMessage("tameda",
					&linebot.BubbleContainer{
						Type: linebot.FlexContainerTypeBubble,
						Body: &linebot.BoxComponent{
							Type:            linebot.FlexComponentTypeBox,
							Layout:          linebot.FlexBoxLayoutTypeBaseline,
							Contents:        []linebot.FlexComponent{
								&linebot.TextComponent{
									Type:       linebot.FlexComponentTypeText,
									Text:       text,
									Gravity:    linebot.FlexComponentGravityTypeBottom,
									Color:      "#e76565",
									Style:      linebot.FlexTextStyleTypeItalic,
									Decoration: linebot.FlexTextDecorationTypeUnderline,
								},
								&linebot.TextComponent{
									Type:       linebot.FlexComponentTypeText,
									Text:       text,
									Contents:   []*linebot.SpanComponent{
										{
											Type:   linebot.FlexComponentTypeSpan,
											Text:   "tamedakouya",
											Size:   linebot.FlexTextSizeTypeXxl,
											Weight: linebot.FlexTextWeightTypeBold,
											Color:  "#91f89d",
											Style:  linebot.FlexTextStyleTypeItalic,
											Decoration: linebot.FlexTextDecorationTypeUnderline,
																				},
									},
									Flex:       linebot.IntPtr(1),
									Margin:     linebot.FlexComponentMarginTypeSm,
									Size:       linebot.FlexTextSizeTypeSm,
									Align:      linebot.FlexComponentAlignTypeEnd,
									Gravity:    linebot.FlexComponentGravityTypeBottom,
									Wrap:       true,
									Weight:     linebot.FlexTextWeightTypeBold,
									Color:      "#ebf891",
									Action:     linebot.NewMessageAction("weather","weather 140010"),
									Style:      linebot.FlexTextStyleTypeItalic,
									Decoration: linebot.FlexTextDecorationTypeUnderline,
									MaxLines:   linebot.IntPtr(5),
								},
							},
							CornerRadius:  linebot.FlexComponentCornerRadiusTypeXl  ,
							BackgroundColor: "#6de765",
							BorderColor:     "#3BAF75",
							Action:         linebot.NewURIAction("tameda","https://twitter.com/TamerNazeda") ,
						},
					},
				)).Do()
			}
			if message.Text=="tamedaVerticalTopWrap" {
				text:=user.DisplayName+"がtamedaVerticalTopWrapと送信しました"
				bot.ReplyMessage(event.ReplyToken,linebot.NewFlexMessage("tameda1",
					&linebot.BubbleContainer{
						Type: linebot.FlexContainerTypeBubble,
						Body: &linebot.BoxComponent{
							Type:            linebot.FlexComponentTypeBox,
							Layout:          linebot.FlexBoxLayoutTypeVertical,
							Contents:        []linebot.FlexComponent{
								&linebot.TextComponent{
									Type:       linebot.FlexComponentTypeText,
									Text:       text,
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
							CornerRadius:  linebot.FlexComponentCornerRadiusTypeXl  ,
							//BackgroundColor: "#6de765",
							BorderColor:     "#3BAF75",
							Action:         linebot.NewURIAction("tameda","https://twitter.com/TamerNazeda") ,
						},
					},
				)).Do()
			}
			if message.Text=="tamedaHorizontalTopWrap" {
				text:=user.DisplayName+"がtamedaHorizontalTopWrapと送信しました"
				bot.ReplyMessage(event.ReplyToken,linebot.NewFlexMessage("tameda1",
					&linebot.BubbleContainer{
						Type: linebot.FlexContainerTypeBubble,
						Body: &linebot.BoxComponent{
							Type:            linebot.FlexComponentTypeBox,
							Layout:          linebot.FlexBoxLayoutTypeHorizontal,
							Contents:        []linebot.FlexComponent{
								&linebot.TextComponent{
									Type:       linebot.FlexComponentTypeText,
									Text:       text,
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
							CornerRadius:  linebot.FlexComponentCornerRadiusTypeXl  ,
							//BackgroundColor: "#6de765",
							BorderColor:     "#3BAF75",
							Action:         linebot.NewURIAction("tameda","https://twitter.com/TamerNazeda") ,
						},
					},
				)).Do()
			}
			if message.Text=="tameda1" {
				text:=user.DisplayName+"がtameda1と送信しました"
				bot.ReplyMessage(event.ReplyToken,linebot.NewFlexMessage("tameda1",
					&linebot.BubbleContainer{
						Type: linebot.FlexContainerTypeBubble,
						Body: &linebot.BoxComponent{
							Type:            linebot.FlexComponentTypeBox,
							Layout:          linebot.FlexBoxLayoutTypeVertical,
							Contents:        []linebot.FlexComponent{
								&linebot.TextComponent{
									Type:       linebot.FlexComponentTypeText,
									Text:       text,
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
									Wrap:       false,
									Weight:     linebot.FlexTextWeightTypeBold,
									Color:      "#6443d9",
									Action:     linebot.NewMessageAction("weather","weather 140010"),
									Style:      linebot.FlexTextStyleTypeItalic,
									Decoration: linebot.FlexTextDecorationTypeNone,
									MaxLines:   linebot.IntPtr(5),
								},
							},
							CornerRadius:  linebot.FlexComponentCornerRadiusTypeXl  ,
							//BackgroundColor: "#6de765",
							BorderColor:     "#3BAF75",
							Action:         linebot.NewURIAction("tameda","https://twitter.com/TamerNazeda") ,
						},
					},
				)).Do()
			}
			if message.Text == "Ctameda" {
				text:="cccc"
				bot.ReplyMessage(event.ReplyToken,linebot.NewFlexMessage("Ctameda",
					&linebot.CarouselContainer{
						Type:     linebot.FlexContainerTypeCarousel,
						Contents: []*linebot.BubbleContainer{
							{
								Type: linebot.FlexContainerTypeBubble,
								Body: &linebot.BoxComponent{
									Type:            linebot.FlexComponentTypeBox,
									Layout:          linebot.FlexBoxLayoutTypeVertical,
									Contents:        []linebot.FlexComponent{
										&linebot.TextComponent{
											Type:       linebot.FlexComponentTypeText,
											Text:       text,
											Gravity:    linebot.FlexComponentGravityTypeBottom,
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
											Align:      linebot.FlexComponentAlignTypeStart,
											Gravity:    linebot.FlexComponentGravityTypeTop,
											Wrap:       true,
											Weight:     linebot.FlexTextWeightTypeBold,
											Color:      "#6443d9",
											Action:     linebot.NewMessageAction("weather","weather 140010"),
											Style:      linebot.FlexTextStyleTypeItalic,
											Decoration: linebot.FlexTextDecorationTypeNone,
											MaxLines:   linebot.IntPtr(5),
										},
									},
									CornerRadius:  linebot.FlexComponentCornerRadiusTypeSm  ,
									//BackgroundColor: "#6de765",
									BorderColor:     "#3BAF75",
									Action:         linebot.NewURIAction("tameda","https://twitter.com/TamerNazeda") ,
								},
							},
							{
								Type: linebot.FlexContainerTypeBubble,
								Body: &linebot.BoxComponent{
									Type:            linebot.FlexComponentTypeBox,
									Layout:          linebot.FlexBoxLayoutTypeBaseline,
									Contents:        []linebot.FlexComponent{
										&linebot.TextComponent{
											Type:       linebot.FlexComponentTypeText,
											Text:       text,
											Gravity:    linebot.FlexComponentGravityTypeCenter,
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
													Size:   linebot.FlexTextSizeTypeXxl,
													Weight: linebot.FlexTextWeightTypeBold,
													Color:  "#91f89d",
													Style:  linebot.FlexTextStyleTypeItalic,
													Decoration: linebot.FlexTextDecorationTypeUnderline,
												},
												{
													Type:   linebot.FlexComponentTypeSpan,
													Text:   "140010",
													Size:   linebot.FlexTextSizeTypeSm,
													Weight: linebot.FlexTextWeightTypeRegular,
													Color:  "#91f89d",
													Style:  linebot.FlexTextStyleTypeNormal,
													Decoration: linebot.FlexTextDecorationTypeNone,
												},
											},
											Flex:       linebot.IntPtr(0),
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
									CornerRadius:  linebot.FlexComponentCornerRadiusTypeXl  ,
									//BackgroundColor: "#6de765",
									BorderColor:     "#3BAF75",
									Action:         linebot.NewURIAction("tameda","https://twitter.com/TamerNazeda") ,
								},
							},
						},
					},
				)).Do()
			}
			if message.Text=="tamedaBotton" {
				text:=user.DisplayName+"がtameda1と送信しました"
				bot.ReplyMessage(event.ReplyToken,linebot.NewFlexMessage("tameda1",
					&linebot.BubbleContainer{
						Type: linebot.FlexContainerTypeBubble,
						Body: &linebot.BoxComponent{
							Type:            linebot.FlexComponentTypeBox,
							Layout:          linebot.FlexBoxLayoutTypeVertical,
							Contents:        []linebot.FlexComponent{
								&linebot.TextComponent{
									Type:       linebot.FlexComponentTypeText,
									Text:       text,
									Gravity:    linebot.FlexComponentGravityTypeBottom,
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
									Align:      linebot.FlexComponentAlignTypeStart,
									Gravity:    linebot.FlexComponentGravityTypeTop,
									Wrap:       true,
									Weight:     linebot.FlexTextWeightTypeBold,
									Color:      "#6443d9",
									Action:     linebot.NewMessageAction("weather","weather 140010"),
									Style:      linebot.FlexTextStyleTypeItalic,
									Decoration: linebot.FlexTextDecorationTypeNone,
									MaxLines:   linebot.IntPtr(5),
								},
							},
							CornerRadius:  linebot.FlexComponentCornerRadiusTypeSm  ,
							//BackgroundColor: "#6de765",
							BorderColor:     "#3BAF75",
							Action:         linebot.NewURIAction("tameda","https://twitter.com/TamerNazeda") ,
						},
					},
				)).Do()
			}
			if message.Text=="tamedaBase" {
				text:=user.DisplayName+"がtameda1と送信しました"
				bot.ReplyMessage(event.ReplyToken,linebot.NewFlexMessage("tameda1",
					&linebot.BubbleContainer{
						Type: linebot.FlexContainerTypeBubble,
						Body: &linebot.BoxComponent{
							Type:            linebot.FlexComponentTypeBox,
							Layout:          linebot.FlexBoxLayoutTypeBaseline,
							Contents:        []linebot.FlexComponent{
								&linebot.TextComponent{
									Type:       linebot.FlexComponentTypeText,
									Text:       text,
									Gravity:    linebot.FlexComponentGravityTypeCenter,
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
											Size:   linebot.FlexTextSizeTypeXxl,
											Weight: linebot.FlexTextWeightTypeBold,
											Color:  "#91f89d",
											Style:  linebot.FlexTextStyleTypeItalic,
											Decoration: linebot.FlexTextDecorationTypeUnderline,
										},
										{
											Type:   linebot.FlexComponentTypeSpan,
											Text:   "140010",
											Size:   linebot.FlexTextSizeTypeSm,
											Weight: linebot.FlexTextWeightTypeRegular,
											Color:  "#91f89d",
											Style:  linebot.FlexTextStyleTypeNormal,
											Decoration: linebot.FlexTextDecorationTypeNone,
										},
									},
									Flex:       linebot.IntPtr(0),
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
							CornerRadius:  linebot.FlexComponentCornerRadiusTypeXl  ,
							//BackgroundColor: "#6de765",
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
	defer db.Close()
	return c.String(http.StatusOK,"OK")

}

