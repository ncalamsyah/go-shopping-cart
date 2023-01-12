package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Product struct {
	kodeProduk string
	namaProduk string
	kuantitas  int
}

var cart = make(map[string]Product)

func tambahProduk(w http.ResponseWriter, r *http.Request) {
	kodeProduk := r.FormValue("kodeProduk")
	namaProduk := r.FormValue("namaProduk")
	kuantitas, _ := strconv.Atoi(r.FormValue("kuantitas"))

	product := cart[kodeProduk]
	if _, ok := cart[kodeProduk]; ok {
		// menambahkan kuantitas produk yang sudah ada
		product.kuantitas += kuantitas
		cart[kodeProduk] = product
	} else {
		// menambahkan produk baru ke dalam cart
		cart[kodeProduk] = Product{kodeProduk, namaProduk, kuantitas}
	}
	fmt.Fprintf(w, "Produk dengan kode %s telah ditambahkan dengan kuantitas %d", kodeProduk, kuantitas)
}

func hapusProduk(w http.ResponseWriter, r *http.Request) {
	kodeProduk := r.FormValue("kodeProduk")
	if _, ok := cart[kodeProduk]; ok {
		delete(cart, kodeProduk)
		fmt.Fprintf(w, "Produk dengan kode %s telah di hapus dari cart", kodeProduk)
	} else {
		fmt.Fprintf(w, "Produk dengan kode %s tidak ditemukan di dalam cart", kodeProduk)
	}
}

func tampilkanCart(w http.ResponseWriter, r *http.Request) {
	namaProdukFilter := r.FormValue("namaProduk")
	kuantitasFilter, _ := strconv.Atoi(r.FormValue("kuantitas"))
	fmt.Fprintf(w, "Isi Cart: \n")
	for _, product := range cart {
		if (namaProdukFilter == "" || product.namaProduk == namaProdukFilter) &&
			(kuantitasFilter == 0 || product.kuantitas == kuantitasFilter) {
			fmt.Fprintf(w, "%s - %s - (%d)\n", product.kodeProduk, product.namaProduk, product.kuantitas)
		}
	}
}

func main() {
	http.HandleFunc("/tambahProduk", tambahProduk)
	http.HandleFunc("/hapusProduk", hapusProduk)
	http.HandleFunc("/tampilkanCart", tampilkanCart)

	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
