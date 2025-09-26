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
	database "github.com/censoredplanet/cp-api/internal/database"
	"github.com/censoredplanet/cp-api/internal/middleware"
	service "github.com/censoredplanet/cp-api/internal/services"
	"github.com/censoredplanet/cp-api/internal/slack"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

const version = "v0.3.0"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	slack := slack.NewSlack()
	slack.Info("version", version, "started")

	clickHouseClient, err := database.ClickHouseConnect()
	if err != nil {
		// slack.Fatal("main.go", "main", "ClickHouseConnect", err.Error())
		// log.Fatalf("Fatal: %s\n", err)
	}

	clickHouseRepo, err := database.NewClickHouse(&clickHouseClient)
	if err != nil {
		slack.Fatal("main.go", "main", "NewClickHouse", err.Error())
	}

	serviceRepo, err := service.NewService(slack, clickHouseRepo)
	if err != nil {
		slack.Fatal("main.go", "main", "NewService", err.Error())
	}

	router := mux.NewRouter()
	router.Use(middleware.Middleware)
	conf := generated.Config{Resolvers: &graph.Resolver{
		Service: serviceRepo,
	}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(conf))

	router.Handle("/query", cors.AllowAll().Handler(srv))
	router.Handle("/", playground.Handler("GraphQL Playground", "/query"))

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("No port has been specified in .env")
	}

	env := os.Getenv("ENV")
	log.Printf("Starting server in %s mode on port %s", env, port)

	var serverErr error
	address := fmt.Sprintf(":%s", port)

	log.Printf("Starting HTTP server on %s", address)
	serverErr = http.ListenAndServe(address, router)

	if serverErr != nil {
		slack.Fatal("main.go", "main", "ListenAndServe", serverErr.Error())
		log.Fatal(serverErr)
	}
}
