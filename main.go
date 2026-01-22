package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Produk struct {
	ID		int		`json:"id"`
	Nama	string	`json:"nama"`
	Harga	int		`json:"harga":`
	Stok	int		`json:"stok"`
}

type Category struct {
	ID			int 	`json:"id"`
	Name		string	`json:"name"`
	Description	string	`json:"description"`
}

var produk = []Produk{
	{ID: 1, Nama :"Mouse", Harga:50000, Stok:10},
	{ID: 2, Nama :"Mouse Pad", Harga:100000, Stok:10},
	{ID: 3, Nama :"Usb Hub", Harga:200000, Stok:20},
}

var category = []Category{
	{ID: 1, Name:"Aksesoris", Description:"Aksesoris Laptop/PC"},
	{ID: 2, Name:"Monitor", Description:"Daftar Monitor Eksternal"},
	{ID: 3, Name:"CPU", Description:"Daftar CPU Eksternal"},
}

// function Get Produk By ID
func getProdukById (w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w, "Invalid Produk Id", http.StatusBadRequest)
		return
	}
	
	for _, p := range produk{
		if p.ID == id{
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	
	http.Error(w, "Produk Tidak Ada", http.StatusNotFound)
}


// function Update Produk By ID
func updateProduk(w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w, "Invalid Produk Id", http.StatusBadRequest)
		return
	}

	var produkBaru Produk
	err = json.NewDecoder(r.Body).Decode(&produkBaru)
	if err != nil{
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i := range produk {
		if produk[i].ID == id{
			produkBaru.ID = id
			produk[i] = produkBaru
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produkBaru)
		}
	}
}

// function Delete Produk By ID
func deleteProduk(w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w, "Invalid Produk Id", http.StatusBadRequest)
		return
	}

	for i, p := range produk{
		if p.ID == id {
			produk = append(produk[:i], produk[1+i:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message" : "Sukses Delete",
			})
			return
		}
	}
	http.Error(w, "Produk Tidak Ada", http.StatusNotFound)

}

// Function Get Category By ID
func getCategoryById (w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w, "Invalid Category Id", http.StatusBadRequest)
		return
	}

	for _, c := range category{
		if c.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	http.Error(w, "Categori Tidak Ada", http.StatusNotFound)

}

// Function Update Category By ID
func updateCategory(w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w, "Invalid Category Id", http.StatusBadRequest)
		return
	}

	var categoryBaru Category
	err = json.NewDecoder(r.Body).Decode(&categoryBaru)
	if err != nil{
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i := range category{
		if category[i].ID == id {
			categoryBaru.ID = id
			category[i] = categoryBaru
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categoryBaru)
		}
	}
}

// Function Delete Category By ID
func deleteCategory(w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w, "Invalid Category Id", http.StatusBadRequest)
		return
	}

	for i, c := range category{
		if c.ID == id {
			category = append(category[:i], category[1+i:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message" : "Sukses Delete Category",
			})
			return
		}
	}

	http.Error(w, "Category Tidak Ada", http.StatusNotFound)
}

func main(){
	// 1. Produk
	// Get localhost:8080/api/produk/{id}
	// Put localhost:8080/api/produk/{id}
	// Delete localhost:8080/api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request){
		if r.Method == "GET" {
			getProdukById(w, r)
		}else if r.Method == "PUT"{
			updateProduk(w, r)
		}else if r.Method == "DELETE"{
			deleteProduk(w,r)
		}
	})

	
	// Get localhost:8080/api/produk
	// Post localhost:8080/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request){
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		}else if r.Method == "POST" {
			var produkBaru Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil{
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}
			
			produkBaru.ID = len(produk)+1
			produk = append(produk, produkBaru)
				
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) //201
			json.NewEncoder(w).Encode(produkBaru)
		}
	})

	// 2. Categories
	// Get localhost:8080/categories 
	// Post localhost:8080/categories
	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request){
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(category)
		}else if r.Method == "POST" {
			var categoryBaru Category
			err := json.NewDecoder(r.Body).Decode(&categoryBaru)
			if err != nil {
				http.Error(w, "Invalid Request", http.StatusBadRequest)
				return
			}

			categoryBaru.ID = len(category)+1
			category = append(category, categoryBaru)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(categoryBaru)
		}
	})

	// Get localhost:8080/categories/{id}
	// PUT localhost:8080/categories/{id}
	// DELETE localhost:8080/categories/{id}
	http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getCategoryById(w,r)
		}else if r.Method == "PUT" {
			updateCategory(w,r)
		}else if r.Method == "DELETE" {
			deleteCategory(w,r)
		}
	})
	

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status" : "oke",
			"message" : "Api running",
		})
	})
	
	fmt.Println("Server running di localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	
	if err != nil{
		fmt.Println("gagal terhubung ke server")
	}
}