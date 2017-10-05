package mutations

import (
	"github.com/graphql-go/graphql"
	data "github.com/tokopedia/project2/db"
	"github.com/tokopedia/project2/types"
)

// GetCreateUserMutation creates a new user and returns it.
func GetCreateProductMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.ProductType,
		Args: graphql.FieldConfigArgument{
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
			product := types.Product{
				ProductName:  params.Args["product_name"].(string),
				ProductCity:  params.Args["product_city"].(string),
				ProductPrice: params.Args["product_price"].(string),
			}
			isAdded := data.AddProductData(product.ProductName, product.ProductCity, product.ProductPrice)

			if isAdded {
				return product, nil
			}

			return isAdded, nil
		},
	}
}
