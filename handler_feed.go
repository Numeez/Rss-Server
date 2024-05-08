package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Numeez/rssAgg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request,user database.User){
	type parameters struct{
		Name string `json"name"`
		URL string `json"url"`
	}
	decoder:=json.NewDecoder(r.Body)
	params:=parameters{}
	err:=decoder.Decode(&params)
	if err !=nil{
		respondWithError(w,400,fmt.Sprintln("Error parsing json : ",err))
		return
	}
	feed,err:= apiCfg.DB.CreateFeed(r.Context(),database.CreateFeedParams{
		ID:uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdateAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
	})
	if err !=nil{
		respondWithError(w,400,fmt.Sprintln("Could not create feed : ",err))
		return
	}

	respondWithJson(w,201,databaseFeed(feed))


}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request){
	feeds,err:= apiCfg.DB.GetFeeds(r.Context())
	if err !=nil{
		respondWithError(w,400,fmt.Sprintln("Could not create feed : ",err))
		return
	}

	respondWithJson(w,201,databaseFeedsToFeeds(feeds))


}