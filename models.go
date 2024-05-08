package main

import (
	"time"
	"github.com/google/uuid"

	"github.com/Numeez/rssAgg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	Name      string `json:"name"`
	ApiKey    string `json:"api_key"`

}

func databaseUser(user database.User) User{
	return User{
		ID: user.ID,
		CreatedAt: user.CreatedAt,
		UpdateAt: user.UpdateAt,
		Name: user.Name,
		ApiKey: user.ApiKey,

	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	Name      string `json:"name"`
	Url    string `json:"url"`
	UserID uuid.UUID `json:"user_id"`

}
type FeedFollower struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	UserID uuid.UUID `json:"user_id"`
	FeedId uuid.UUID `json:"feed_id"`

}
func databaseFeed(feed database.Feed)Feed{
	return Feed{
		ID: feed.ID,
		UserID: feed.UserID,
		UpdateAt: feed.UpdateAt,
		Name: feed.Name,
		Url: feed.Url,
		CreatedAt: feed.CreatedAt,
	}
}
func databaseFeedsToFeeds(feeds [] database.Feed)[]Feed{
	allFeeds:= []Feed{} 
	for _,feed :=range feeds{
		allFeeds = append(allFeeds,databaseFeed(feed))
	}
	return allFeeds
}
func databaseFeedFollowerToFeedFollower(feedfollower database.Feedfollower)FeedFollower{
	return FeedFollower{
		ID: feedfollower.FeedID,
		UserID: feedfollower.UserID,
		UpdateAt: feedfollower.UpdateAt,
		CreatedAt: feedfollower.CreatedAt,
		FeedId: feedfollower.FeedID,
	}
}
 func databaseGetFeedFollowersList(feedfollowers []database.Feedfollower)[]FeedFollower{
	ListOfFeedFollowers:= []FeedFollower{}
	for _,feedFollower:= range feedfollowers{
		ListOfFeedFollowers=append(ListOfFeedFollowers,databaseFeedFollowerToFeedFollower(feedFollower))
	}
	return ListOfFeedFollowers
 }

 func getUsersList(userList []database.User) [] User{
	users:=[]User{}
	for _,user:=range userList{
		users = append(users,databaseUser(user) )
	}
	return users
 }
