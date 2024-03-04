package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create an order")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all Orders")
}

func (o *Order) GetbyID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get order by ID")
}

func (o *Order) UpdatebyID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update order by ID")
}


func (o *Order) DeletebyID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update order by ID")
}
