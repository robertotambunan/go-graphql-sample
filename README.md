First Step,
Create your MySQL Table, with this structure


CREATE DATABASE `belajar` /*!40100 DEFAULT CHARACTER SET utf8 */;

CREATE TABLE `product` (
  `product_id` int(11) NOT NULL AUTO_INCREMENT,
  `product_name` varchar(200) NOT NULL,
  `product_price` varchar(20) NOT NULL,
  `product_city` varchar(45) NOT NULL,
  PRIMARY KEY (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;



Note : This project will handle how showing data and add data (Query and Mutation) in graphQL + Go ;)

How to make it run?
1. Of course go get first, then go build && ./{{project name}}

2. localhost:9090 will be listening

3. Open your Postman (If you have other, then open it)

4. For Showing Data (Query), Use Post Method, open localhost:9090, set your Content-Type -> application/json, then send this json as body

    {"query": "query { products { product_id, product_name, product_city, product_price } }"}

    It will show you the atribute listed above from products

5. For Adding Data (on of Mutation's function),  Use Post Method, open localhost:9090, set your Content-Type -> application/json, then send this json as body

    {"query": "mutation { createProduct(product_name: \"Baju\", product_price: \"500000\", product_city: \"Jakarta\") { product_name,product_price, product_city } }"}
    
    Fill product_name, product_price, product_city with value that you want, of course in string
    Then data will be added to your db