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

//// NewMockDbClient mocks
//func NewMockDbClient(ctrl *gomock.Controller) *MockDataBase {
//	mock := &MockDataBase{ctrl: ctrl}
//	mock.recorder = &MockDataBaseMockRecorder{mock}
//	return mock
//}

// TestInsertInvoice tests the /api/insertInvoice endpoint
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
}
// TestGetAll tests the /api/getAll endpoint
func TestGetAll(t *testing.T) {
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
	cm.EXPECT().GetAll().Return(ex, nil)
	req, err := http.NewRequest("GET", "/api/getall", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.GetAll)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `[{"id":1,"referenceMonth":7,"referenceYear":2022,"document":"00000000000014","description":"insert","amount":10,"isActive":1,"createdAt":"2020-07-29 17:18:04","deactivatedAt":"2020-01-01 00:01:00"}]`
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
// TestEditInvoice tests the /api/up/ endpoint
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
}
// TestPatchInvoice tests the /api/patch/ endpoint
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
	expected := `{"id":1,"referenceMonth":2,"referenceYear":2022,"document":"00000000000011","description":"patched","amount":12,"isActive":1,"createdAt":"2020-07-20 10:07:37","deactivatedAt":"2020-01-01 00:01:00"}`
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
// TestDeleteInvoice tests the /api/del/ endpoint
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
	expected := `{"ID":1}`
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
// TestPagination tests the /api/pagination endpoint
func TestPagination(t *testing.T) {
	c := gomock.NewController(t)
	cm := databaseMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Db: cm,
	}
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
	cm.EXPECT().Pagination(gomock.Any()).Return(ex, nil)
	req, err := http.NewRequest("GET", "/api/pagination", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.Pagination)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `[{"id":1,"referenceMonth":7,"referenceYear":2022,"document":"00000000000014","description":"insert","amount":10,"isActive":1,"createdAt":"2020-07-29 17:18:04","deactivatedAt":"2020-01-01 00:01:00"}]`
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
// TestPaginationOrderByMonth tests the /api/pagination/{offset}/month endpoint
func TestPaginationOrderByMonth(t *testing.T) {
	c := gomock.NewController(t)
	cm := databaseMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Db: cm,
	}
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
	cm.EXPECT().PaginationOrderByMonth(gomock.Any()).Return(ex, nil)
	req, err := http.NewRequest("GET", "/api/pagination/{offset}/month", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.PaginationOrderByMonth)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `[{"id":1,"referenceMonth":7,"referenceYear":2022,"document":"00000000000014","description":"insert","amount":10,"isActive":1,"createdAt":"2020-07-29 17:18:04","deactivatedAt":"2020-01-01 00:01:00"}]`
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
// TestPaginationOrderByYear tests the /api/pagination/{offset}/year endpoint
func TestPaginationOrderByYear(t *testing.T) {
	c := gomock.NewController(t)
	cm := databaseMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Db: cm,
	}
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
	cm.EXPECT().PaginationOrderByYear(gomock.Any()).Return(ex, nil)
	req, err := http.NewRequest("GET", "/api/pagination/{offset}/year", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.PaginationOrderByYear)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `[{"id":1,"referenceMonth":7,"referenceYear":2022,"document":"00000000000014","description":"insert","amount":10,"isActive":1,"createdAt":"2020-07-29 17:18:04","deactivatedAt":"2020-01-01 00:01:00"}]`
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
// TestPaginationOrderByDocument tests the /api/pagination/{offset}/document/ endpoint
func TestPaginationOrderByDocument(t *testing.T) {
	c := gomock.NewController(t)
	cm := databaseMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Db: cm,
	}
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
	cm.EXPECT().PaginationOrderByDocument(gomock.Any()).Return(ex, nil)
	req, err := http.NewRequest("GET", "/api/pagination/{offset}/document", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.PaginationOrderByDocument)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `[{"id":1,"referenceMonth":7,"referenceYear":2022,"document":"00000000000014","description":"insert","amount":10,"isActive":1,"createdAt":"2020-07-29 17:18:04","deactivatedAt":"2020-01-01 00:01:00"}]`
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
// TestPaginationOrderByMonthYear tests the /api/pagination/{offset}/month/year/ endpoint
func TestPaginationOrderByMonthYear(t *testing.T) {
	c := gomock.NewController(t)
	cm := databaseMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Db: cm,
	}
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
	cm.EXPECT().PaginationOrderByMonthYear(gomock.Any()).Return(ex, nil)
	req, err := http.NewRequest("GET", "/api/pagination/{offset}/month/year/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.PaginationOrderByMonthYear)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `[{"id":1,"referenceMonth":7,"referenceYear":2022,"document":"00000000000014","description":"insert","amount":10,"isActive":1,"createdAt":"2020-07-29 17:18:04","deactivatedAt":"2020-01-01 00:01:00"}]`
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
// TestPaginationOrderByMonthDocument tests the /api/pagination/{offset}/month/document/ endpoint
func TestPaginationOrderByMonthDocument(t *testing.T) {
	c := gomock.NewController(t)
	cm := databaseMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Db: cm,
	}
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
	cm.EXPECT().PaginationOrderByMonthDocument(gomock.Any()).Return(ex, nil)
	req, err := http.NewRequest("GET", "/api/pagination/{offset}/month/document/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.PaginationOrderByMonthDocument)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `[{"id":1,"referenceMonth":7,"referenceYear":2022,"document":"00000000000014","description":"insert","amount":10,"isActive":1,"createdAt":"2020-07-29 17:18:04","deactivatedAt":"2020-01-01 00:01:00"}]`
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
// TestPaginationOrderByYearDocument tests the /api/pagination/{offset}/year/document/ endpoint
func TestPaginationOrderByYearDocument(t *testing.T) {
	c := gomock.NewController(t)
	cm := databaseMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Db: cm,
	}
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
	cm.EXPECT().PaginationOrderByYearDocument(gomock.Any()).Return(ex, nil)
	req, err := http.NewRequest("GET", "/api/pagination/{offset}/year/document/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.PaginationOrderByYearDocument)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `[{"id":1,"referenceMonth":7,"referenceYear":2022,"document":"00000000000014","description":"insert","amount":10,"isActive":1,"createdAt":"2020-07-29 17:18:04","deactivatedAt":"2020-01-01 00:01:00"}]`
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
// TestPaginationByMonth tests the /api/pagination/{offset}/month/{referenceMonth} endpoint
func TestPaginationByMonth(t *testing.T) {
	c := gomock.NewController(t)
	cm := databaseMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Db: cm,
	}
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
	cm.EXPECT().PaginationByMonth(gomock.Any(),gomock.Any()).Return(ex, nil)
	req, err := http.NewRequest("GET", "/api/pagination/{offset}/month/{referenceMonth}", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.PaginationByMonth)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `[{"id":1,"referenceMonth":7,"referenceYear":2022,"document":"00000000000014","description":"insert","amount":10,"isActive":1,"createdAt":"2020-07-29 17:18:04","deactivatedAt":"2020-01-01 00:01:00"}]`
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
// TestPaginationByYear tests the /api/pagination/{offset}/year/{referenceYear} endpoint
func TestPaginationByYear(t *testing.T) {
	c := gomock.NewController(t)
	cm := databaseMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Db: cm,
	}
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
	cm.EXPECT().PaginationByYear(gomock.Any(),gomock.Any()).Return(ex, nil)
	req, err := http.NewRequest("GET", "/api/pagination/{offset}/year/{referenceYear}", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.PaginationByYear)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `[{"id":1,"referenceMonth":7,"referenceYear":2022,"document":"00000000000014","description":"insert","amount":10,"isActive":1,"createdAt":"2020-07-29 17:18:04","deactivatedAt":"2020-01-01 00:01:00"}]`
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
// TestPaginationByDocument tests the /api/pagination/{offset}/document/{document} endpoint
func TestPaginationByDocument(t *testing.T) {
	c := gomock.NewController(t)
	cm := databaseMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Db: cm,
	}
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
	cm.EXPECT().PaginationByDocument(gomock.Any(),gomock.Any()).Return(ex, nil)
	req, err := http.NewRequest("GET", "/api/pagination/{offset}/document/{document}", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.PaginationByDocument)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `[{"id":1,"referenceMonth":7,"referenceYear":2022,"document":"00000000000014","description":"insert","amount":10,"isActive":1,"createdAt":"2020-07-29 17:18:04","deactivatedAt":"2020-01-01 00:01:00"}]`
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
//