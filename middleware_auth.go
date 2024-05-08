package main

import (
	"fmt"
	"net/http"

	"github.com/Numeez/rssAgg/internal/database"
	"github.com/Numeez/rssAgg/internal/database/auth"
)

type authHandler func(http.ResponseWriter,*http.Request,database.User)

func (cfg *apiConfig) middleWareAuth(handler authHandler) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		apiKey,err:= auth.GetApIKey(r.Header)
		if err !=nil {
			respondWithError(w,400,fmt.Sprintln("Auth error : ",err))
			return
		}
	
		user,err:=cfg.DB.GetUserByAPIKey(r.Context(),apiKey)
		if err!=nil{
			respondWithError(w,400,fmt.Sprintln("Could not get user : ",err))
			return
		}
		handler(w,r,user)	
	}
}