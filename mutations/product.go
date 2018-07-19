package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/robertotambunan/go-graphql-sample/data"
)

// GetCreateProductMutation creates a new product and returns it.
func GetCreateProductMutation() *graphql.Field {
	return &graphql.Field{
		Type: data.ProductType,
		Args: graphql.FieldConfigArgument{
			"product_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"product_name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"product_price": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"product_city": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			product := data.Product{
				ProductID:    params.Args["product_id"].(int),
				ProductName:  params.Args["product_name"].(string),
				ProductCity:  params.Args["product_city"].(string),
				ProductPrice: params.Args["product_price"].(string),
			}
			isAdded := data.AddProductData(product)

			if isAdded {
				return product, nil
			}

			return isAdded, nil
		},
	}
}
