package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/raamas/ShopAPI/controllers"
	"github.com/raamas/ShopAPI/db"
	"github.com/raamas/ShopAPI/models"
)

func main() {
	db.Connect()
	db.DB.AutoMigrate(&models.Product{})

	methods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization"})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Print("DB Is Ready")
	r := mux.NewRouter()

	r.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/products/{product_id}", controllers.GetProduct).Methods("GET")
	r.HandleFunc("/products/category/{product_category}", controllers.GetProductsByCategory).Methods("GET")

	s := r.PathPrefix("/admin").Subrouter()

	s.HandleFunc("/products/add", controllers.AddProduct).Methods("POST")
	s.HandleFunc("/products/{product_id}/delete", controllers.DeleteProduct).Methods("DELETE")
	s.HandleFunc("/products/{product_id}/update", controllers.PutProduct).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(r)))
}
