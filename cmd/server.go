package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/thegoldengator/APIv2/internal/apis"
	"github.com/thegoldengator/APIv2/internal/apis/twitch"
	"github.com/thegoldengator/APIv2/internal/config"
	"github.com/thegoldengator/APIv2/internal/database"
	"github.com/thegoldengator/APIv2/internal/events"
	"github.com/thegoldengator/APIv2/internal/gql/graph/generated"
	"github.com/thegoldengator/APIv2/internal/gql/resolvers"
	"github.com/thegoldengator/APIv2/internal/sse"
)

const defaultPort = "8080"

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

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:3000"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == "example.org"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	router.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	router.HandleFunc("/sse", func(w http.ResponseWriter, r *http.Request) {
		go func() {
			// Received browser disconnection
			<-r.Context().Done()
			println("Client disconnected")
		}()
		sse.Server.ServeHTTP(w, r)
	})

	/* s := gocron.NewScheduler(time.UTC)
	s.Every(5).Minutes().Do(routines.ViewCount())
	s.Every(24).Hours().Do(routines.Pfp()) */

	router.HandleFunc("/test/createstreams", func(w http.ResponseWriter, r *http.Request) {
		apis.Twitch.CreateStreams()
	})

	router.HandleFunc("/test/event", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var body map[string]interface{}
		errDecode := decoder.Decode(&body)
		if errDecode != nil {
			return
		}

		jsonStr, _ := json.Marshal(body)

		var toSend sse.SSEMessage
		json.Unmarshal(jsonStr, &toSend)
		// sse.PublishMessage(sse.SSEChannelEvents, toSend)
		events.StreamOnline(twitch.EventSubStreamOnlineEvent{
			BroadcasterUserID:    "208953352",
			BroadcasterUserLogin: "their0njew",
			BroadcasterUserName:  "THEIR0NJEW",
			Type:                 "live",
			StartedAt:            time.Now(),
		})
		fmt.Println("Published message")
	})

	/* router.HandleFunc("/eventsub", func(w http.ResponseWriter, r *http.Request) {
		routes.EventsubRecievedNotification(w, r)
	}).Methods("POST")
	*/
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
