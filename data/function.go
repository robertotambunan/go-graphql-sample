package data

import "log"

func GetProductData() (result []Product) {
	if len(DataSetProduct) > 0 {
		result = DataSetProduct
	}
	return result

}

func AddProductData(data Product) bool {
	if data.ProductName != "" && data.ProductID > 0 && data.ProductCity != "" && data.ProductPrice != "" {
		DataSetProduct = append(DataSetProduct, data)
		return true
	}
	log.Println("Data is not complete")
	return false
}
