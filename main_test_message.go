/*
// main.go

package main

import (
	"fmt"
	"log"

	"github.com/favtuts/go-grpc-gateway-example/protogen/golang/orders"
	"github.com/favtuts/go-grpc-gateway-example/protogen/golang/product"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	orderItem := orders.Order{
		OrderId:    10,
		CustomerId: 11,
		IsActive:   true,
		OrderDate:  &date.Date{Year: 2024, Month: 9, Day: 19},
		Products: []*product.Product{
			{ProductId: 1, ProductName: "CocaCola", ProductType: product.ProductType_DRINK},
		},
	}

	bytes, err := protojson.Marshal(&orderItem)
	if err != nil {
		log.Fatal("deserialization error:", err)
	}

	fmt.Println(string(bytes))
}
*/