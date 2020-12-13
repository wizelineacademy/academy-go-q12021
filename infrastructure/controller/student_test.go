package controller

//
//type Controller interface {
//	ReadStudentsHandler(w http.ResponseWriter, r *http.Request)
//	StoreStudentURLHandler(w http.ResponseWriter, r *http.Request)
//}
//
////
//func TestReadStudentsHandler(t *testing.T) {
//	req, err := http.NewRequest("GET", "/readcsv", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc("/readcsv", Controller.ReadStudentsHandler)
//	handler.ServeHTTP(rr, req)
//
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusOK)
//	}
//
//	// Check the response body is what we expect.
//	expected := `[{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb2405@gmail.com","phone_number":"0987654321"},{"id":2,"first_name":"Rick","last_name":"Parker","email_address":"rick.parker@gmail.com","phone_number":"1234567890"},{"id":3,"first_name":"Kelly","last_name":"Franco","email_address":"kelly.franco@gmail.com","phone_number":"1112223333"}]`
//	if rr.Body.String() != expected {
//		t.Errorf("handler returned unexpected body: got %v want %v",
//			rr.Body.String(), expected)
//	}
//	w.status
//	r
//	if w.StatusCode != http.StatusOK {
//		t.Errorf("expected %d, got: %d", http.StatusOK, res.StatusCode)
//	}
//}
