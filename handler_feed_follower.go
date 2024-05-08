package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Numeez/rssAgg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)


func (apiCfg *apiConfig) handlerCreateFeedFollower(w http.ResponseWriter, r *http.Request, user database.User){
	type parameters struct{
		Feed_Id uuid.UUID `json"feed_id"`
	}

	params:=parameters{}
	decoder:=json.NewDecoder(r.Body)
	err:=decoder.Decode(&params)
	if err !=nil{
		respondWithError(w,404,fmt.Sprintln("Error parsing json : ",err))
		return
	}
	feedfollower,err:=apiCfg.DB.CreateFeedFollower(r.Context(),database.CreateFeedFollowerParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdateAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.Feed_Id,

	})
	if err !=nil{
		respondWithError(w,404,fmt.Sprintln("Feed follower not added : ",err))
		return
	}

	respondWithJson(w,201,databaseFeedFollowerToFeedFollower(feedfollower))


}

func (apiCfg *apiConfig) handlerGetFeedFollower(w http.ResponseWriter, r *http.Request,user database.User){
	
	feedfollowers,err:= apiCfg.DB.GetFeedFollows(r.Context(),user.ID)
	if err !=nil{
		respondWithError(w,404,fmt.Sprintln("Feed follower not added : ",err))
		return
	}

	respondWithJson(w,201,databaseGetFeedFollowersList(feedfollowers))

}

func (apiCfg *apiConfig) handlerDeleteFeedFollower(w http.ResponseWriter,r *http.Request,user database.User){
	feedFollowIdString:=chi.URLParam(r,"feedFollowerId")	
	feedFollowId,err:=uuid.Parse(feedFollowIdString)
	if err!=nil{
		respondWithError(w,400,fmt.Sprintln("Could not parse feed follow id : ",err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollows(r.Context(),database.DeleteFeedFollowsParams{
		ID: feedFollowId,
		UserID: user.ID,
	})
	if err!=nil{
		respondWithError(w,400,fmt.Sprintln("Feed follow delete operation failed",err))
		return
	}
	respondWithJson(w,200,struct{}{})

}

