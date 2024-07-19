package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	graph "github.com/censoredplanet/cp-api/internal/api/graphql"
	"github.com/censoredplanet/cp-api/internal/api/graphql/generated"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// TODO: init database connection, middleware integration and slack connection

	router := mux.NewRouter()

	conf := generated.Config{Resolvers: &graph.Resolver{}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(conf))

	router.Handle("/query", srv)
	router.Handle("/", playground.Handler("GraphQL Playground", "/query"))

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("No port has been specified in .env")
	}
	log.Printf("Server is running on localhost:%s", port)
	if err := (http.ListenAndServe(fmt.Sprintf(":%s", port), router)); err != nil {
		log.Println(err)
	}
}
