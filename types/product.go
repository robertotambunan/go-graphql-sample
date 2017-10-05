package types

import (
	"github.com/graphql-go/graphql"
)

// Product type definition.
type Product struct {
	ProductID    int    `db:"product_id" json:"product_id"`
	ProductName  string `db:"product_name" json:"product_name"`
	ProductPrice string `db:"product_price" json:"product_price"`
	ProductCity  string `db:"product_city" json:"product_city"`
}

// ProductType is the GraphQL schema for the user type.
var ProductType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
		"product_id":    &graphql.Field{Type: graphql.Int},
		"product_name":  &graphql.Field{Type: graphql.String},
		"product_price": &graphql.Field{Type: graphql.String},
		"product_city":  &graphql.Field{Type: graphql.String},
	},
})
