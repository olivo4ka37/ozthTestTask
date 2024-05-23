package main

import (
	graph2 "PostCommentService/Internal/graph"
	"flag"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	// Set true if you want to use machine memory for this or use PostgreSQL
	useMemory := flag.Bool("useMemory", false, "Use in-memory storage")
	flag.Parse()
	resolver := graph2.NewResolver(*useMemory)
	srv := handler.NewDefaultServer(graph2.NewExecutableSchema(graph2.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:8080/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
