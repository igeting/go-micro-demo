package model

import "strconv"

type prodModel struct {
	ProdID   int
	ProdName string
}

func NewProd(id int, name string) *prodModel {
	return &prodModel{ProdID: id, ProdName: name}
}

func NewProdList(n int) []*prodModel {
	result := make([]*prodModel, 0)
	for i := 0; i < n; i++ {
		result = append(result, NewProd(i+10000, "prodName"+strconv.Itoa(i+10000)))
	}
	return result
}
