
package main

import (
	"log"
	"milena/graph"
	"net/http"
	"os"
	"milena/package/database"
	"github.com/joho/godotenv"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/gorilla/websocket"
	_ "github.com/joho/godotenv/autoload"
)

const defaultPort = "8080"

func init() {
    // loads values from .env into the system
    if err := godotenv.Load("../.env"); err != nil {
	log.Print("No .env file found")
    }
}

func main() {
	var dataBase = database.Connect()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	// Chi router
	router := chi.NewRouter()
	// CORS
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://studio.apollographql.com", "http://localhost:8080"},
		AllowCredentials: true,
		AllowedMethods: []string{"GET", "POST"},
		Debug:            false,
	}).Handler)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: dataBase}}))

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

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
