package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestTambahProduk(t *testing.T) {
	// Test tambah produk baru
	req, _ := http.NewRequest("POST", "/tambahProduk", nil)
	req.Form = url.Values{}
	req.Form.Add("kodeProduk", "P001")
	req.Form.Add("namaProduk", "Produk 1")
	req.Form.Add("kuantitas", "5")

	w := httptest.NewRecorder()
	tambahProduk(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected status code. Got %v, expected %v", w.Code, http.StatusOK)
	}

	if _, ok := cart["P001"]; !ok {
		t.Errorf("Produk dengan kode P001 tidak ditemukan di cart")
	}

	// Test tambah kuantitas produk yang sudah ada
	req, _ = http.NewRequest("POST", "/tambahProduk", nil)
	req.Form = url.Values{}
	req.Form.Add("kodeProduk", "P001")
	req.Form.Add("kuantitas", "3")
	w = httptest.NewRecorder()
	tambahProduk(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected status code. Got %v, expected %v", w.Code, http.StatusOK)
	}

	if cart["P001"].kuantitas != 8 {
		t.Errorf("Kuantitas produk P001 salah. Got %d, expected 8", cart["P001"].kuantitas)
	}
}
