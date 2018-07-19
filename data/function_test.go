package data

import (
	"reflect"
	"testing"
)

func TestAddProductData(t *testing.T) {
	type args struct {
		data Product
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1: Data Not Complete - No ProductID",
			args: args{
				data: Product{
					ProductCity:  "Medan",
					ProductName:  "Kursi",
					ProductPrice: "Rp. 540.000",
				},
			},
			want: false,
		},
		{
			name: "Test 2: Data Not Complete - No ProductCity",
			args: args{
				data: Product{
					ProductID:    40,
					ProductName:  "Kursi",
					ProductPrice: "Rp. 540.000",
				},
			},
			want: false,
		},
		{
			name: "Test 3: Data Complete",
			args: args{
				data: Product{
					ProductID:    40,
					ProductName:  "Kursi",
					ProductPrice: "Rp. 540.000",
					ProductCity:  "Medan",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddProductData(tt.args.data); got != tt.want {
				t.Errorf("AddProductData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetProductData(t *testing.T) {
	DataSetProduct = fillDataSet()
	tests := []struct {
		name       string
		wantResult []Product
	}{
		{
			name:       "Test 1",
			wantResult: fillDataSet(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := GetProductData(); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("GetProductData() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
