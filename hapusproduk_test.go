package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHapusProduk(t *testing.T) {
	// Test hapus produk yang ada
	req, _ := http.NewRequest("POST", "/hapusProduk", nil)
	req.Form = url.Values{}
	req.Form.Add("kodeProduk", "P001")
	w := httptest.NewRecorder()
	hapusProduk(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected status code. Got %v, expected %v", w.Code, http.StatusOK)
	}

	if _, ok := cart["P001"]; ok {
		t.Errorf("Produk dengan kode P001 masih ada di cart")
	}

	// Test hapus produk yang tidak ada
	req, _ = http.NewRequest("POST", "/hapusProduk", nil)
	req.Form = url.Values{}
	req.Form.Add("kodeProduk", "P002")

	w = httptest.NewRecorder()
	hapusProduk(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected status code. Got %v, expected %v", w.Code, http.StatusOK)
	}
}
