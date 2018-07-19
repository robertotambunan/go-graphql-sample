package data

var DataSetProduct []Product

func Initialize() {
	DataSetProduct = fillDataSet()
}

func fillDataSet() []Product {
	tempData := []Product{
		Product{
			ProductID:    10,
			ProductName:  "Tapai",
			ProductPrice: "Rp. 10.000,00",
			ProductCity:  "Medan",
		},
		Product{
			ProductID:    11,
			ProductName:  "Bika Ambon",
			ProductPrice: "Rp. 70.000,00",
			ProductCity:  "Medan",
		},
		Product{
			ProductID:    12,
			ProductName:  "Ayam Gepuk Pak Gembus (Refill Nasi)",
			ProductPrice: "Rp. 21.000,00",
			ProductCity:  "Medan",
		},
		Product{
			ProductID:    12,
			ProductName:  "Televisi",
			ProductPrice: "Rp. 2.121.000,00",
			ProductCity:  "Medan",
		},
	}
	return tempData
}
