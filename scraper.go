package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/Numeez/rssAgg/internal/database"
)

func startScraping (db *database.Queries,concurrency int,timeBetweenRequest time.Duration){
	log.Printf("Scraping on %v coroutined every %s duration",concurrency,timeBetweenRequest)
	ticker:=time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C{
		feeds,err:=db.GetNextFeedsToFetch(context.Background(),int32(concurrency))
		if err !=nil{
			log.Println("Error fetching feeds : ",err)
			continue
		}
		
		wg := &sync.WaitGroup{}

		for _,feed :=range feeds{
			wg.Add(1)
			go scrapeFeed(wg,db,feed)
		}
		wg.Wait()

	}
}

func scrapeFeed(wg *sync.WaitGroup,db *database.Queries, feed database.Feed){
	defer wg.Done()

	_,err:=db.MarkFeedFetch(context.Background(),feed.ID)
	if err !=nil{
		log.Println("Error marking feed as a fetched one : ",err)
	}
	rssFeed,err:=URLToFeed(feed.Url)
	if err !=nil{
		log.Println("Error fetching the rss feed from internet : ",err)
	}
	for _,item :=range rssFeed.Channel.Item{
		log.Println("Found post",item.Title,"on feed",feed.Name)
	}

	log.Printf("Feed %s collected , %v posts found",feed.Name,len(rssFeed.Channel.Item))
}