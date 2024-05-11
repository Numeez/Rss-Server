package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Numeez/rssAgg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)
type apiConfig struct{
	DB *database.Queries
}

func main(){
	feed,err:=URLToFeed("https://wagslane.dev/index.xml")
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println(feed)
	godotenv.Load()
	port:=os.Getenv("PORT")
	dbURL:=os.Getenv("DB_URL")
	if dbURL==""{
		log.Fatal("DB_URL is not found in the environment")
	}
	conn,err:=sql.Open("postgres",dbURL)
	if err!=nil{
		log.Fatal("Can't connect to the database")
	}
	db:=database.New(conn)
	apiCfg:= apiConfig{
		DB: db,
	}
	if port==""{
		log.Fatal("Port is not found in the environment")
	}
	log.Printf("Server starting on port %v",port)
	go startScraping(db,10,time.Minute)
	router:= chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
	AllowedOrigins: []string{"https://","http://"},
	AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
	AllowCredentials: false,
	AllowedHeaders: []string{"*"},
	MaxAge: 300,
	ExposedHeaders: []string{"Link"},	
	}))

	v1Router:= chi.NewRouter()
	v1Router.Get("/healthz",handlerReadiness)
	v1Router.Get("/err",handlerError)
	v1Router.Post("/users",apiCfg.handlerCreateUser)
	v1Router.Get("/users",apiCfg.middleWareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/feeds",apiCfg.middleWareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds",apiCfg.handlerGetFeeds)
	v1Router.Post("/feedfollowers",apiCfg.middleWareAuth(apiCfg.handlerCreateFeedFollower))
	v1Router.Get("/feedfollowers",apiCfg.middleWareAuth(apiCfg.handlerGetFeedFollower))
	v1Router.Delete("/feedfollowers/{feedFollowerId}",apiCfg.middleWareAuth(apiCfg.handlerDeleteFeedFollower))
	v1Router.Get("/getAllUsers",apiCfg.handlerGetAllUser)
	router.Mount("/v1",v1Router)


	srv:= &http.Server{
		Handler: router,
		Addr: ":"+port,
	}
	serverError :=srv.ListenAndServe()
	if serverError!=nil{
		log.Fatal(serverError)
	}
	
}