package model

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/xxarupakaxx/sample1-bot/domain"
	"log"
	"time"
)

/*func SerarchYoutubeChannel(channelName string,s *youtube.Service) ([]domain.Video, error) {
	list:=youtube.NewChannelsService(s)
	result,err:=list.List([]string{"id","snippet"}).ForUsername(channelName).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Items)
	url:="https://www.googleapis.com/youtube/v3/search"

	req,err:=http.NewRequest("GET",url,nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("content-type", "application/json")
	params:=req.URL.Query()
	params.Add("key",os.Getenv("YOUTUBE_APIKEY"))
	params.Set("part","id,snippet")
	params.Set("q",channelName)

	req.URL.RawQuery=params.Encode()

	timeout:=time.Duration(5*time.Second)
	client:=&http.Client{
		Timeout: timeout,
	}

	res,err:=client.Do(req)
	if err != nil{
		log.Fatal(err)
	}

	defer res.Body.Close()

	body,_:=ioutil.ReadAll(res.Body)

	index:=strings.Index(string(body),"items")
	s:=string(body)[index+7: len(string(body))-2]
	//fmt.Println(s)
	data:=make([]*domain.VideoInfo,0)
	if err = json.Unmarshal([]byte(s), &data); err != nil {
		log.Fatal(err)
	}
	videos:=make([]domain.Video,0,len(data))
	for _,v:=range data{
		videos=append(videos,domain.Video{
			VideoID:      v.ID.VideoID,
			PublishedAt:  v.Snippet.PublishedAt,
			ChannelID:    v.Snippet.ChannelID,
			Title:        v.Snippet.Title,
			ChannelTitle: v.Snippet.ChannelTitle,
		})
	}
	return videos,err
}*/

func UserListGET(bot *linebot.Client,event *linebot.Event) []domain.Video {

	user:=bot.GetProfile(event.Source.UserID)
	userProfile,err:= user.Do()
	if err != nil {
		log.Fatal(err)
	}
	db:=DBConnect()
	result,err:=db.Query("SELECT * FROM video ORDER BY user DESC ")
	if err != nil {
		log.Fatal(err)
	}
	videos:=[]domain.Video{}
	for result.Next() {
		var (
			videoId string
			published_at time.Time
			channelID string
			title string
			description string
			channelTitle string
		)

		result.Scan(&userProfile.UserID,&videoId,&published_at,&channelID,&title,&description,&channelTitle)
		videos=append(videos,domain.Video{
			VideoID:      videoId,
			PublishedAt:  published_at,
			ChannelID:    channelID,
			Title:        title,
			Description:  description,
			ChannelTitle: channelTitle,
		})
	}
	return videos
}
/*
func UserVideoPOST(bot *linebot.Client, event *linebot.Event) domain.Video {
	user:=bot.GetProfile(event.Source.UserID)
	userProfile,err:= user.Do()
	if err != nil {
		log.Fatal(err)
	}
	db:=DBConnect()

	result,err:=db.Query("INSERT INTO video (user,videoId,published_at,channelID,title,description,channelTitle) VALUES (?,?,?,?,?,?,?)",title,now,now)
}*/
