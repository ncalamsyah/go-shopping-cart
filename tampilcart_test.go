package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestTampilkanCart(t *testing.T) {
	// Test tampilkan cart tanpa filter
	req, _ := http.NewRequest("GET", "/tampilkanCart", nil)
	w := httptest.NewRecorder()
	tampilkanCart(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected status code. Got %v, expected %v", w.Code, http.StatusOK)
	}

	// Test tampilkan cart dengan filter nama produk
	req, _ = http.NewRequest("GET", "/tampilkanCart", nil)
	req.Form = url.Values{}
	req.Form.Add("namaProduk", "Produk 2")

	w = httptest.NewRecorder()
	tampilkanCart(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected status code. Got %v, expected %v", w.Code, http.StatusOK)
	}

	// Test tampilkan cart dengan filter kuantitas
	req, _ = http.NewRequest("GET", "/tampilkanCart", nil)
	req.Form = url.Values{}
	req.Form.Add("kuantitas", "7")
	w = httptest.NewRecorder()
	tampilkanCart(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected status code. Got %v, expected %v", w.Code, http.StatusOK)
	}
}
