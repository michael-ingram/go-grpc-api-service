package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	//"github.com/99designs/gqlgen/graphql/playground"
	"go-api-service/configurations"
	"go-api-service/graph"
	"go-api-service/grpc/clients"
	"log"
	"net/http"
)

const defaultPort = "8080"

func main() {
	// Load configuration
	cfg := configurations.LoadConfiguration()

	// Set up gRPC clients
	navblueClient := grpc.NewNavblueClient(cfg.GrpcNavblueAddr)
	defer func() {
		if err := navblueClient.Close(); err != nil {
			log.Printf("error closing gRPC client: %v", err)
		}
	}()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			NavblueClient: navblueClient,
		},
	}))

	//http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("Server running at http://localhost:8080/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
