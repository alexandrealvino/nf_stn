package tests

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"nf_stn/adapter"
	databaseMock "nf_stn/database/mock"
	"nf_stn/entities"
	"testing"
)

// TestInsertInvoice tests the /api POST endpoint
func TestInsertInvoice(t *testing.T) {
	c := gomock.NewController(t)
	cm := databaseMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Db: cm,
	}
	cm.EXPECT().InsertInvoice(gomock.Any()).Return(nil)
	var jsonStr = []byte(`{"id":1,"referenceMonth":1,"referenceYear":2020,"document":"00000000000011","description":"insert","amount":10,"isActive":1,"createdAt":"2020-07-20 10:07:37","deactivatedAt":"2020-07-20 10:07:37"}`)

	req, err := http.NewRequest("POST", "/insertInvoice", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.InsertInvoice)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	expected := `{"authenticationStatus":"authorized","requestMethod":"POST","contentType":"application/json","page":"1","totalPages":"1","numberOfInvoices":1,"invoices":[{"id":1,"referenceMonth":1,"referenceYear":2020,"document":"00000000000011","description":"insert","amount":10,"isActive":1,"createdAt":"2020-07-20 10:07:37","deactivatedAt":"2020-07-20 10:07:37"}]}`
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
}
// TestGetInvoices tests the /api GET endpoint
func TestGetInvoices(t *testing.T) {
	c := gomock.NewController(t)
	cm := databaseMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Db: cm,
	}
	//var ex []entities.Invoice
	ex := []entities.Invoice{
		{
			ID: 1,
			ReferenceMonth: 7,
			ReferenceYear: 2022,
			Document: "00000000000014",
			Description: "insert",
			Amount: 10,
			IsActive: 1,
			CreatedAt: "2020-07-29 17:18:04",
			DeactivatedAt: "2020-01-01 00:01:00",
		},
	}
	i := 1
	cm.EXPECT().GetInvoices(gomock.Any()).Return(ex, i, nil)
	req, err := http.NewRequest("GET", "/api/getall", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.GetInvoices)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `{"authenticationStatus":"authorized","requestMethod":"GET","contentType":"application/json","page":"1","totalPages":"1","numberOfInvoices":1,"invoices":[{"id":1,"referenceMonth":7,"referenceYear":2022,"document":"00000000000014","description":"insert","amount":10,"isActive":1,"createdAt":"2020-07-29 17:18:04","deactivatedAt":"2020-01-01 00:01:00"}]}`
	buffer := new(bytes.Buffer)
	err = json.Compact(buffer, rr.Body.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	_ = json.Unmarshal(rr.Body.Bytes(), &buffer)
	if buffer.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
// TestEditInvoice tests the /api PUT endpoint
func TestEditInvoice(t *testing.T) {
	c := gomock.NewController(t)
	cm := databaseMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Db: cm,
	}
	ex := entities.Invoice{
			ID: 1,
			ReferenceMonth: 7,
			ReferenceYear: 2022,
			Document: "00000000000014",
			Description: "insert",
			Amount: 10,
			IsActive: 1,
			CreatedAt: "2020-07-29 17:18:04",
			DeactivatedAt: "2020-01-01 00:01:00",
		}
	var jsonStr = []byte(`{"id":1,"referenceMonth":1,"referenceYear":2020,"document":"00000000000011","description":"update","amount":10,"isActive":1,"createdAt":"2020-07-20 10:07:37","deactivatedAt":"2020-07-20 10:07:37"}`)
	cm.EXPECT().InvoiceExists(gomock.Any()).Return(ex,nil)
	cm.EXPECT().UpdateInvoice(gomock.Any()).Return(nil)
	req, err := http.NewRequest("PUT", "/api/up/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.UpdateInvoice)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"authenticationStatus":"authorized","requestMethod":"PUT","contentType":"application/json","page":"1","totalPages":"1","numberOfInvoices":1,"invoices":[{"id":1,"referenceMonth":1,"referenceYear":2020,"document":"00000000000011","description":"update","amount":0.1,"isActive":1,"createdAt":"2020-07-20 10:07:37","deactivatedAt":"2020-07-20 10:07:37"}]}`
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
}
// TestPatchInvoice tests the /api PATCH endpoint
func TestPatchInvoice(t *testing.T) {
	c := gomock.NewController(t)
	cm := databaseMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Db: cm,
	}
	var jsonStr = []byte(`{"id":1,"referenceMonth":2,"referenceYear":2022,"amount":12,"description":"patched"}`)
	ex := entities.Invoice{
		ID: 1,
		ReferenceMonth: 2,
		ReferenceYear: 2022,
		Document: "00000000000011",
		Description: "patched",
		Amount: 12,
		IsActive: 1,
		CreatedAt: "2020-07-20 10:07:37",
		DeactivatedAt: "2020-01-01 00:01:00",
	}
	cm.EXPECT().InvoiceExists(gomock.Any()).Return(ex,nil)
	cm.EXPECT().GetInvoiceByID(gomock.Any()).Return(ex,nil)
	cm.EXPECT().PatchInvoice(gomock.Any()).Return(nil)
	req, err := http.NewRequest("PUT", "/api/patch/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.PatchInvoice)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"authenticationStatus":"authorized","requestMethod":"PUT","contentType":"application/json","page":"1","totalPages":"1","numberOfInvoices":1,"invoices":[{"id":1,"referenceMonth":2,"referenceYear":2022,"document":"00000000000011","description":"patched","amount":12,"isActive":1,"createdAt":"2020-07-20 10:07:37","deactivatedAt":"2020-01-01 00:01:00"}]}`
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
}
// TestDeleteInvoice tests the /api DELETE endpoint
func TestDeleteInvoice(t *testing.T) {
	c := gomock.NewController(t)
	cm := databaseMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Db: cm,
	}
	var jsonStr = []byte(`{"id":1}`)
	ex := entities.Invoice{
		ID: 1,
		ReferenceMonth: 2,
		ReferenceYear: 2022,
		Document: "00000000000011",
		Description: "patched",
		Amount: 12,
		IsActive: 1,
		CreatedAt: "2020-07-20 10:07:37",
		DeactivatedAt: "2020-01-01 00:01:00",
	}
	cm.EXPECT().GetInvoiceByID(gomock.Any()).Return(ex,nil)
	cm.EXPECT().DeleteInvoice(gomock.Any()).Return(nil)
	req, err := http.NewRequest("DELETE", "/api/del/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	handler := http.HandlerFunc(routes.DeleteInvoice)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"authenticationStatus":"authorized","requestMethod":"DELETE","contentType":"application/json","page":"1","totalPages":"1","numberOfInvoices":1,"invoices":[{"id":1,"referenceMonth":2,"referenceYear":2022,"document":"00000000000011","description":"patched","amount":12,"isActive":0,"createdAt":"2020-07-20 10:07:37","deactivatedAt":"2020-01-01 00:01:00"}]}`
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
}
// TestInit tests the Init function
func TestInit(t *testing.T) {
	c := gomock.NewController(t)
	cm := databaseMock.NewMockDbClient(c)
	cm.EXPECT().Init().Return()
}
// TestGetInvoiceByDocument tests the GetInvoiceByDocument function
func TestGetInvoiceByDocument(t *testing.T) {
	c := gomock.NewController(t)
	cm := databaseMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Db: cm,
	}
	ex := entities.Invoice{
			ID: 1,
			ReferenceMonth: 7,
			ReferenceYear: 2022,
			Document: "00000000000014",
			Description: "insert",
			Amount: 10,
			IsActive: 1,
			CreatedAt: "2020-07-29 17:18:04",
			DeactivatedAt: "2020-01-01 00:01:00",
		}
	cm.EXPECT().GetInvoiceByDocument(gomock.Any()).Return(ex,nil)
	rr, _ := routes.Db.GetInvoiceByDocument("00000000000014")
	expected := entities.Invoice{
		ID: 1,
		ReferenceMonth: 7,
		ReferenceYear: 2022,
		Document: "00000000000014",
		Description: "insert",
		Amount: 10,
		IsActive: 1,
		CreatedAt: "2020-07-29 17:18:04",
		DeactivatedAt: "2020-01-01 00:01:00",
	}
	if rr != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr, expected)
	}
}
// TestGetInvoiceByID tests the GetInvoiceByID function
func TestGetInvoiceByID(t *testing.T) {
	c := gomock.NewController(t)
	cm := databaseMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Db: cm,
	}
	ex := entities.Invoice{
		ID: 1,
		ReferenceMonth: 7,
		ReferenceYear: 2022,
		Document: "00000000000014",
		Description: "insert",
		Amount: 10,
		IsActive: 1,
		CreatedAt: "2020-07-29 17:18:04",
		DeactivatedAt: "2020-01-01 00:01:00",
	}
	cm.EXPECT().GetInvoiceByID(gomock.Any()).Return(ex,nil)
	rr, _ := routes.Db.GetInvoiceByID(1)
	expected := entities.Invoice{
		ID: 1,
		ReferenceMonth: 7,
		ReferenceYear: 2022,
		Document: "00000000000014",
		Description: "insert",
		Amount: 10,
		IsActive: 1,
		CreatedAt: "2020-07-29 17:18:04",
		DeactivatedAt: "2020-01-01 00:01:00",
	}
	if rr != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr, expected)
	}
}
// TestGetUser tests the GetUser function
func TestGetUser(t *testing.T) {
	c := gomock.NewController(t)
	cm := databaseMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Db: cm,
	}
	ex := entities.User{
		ID: 1,
		Username: "username",
		Hash: "$2a$04$/GvrVH49FLVOVqbtXd99oul2Ma8Nw84dHbYqapq93R042Q98OpEAW",
	}
	cm.EXPECT().GetUser(gomock.Any()).Return(ex.ID,ex.Username,ex.Hash,nil)
	var u entities.User
	u.ID, u.Username, u.Hash, _ = routes.Db.GetUser(ex.Username)
	expected := entities.User{
		ID: 1,
		Username: "username",
		Hash: "$2a$04$/GvrVH49FLVOVqbtXd99oul2Ma8Nw84dHbYqapq93R042Q98OpEAW",
	}
	if u != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			u, expected)
	}
}
// TestInvoiceExists tests the InvoiceExists function
func TestInvoiceExists(t *testing.T) {
	c := gomock.NewController(t)
	cm := databaseMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Db: cm,
	}
	ex := entities.Invoice{
		ID: 1,
		ReferenceMonth: 7,
		ReferenceYear: 2022,
		Document: "00000000000014",
		Description: "insert",
		Amount: 10,
		IsActive: 1,
		CreatedAt: "2020-07-29 17:18:04",
		DeactivatedAt: "2020-01-01 00:01:00",
	}
	cm.EXPECT().InvoiceExists(gomock.Any()).Return(ex,nil)
	rr, _ := routes.Db.InvoiceExists("00000000000014")
	expected := entities.Invoice{
		ID: 1,
		ReferenceMonth: 7,
		ReferenceYear: 2022,
		Document: "00000000000014",
		Description: "insert",
		Amount: 10,
		IsActive: 1,
		CreatedAt: "2020-07-29 17:18:04",
		DeactivatedAt: "2020-01-01 00:01:00",
	}
	if rr != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr, expected)
	}
}
//