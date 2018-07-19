[![Build Status](https://travis-ci.org/robertotambunan/go-graphql-sample.svg?branch=master)](https://travis-ci.org/robertotambunan/go-graphql-sample)
# Go-GraphQL-Sample
This is a very simple project about how to implement graphQL in your golang project. The flow is so simple and very recommended for just quick learning.
Dataset is just a simply array of struct.

## Getting Started
Simply follow steps below to run this project

## How To Install and Run
```
1. Clone the repository to your local machine (https://github.com/robertotambunan/go-graphql-sample.git)
2. go get -v
3. go build && ./go-graphql-sample
Now your project is listening
```

## How To Access Your GraphQL
The project is running in localhost:9090/gql
### Query
Hit localhost:9090/gql with method POST with body
```
    {"query": "query { products { product_id, product_name, product_city, product_price } }"}
```
It will show you the data of your dataset
# Mutation
Hit localhost:9090/gql with method POST with body
```
    {"query": "mutation { createProduct(product_id: 17, product_name: \"Baju Boxer\", product_price: \"Rp.500.000\", product_city: \"Jakarta\") { product_name,product_price, product_city } }"}
```
It will add new data to your dataset, you can check the data using Query