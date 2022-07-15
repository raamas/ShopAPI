package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/raamas/ShopAPI/db"
	"github.com/raamas/ShopAPI/models"
)

var Products []models.Product

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db.DB.Find(&Products)
	json.NewEncoder(w).Encode(&Products)
}

func GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	db.DB.Find(&Products, "product_category = ?", params["product_category"])
	json.NewEncoder(w).Encode(&Products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var Product models.Product

	db.DB.First(&Product, "product_id = ?", params["product_id"])

	if Product.ProductID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Product Not Found"))
		return
	}

	json.NewEncoder(w).Encode(&Product)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	w.Header().Set("Content-Type", "application/json")

	var Product models.Product

	_ = json.NewDecoder(r.Body).Decode(&Product)
	Product.ProductID = uint(rand.Intn(25)) // Mock ID - not safe
	createdProduct := db.DB.Create(&Product)
	err := createdProduct.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&Product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var Product models.Product

	db.DB.First(&Product, "product_id = ?", params["product_id"])

	if Product.ProductID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Product Not Found"))
		return
	}

	db.DB.Unscoped().Delete(&Product)
	db.DB.Find(&Product)
	json.NewEncoder(w).Encode(&Product)
}

func PutProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var Product models.Product

	db.DB.First(&Product, "product_id = ?", params["product_id"])
	_ = json.NewDecoder(r.Body).Decode(&Product)
	db.DB.Save(&Product)

	if Product.ProductID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Product Not Found"))
	}

	json.NewEncoder(w).Encode(&Product)
}
