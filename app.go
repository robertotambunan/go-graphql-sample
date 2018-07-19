package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/tokopedia/go-graphql-sample/data"
	"github.com/tokopedia/go-graphql-sample/mutations"
	"github.com/tokopedia/go-graphql-sample/queries"
)

func init() {
	data.Initialize()
}

func main() {

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: queries.GetRootFields(),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootMutation",
			Fields: mutations.GetRootFields(),
		}),
	}

	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("Failed to create new schema, error: %v", err)
	}

	httpHandler := handler.New(&handler.Config{
		Schema: &schema,
	})

	http.Handle("/gql", httpHandler)
	log.Print("ready: listening...\n")
	http.ListenAndServe(":9090", nil)

}
