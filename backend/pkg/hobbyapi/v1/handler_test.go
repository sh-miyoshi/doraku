package hobbyapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/sh-miyoshi/doraku/pkg/hobbydb"
)

func TestGetAllHobbyHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/hobby/all", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllHobbyHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetTodayHobbyHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/hobby/today", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Initalize DB for API call
	hobbyFilePath := "../../../database/hobby.csv"
	descFilePath := "../../../database/description.csv"
	hobbydb.GetInst().Initialize(hobbyFilePath, descFilePath)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTodayHobbyHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetHobbyDetailsHandler(t *testing.T) {
	// Initalize DB for API call
	hobbyFilePath := "../../../database/hobby.csv"
	descFilePath := "../../../database/description.csv"
	hobbydb.GetInst().Initialize(hobbyFilePath, descFilePath)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Test correct data
	req, err := http.NewRequest("GET", "/api/v1/hobby/details/0", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/hobby/details/{id}", GetHobbyDetailsHandler)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
