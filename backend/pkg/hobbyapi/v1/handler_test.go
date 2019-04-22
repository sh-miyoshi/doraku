package hobbyapi

import (
	"fmt"
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

func TestGetRecommendedHobbyHandler(t *testing.T) {
	// Initalize DB for API call
	hobbyFilePath := "../../../database/hobby.csv"
	descFilePath := "../../../database/description.csv"
	hobbydb.GetInst().Initialize(hobbyFilePath, descFilePath)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetRecommendedHobbyHandler)

	req, err := http.NewRequest("GET", "/api/v1/hobby/recommended", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// TODO(test all input pattern)
}

func TestGetHobbyDetailsHandler(t *testing.T) {
	// Initalize DB for API call
	hobbyFilePath := "../../../database/hobby.csv"
	descFilePath := "../../../database/description.csv"
	hobbydb.GetInst().Initialize(hobbyFilePath, descFilePath)

	const BasePath = "/api/v1/hobby/details"

	// Test Cases
	tt := []struct {
		routeVariable    string
		expectStatusCode int
	}{
		{"0", http.StatusOK},
		{"21", http.StatusOK},
		{"22", http.StatusNotFound},
		{"-1", http.StatusNotFound},
		{"test", http.StatusNotFound},
		{"", http.StatusNotFound},
	}

	for _, tc := range tt {
		path := fmt.Sprintf("%s/%s", BasePath, tc.routeVariable)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()
		router.HandleFunc(BasePath+"/{id}", GetHobbyDetailsHandler)
		router.ServeHTTP(rr, req)

		// In this case, our MetricsHandler returns a non-200 response
		// for a route variable it doesn't know about.
		if rr.Code != tc.expectStatusCode {
			t.Errorf("handler should have failed on routeVariable %s: got %v want %v",
				tc.routeVariable, rr.Code, http.StatusOK)
		}
	}
}
