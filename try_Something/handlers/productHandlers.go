package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	. "test/models"
	"time"

	"github.com/gorilla/mux"
)

var productStore = make(map[string]Product)
var id int = 0

// Hata Fonksiyonu
func CheckError(err error) {
	fmt.Printf("Hata var hataaaa : %v\n", err)
}

// HTTP Post - /api/products
func PostProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	CheckError(err)

	product.CreateOn = time.Now()
	id++
	product.ID = id
	key := strconv.Itoa(id)
	productStore[key] = product

	data, err := json.Marshal(product)
	CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)

}

// HTTP Get - /api/products
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []Product
	for _, product := range productStore {
		products = append(products, product)
	}

	data, err := json.Marshal(products)
	CheckError(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// HTTP Get - /api/products/{id}
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	for _, prd := range productStore {
		if prd.ID == key {
			product = prd
		}
	}

	data, err := json.Marshal(product)
	CheckError(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// HTTP Put - /api/products/{id} --->UPDATE için
func PutProductHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	key := vars["id"]

	var prodUp Product
	err = json.NewDecoder(r.Body).Decode(&prodUp)
	CheckError(err)

	if _, ok := productStore[key]; ok {
		prodUp.ID, _ = strconv.Atoi(key)
		prodUp.ChangeOn = time.Now()
		delete(productStore, key)
		productStore[key] = prodUp
	} else {
		log.Printf("Değer bulunamadı : %s", key)
	}
	w.WriteHeader(http.StatusOK)
}

// HTTP Delete - /api/products/{id}
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	if _, ok := productStore[key]; ok {
		delete(productStore, key)
	} else {
		log.Printf("Değer bulunamadı : %s", key)
	}
	w.WriteHeader(http.StatusOK)
}
