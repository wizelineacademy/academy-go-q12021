package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBook(t *testing.T) {
	req, err := http.NewRequest("GET", "/book", nil)
	if err != nil {
		t.Fatal(err)
	}
	expected := `[{"id":"1","title":"The Adventures of Duck and Goose","author":"Sir Quackalot","format":"paperback","price":""},{"id":"2","title":"The Return of Duck and Goose","author":"Sir Quackalot","format":"paperback","price":""},{"id":"3","title":"More Fun with Duck and Goose","author":"Sir Quackalot","format":"paperback","price":""},{"id":"4","title":"Duck and Goose on Holiday","author":"Sir Quackalot","format":"paperback","price":""},{"id":"5","title":"The Return of Duck and Goose","author":"Sir Quackalot","format":"hardback","price":""},{"id":"6","title":"The Adventures of Duck and Goose","author":"Sir Quackalot","format":"hardback","price":""},{"id":"7","title":"My Friend is a Duck","author":"A. Parrot","format":"paperback","price":""},{"id":"8","title":"Annotated Notes on the ‘Duck and Goose’ chronicles","author":"Prof Macaw","format":"ebook","price":""},{"id":"9","title":"‘Duck and Goose’ Cheat Sheet for Students","author":"Polly Parrot","format":"ebook","price":""},{"id":"10","title":"‘Duck and Goose’: an allegory for modern times?","author":"Bor Ing","format":"hardback","price":""}]`
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBook)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetBookById(t *testing.T) {
	req, err := http.NewRequest("GET", "/book/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()
	expected := `{"id":"2","title":"The Return of Duck and Goose","author":"Sir Quackalot","format":"paperback","price":""}`
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBookById)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestConsumeAPI(t *testing.T) {
	req, err := http.NewRequest("GET", "/consume", nil)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"id":"5577006791947779410","title":"Securing DevOps","author":"Julien Vehent","format":"Digital","price":"$39.65"}`
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBook)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}
