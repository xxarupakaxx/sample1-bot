package model

import (
	"encoding/json"
	"github.com/xxarupakaxx/sample1-bot/domain"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func SerarchYoutubeChannel(channelName string) ([]domain.Video, error) {
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
}
