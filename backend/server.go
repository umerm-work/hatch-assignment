package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/umerm-work/hatch-assignment/config"
	"github.com/umerm-work/hatch-assignment/graph"
	"github.com/umerm-work/hatch-assignment/graph/generated"
	"github.com/umerm-work/hatch-assignment/infrastructure/postgres"
	"log"
	"net/http"
)

func main() {

	ctx := context.Background()
	c := config.ReadConfig()
	storage := postgres.New(ctx, c)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.BuildResolver(storage)}))
	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", c.DefaultAppPort)
	log.Fatal(http.ListenAndServe(":"+c.DefaultAppPort, router))
}
