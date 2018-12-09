package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"github.com/traggo/server/tag"
	"github.com/jinzhu/gorm"
)

// Handler creates a graphql handler.
func Handler(db *gorm.DB) *handler.Handler {
	queryFields := merge(tag.Queries(db))
	mutationFields := merge(tag.Mutations(db))

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: queryFields}
	rootMutations := graphql.ObjectConfig{Name: "Mutations", Fields: mutationFields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: graphql.NewObject(rootMutations)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	return handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     false,
		GraphiQL:   true,
		Playground: true,
	})
}

func merge(toMerge ...graphql.Fields) graphql.Fields {
	var result = graphql.Fields{}
	for _, subset := range toMerge {
		for key, value := range subset {
			result[key] = value
		}
	}
	return result
}