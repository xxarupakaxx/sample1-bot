package model

import (
	"encoding/json"
	"github.com/xxarupakaxx/sample1-bot/domain"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func serarchYoutubeList(apiKey string) ([]domain.video, error) {
	url:="https://www.googleapis.com/youtube/v3/search"

	req,err:=http.NewRequest("GET",url,nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("content-type", "application/json")
	params:=req.URL.Query()
	params.Add("key",apiKey)
	params.Add("q","amazarashi Official YouTube Channel")
	params.Add("part","snippet, id")

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
	data:=make([]*videoInfo,0)
	if err = json.Unmarshal([]byte(s), &data); err != nil {
		log.Fatal(err)
	}
	videos:=make([]video,0,len(data))
	for _,v:=range data{
		videos=append(videos,video{
			VideoID:      v.ID.VideoID,
			PublishedAt:  v.Snippet.PublishedAt,
			ChannelID:    v.Snippet.ChannelID,
			Title:        v.Snippet.Title,
			ChannelTitle: v.Snippet.ChannelTitle,
		})
	}
	return videos,err
}
