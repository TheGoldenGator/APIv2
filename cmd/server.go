package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/thegoldengator/APIv2/internal/apis"
	"github.com/thegoldengator/APIv2/internal/config"
	"github.com/thegoldengator/APIv2/internal/database"
	"github.com/thegoldengator/APIv2/internal/gql/graph/generated"
	"github.com/thegoldengator/APIv2/internal/gql/resolvers"
	"github.com/thegoldengator/APIv2/internal/routes"
	"github.com/thegoldengator/APIv2/internal/sse"
)

const defaultPort = "8000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	connectErr := database.Connect(config.Config.GetString("mongo_uri"))
	if connectErr != nil {
		panic(connectErr)
	}

	// Initialize SSE
	sse.Connect()

	errVrcLogin := apis.VRChat.Login(config.Config.GetString("vrc_username"), config.Config.GetString("vrc_password"))
	if errVrcLogin != nil {
		panic(errVrcLogin)
	}

	go func() {
		/* s := gocron.NewScheduler(time.UTC)
		s.Every(5).Minutes().Do(func() {
			err := routines.ViewCount()
			if err != nil {
				fmt.Println("Error updating view count", err)
			}
		})
		s.Every(24).Hours().Do(func() {
			err := routines.Pfp()
			fmt.Println("Error updating pfps", err)
		})
		s.Every(30).Seconds().Do(func() {
			sse.PublishPing(sse.SSEChannelEvents)
		})

		s.StartAsync() */
	}()

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:         []string{"http://localhost:8080", "http://localhost:3000", "https://thegoldengator.tv", "http://api.thegoldengator.tv"},
		AllowOriginRequestFunc: func(r *http.Request, origin string) bool { return true },
		AllowCredentials:       true,
		Debug:                  false,
		AllowedHeaders:         []string{"X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == "thegoldengator.tv" || r.Host == "localhost"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	router.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	router.HandleFunc("/sse", sse.SSEServer.ServeHTTP)
	router.Get("/vrc/world/{username}", routes.VRCWorld)
	router.Post("/eventsub", routes.EventsubRecievedNotification)

	router.HandleFunc("/test/createstreams", func(w http.ResponseWriter, r *http.Request) {
		apis.Twitch.CreateStreams()
	})

	/* router.HandleFunc("/test/event", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var body map[string]interface{}
		errDecode := decoder.Decode(&body)
		if errDecode != nil {
			return
		}

		jsonStr, _ := json.Marshal(body)
		fmt.Println(body)

		var toSend sse.SSEMessage
		json.Unmarshal(jsonStr, &toSend)
		//sse.PublishMessage(sse.SSEChannelEvents, toSend)
		if toSend.Event == sse.SSEMessageEventStreamOnline {
			events.StreamOnline(twitch.EventSubStreamOnlineEvent{
				BroadcasterUserID:    "208953352",
				BroadcasterUserLogin: "their0njew",
				BroadcasterUserName:  "THEIR0NJEW",
				Type:                 "live",
				StartedAt:            time.Now(),
			})
			return
		} else if toSend.Event == sse.SSEMessageEventChannelUpdate {
			events.ChannelUpdate(twitch.EventSubChannelUpdateEvent{
				BroadcasterUserID:    body["data"].(map[string]interface{})["broadcaster_user_id"].(string),
				BroadcasterUserLogin: body["data"].(map[string]interface{})["broadcaster_user_login"].(string),
				BroadcasterUserName:  body["data"].(map[string]interface{})["broadcaster_user_name"].(string),
				Title:                body["data"].(map[string]interface{})["title"].(string),
				Language:             body["data"].(map[string]interface{})["language"].(string),
				CategoryID:           body["data"].(map[string]interface{})["category_id"].(string),
				CategoryName:         body["data"].(map[string]interface{})["category_name"].(string),
				IsMature:             body["data"].(map[string]interface{})["is_mature"].(bool),
			})
		} else if toSend.Event == sse.SSEMessageEventStreamOffline {
			events.StreamOffline(twitch.EventSubStreamOfflineEvent{
				BroadcasterUserID:    "208953352",
				BroadcasterUserLogin: "their0njew",
				BroadcasterUserName:  "THEIR0NJEW",
			})
			return
		}

		fmt.Println("Published message")
	}) */

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
