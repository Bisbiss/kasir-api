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

var produk = []Produk{
	{ID: 1, Nama :"Mouse", Harga:50000, Stok:10},
	{ID: 2, Nama :"Mouse Pad", Harga:100000, Stok:10},
	{ID: 3, Nama :"Usb Hub", Harga:200000, Stok:20},
}

// Get localhost:8080/api/produk/{id}
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

func updateProduk(w http.ResponseWriter, r *http.Request){
		idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
		id, err := strconv.Atoi(idStr)
		if err != nil{
			http.Error(w, "Invalid Produk Id", http.StatusBadRequest)
			return
		}

		var updateProduk Produk
		err = json.NewDecoder(r.Body).Decode(&updateProduk)
		if err != nil{
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		for i := range produk {
			if produk[i].ID == id{
				updateProduk.ID = id
				produk[i] = updateProduk
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(updateProduk)
			}
		}
	}

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

func main(){
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