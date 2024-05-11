package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct{
	Channel struct{
		Title string `xml:"title"`
		Link string `xml:"link"`
		Description string `xml:"description"`
		PubDate string `xml:"pubdate"`
		Language string `xml:"language"`
		Item []RSSItem `xml:"item"`
	} `xml:"channel"`

}

type RSSItem struct{
	Title string `xml:"title"`
	Link string `xml:"link"`
	Description string `xml:"description"`
	PubDate string `xml:"pubdate"`
}

func URLToFeed(url string) (RSSFeed,error){
	rssFeed:= RSSFeed{}
	httpClient := http.Client{
		Timeout: 10* time.Second,
	}

	res,err:= httpClient.Get(url)

	if err !=nil{
		return RSSFeed{} ,err
	}

	defer res.Body.Close()

	data,err:=io.ReadAll(res.Body)
	if err !=nil{
		return RSSFeed{} ,err
	}

	xml.Unmarshal(data,&rssFeed)

	return rssFeed,nil

}