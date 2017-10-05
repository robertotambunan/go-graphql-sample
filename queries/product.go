package queries

import (
	"github.com/graphql-go/graphql"
	data "github.com/tokopedia/project2/db"
	"github.com/tokopedia/project2/types"
)

// GetUserQuery returns the queries available against user type.
func GetUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.ProductType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			products := data.GetProductData()
			return products, nil
		},
	}
}
