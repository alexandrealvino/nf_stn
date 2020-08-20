package tests

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"
	"net/http"
	"nf_stn/adapter"
	tokenMock "nf_stn/authentication/mock"
	"nf_stn/entities"
	"testing"
)

// TestInit tests the Init function
func TestInitToken(t *testing.T) {
	c := gomock.NewController(t)
	cm := tokenMock.NewMockDbClient(c)
	cm.EXPECT().Init().Return()
}
// TestExtractToken tests the ExtractToken function
func TestExtractToken(t *testing.T) {
	c := gomock.NewController(t)
	cm := tokenMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Au: cm,
	}
	ex := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjM4ODYzOWQ4LWQ5NjEtNGM2Ni04MzEyLTI2ZDAxM2NlNjMzOCIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTU5NjM4OTgyMiwidXNlcl9pZCI6MSwidXNlcm5hbWUiOiJ1c2VybmFtZSJ9.cefD1NQ9jirg4vnTudLVi0_pE0VTFZVNvVTexTGSxnk`
	cm.EXPECT().ExtractToken(gomock.Any()).Return(ex)
	req, err := http.NewRequest("POST", "/api/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := routes.Au.ExtractToken(req)
	// Check the response body is what we expect.
	expected := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjM4ODYzOWQ4LWQ5NjEtNGM2Ni04MzEyLTI2ZDAxM2NlNjMzOCIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTU5NjM4OTgyMiwidXNlcl9pZCI6MSwidXNlcm5hbWUiOiJ1c2VybmFtZSJ9.cefD1NQ9jirg4vnTudLVi0_pE0VTFZVNvVTexTGSxnk`
	if rr != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr, expected)
	}
}
// TestVerifyToken tests the VerifyToken function
func TestVerifyToken(t *testing.T) {
	c := gomock.NewController(t)
	cm := tokenMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Au: cm,
	}
	var ex *jwt.Token
	cm.EXPECT().VerifyToken(gomock.Any()).Return(ex, nil)
	req, err := http.NewRequest("POST", "/api/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr, err := routes.Au.VerifyToken(req)
	// Check the response body is what we expect.
	var expected *jwt.Token
	if rr != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr, expected)
	}
}
// TestTokenValid tests the TokenValid function
func TestTokenValid(t *testing.T) {
	c := gomock.NewController(t)
	cm := tokenMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Au: cm,
	}
	//var ex *jwt.Token
	cm.EXPECT().TokenValid(gomock.Any()).Return( nil)
	req, err := http.NewRequest("POST", "/api/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	err = routes.Au.TokenValid(req)
	var expected error
	// Check the response body is what we expect.
	if err != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			err, expected)
	}
}
// TestExtractTokenMetadata tests the ExtractTokenMetadata function
func TestExtractTokenMetadata(t *testing.T) {
	c := gomock.NewController(t)
	cm := tokenMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Au: cm,
	}
	var ex *entities.AccessDetails
	cm.EXPECT().ExtractTokenMetadata(gomock.Any()).Return(ex, nil)
	req, err := http.NewRequest("POST", "/api/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr, err := routes.Au.ExtractTokenMetadata(req)
	// Check the response body is what we expect.
	var expected *entities.AccessDetails
	if rr != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr, expected)
	}
}
// TestFetchAuth tests the FetchAuth function
func TestFetchAuth(t *testing.T) {
	c := gomock.NewController(t)
	cm := tokenMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Au: cm,
	}
	ex :=uint64(1)
	cm.EXPECT().FetchAuth(gomock.Any()).Return(ex, nil)
	//req, err := http.NewRequest("POST", "/api/login", nil)
	//if err != nil {
	//	t.Fatal(err)
	//}
	var es *entities.AccessDetails
	rr, _ := routes.Au.FetchAuth(es)
	// Check the response body is what we expect.
	//var expected *entities.AccessDetails
	expected := uint64(1)
	if rr != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr, expected)
	}
}
// TestCreateToken tests the CreateToken function
func TestCreateToken(t *testing.T) {
	c := gomock.NewController(t)
	cm := tokenMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Au: cm,
	}
	var ex *entities.TokenDetails
	cm.EXPECT().CreateToken(gomock.Any(),gomock.Any()).Return(ex, nil)
	i := uint64(1)
	user := "username"
	es , _ := routes.Au.CreateToken(i,user)
	// Check the response body is what we expect.
	var expected  *entities.TokenDetails
	if es != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			es, expected)
	}
}
// TestCreateAuth tests the CreateAuth function
func TestCreateAuth(t *testing.T) {
	c := gomock.NewController(t)
	cm := tokenMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Au: cm,
	}
	var ex *entities.TokenDetails
	i:=uint64(1)
	cm.EXPECT().CreateAuth(gomock.Any(),gomock.Any()).Return(nil)
	err := routes.Au.CreateAuth(i,ex)
	// Check the response body is what we expect.
	var expected  error
	if err != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			err, expected)
	}
}
// TestDeleteAuth tests the DeleteAuth function
func TestDeleteAuth(t *testing.T) {
	c := gomock.NewController(t)
	cm := tokenMock.NewMockDbClient(c)
	routes := adapter.Routes{
		Au: cm,
	}
	ex := "00000000000044"
	i:=int64(1)
	cm.EXPECT().DeleteAuth(gomock.Any()).Return(i,nil)
	del, _ := routes.Au.DeleteAuth(ex)
	// Check the response body is what we expect.
	expected := int64(1)
	if del != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			del, expected)
	}
}
//