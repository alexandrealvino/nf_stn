package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"nf_stn/adapter"
	"nf_stn/database"
	"testing"
)
//
func TestInsertInvoice(t *testing.T) {
	database.ClearTable()
	var jsonStr = []byte(`{"id":1,"referenceMonth":1,"referenceYear":2020,"document":"00000000000011","description":"insert","amount":10,"isActive":1,"createdAt":"2020-07-20 10:07:37","deactivatedAt":"2020-07-20 10:07:37"}`)

	req, err := http.NewRequest("POST", "/insertInvoice", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(adapter.InsertInvoice)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	expected := `{"id":1,"referenceMonth":1,"referenceYear":2020,"document":"00000000000011","description":"insert","amount":10,"isActive":1,"createdAt":"2020-07-20 10:07:37","deactivatedAt":"2020-07-20 10:07:37"}`
	buffer := new(bytes.Buffer)
	err = json.Compact(buffer, rr.Body.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	_ = json.Unmarshal(rr.Body.Bytes(), buffer)
	if buffer.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
} // tests the /api/insertInvoice endpoint
//
func TestGetAll(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/getall", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(adapter.GetAll)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `[{"id":1,"referenceMonth":1,"referenceYear":2020,"document":"00000000000011","description":"insert","amount":10,"isActive":1,"createdAt":"2020-07-20 10:07:37","deactivatedAt":"2020-07-20 10:07:37"}]`
	buffer := new(bytes.Buffer)
	err = json.Compact(buffer, rr.Body.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	_ = json.Unmarshal(rr.Body.Bytes(), buffer)
	if buffer.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
} // tests the /api/getAll endpoint
//
func TestEditInvoice(t *testing.T) {

	var jsonStr = []byte(`{"id":1,"referenceMonth":1,"referenceYear":2020,"document":"00000000000011","description":"update","amount":10,"isActive":1,"createdAt":"2020-07-20 10:07:37","deactivatedAt":"2020-07-20 10:07:37"}`)

	req, err := http.NewRequest("PUT", "/api/up/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(adapter.UpdateInvoice)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":1,"referenceMonth":1,"referenceYear":2020,"document":"00000000000011","description":"update","amount":10,"isActive":1,"createdAt":"2020-07-20 10:07:37","deactivatedAt":"2020-07-20 10:07:37"}`
	buffer := new(bytes.Buffer)
	err = json.Compact(buffer, rr.Body.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	_ = json.Unmarshal(rr.Body.Bytes(), buffer)
	if buffer.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
} // tests the /api/up/ endpoint
//
func TestPatchInvoice(t *testing.T) {

	var jsonStr = []byte(`{"id":1,"referenceMonth":2,"referenceYear":2022,"amount":12,"description":"patched"}`)

	req, err := http.NewRequest("PUT", "/api/patch/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(adapter.PatchInvoice)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":1,"referenceMonth":2,"referenceYear":2022,"document":"00000000000011","description":"patched","amount":12,"isActive":1,"createdAt":"2020-07-20 10:07:37","deactivatedAt":"2020-07-20 10:07:37"}`
	buffer := new(bytes.Buffer)
	err = json.Compact(buffer, rr.Body.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	_ = json.Unmarshal(rr.Body.Bytes(), buffer)
	if buffer.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
} // tests the /api/patch/ endpoint
//
func TestDeleteInvoice(t *testing.T) {
	var jsonStr = []byte(`{"id":1}`)

	req, err := http.NewRequest("DELETE", "/api/del/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	handler := http.HandlerFunc(adapter.DeleteInvoice)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
	expected := "1"
	buffer := new(bytes.Buffer)
	err = json.Compact(buffer, rr.Body.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	_ = json.Unmarshal(rr.Body.Bytes(), buffer)
	if buffer.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
} // tests the /api/del/ endpoint
//

