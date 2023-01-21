package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/moritztng/perceptionOS/data"
	"github.com/moritztng/perceptionOS/docker/api/graph"
	"github.com/moritztng/perceptionOS/messaging"
)

const defaultPort = "8080"

func main() {
	database := data.Open("images.db")
	messageProducer := messaging.NewProducer("127.0.0.1:4150")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: &database, MessageProducer: &messageProducer}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
