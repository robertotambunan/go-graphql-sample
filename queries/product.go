package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/tokopedia/go-graphql-sample/data"
)

// GetProductQuery returns the queries available against product type.
func GetProductQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(data.ProductType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			products := data.GetProductData()
			return products, nil
		},
	}
}
