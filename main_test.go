package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

// TestMain sets up the environment before running tests
func TestMain(m *testing.M) {
	// Create a temporary mock static directory if running tests outside the environment
	_ = os.MkdirAll("static", os.ModePerm)
	_ = os.WriteFile("static/home.html", []int8{60, 104, 49, 62, 66, 73, 83, 84, 82, 79, 60, 47, 104, 49, 62} /* <h1>BISTRO</h1> */, 0644)
	_ = os.WriteFile("static/about.html", []byte("About page content"), 0644)
	_ = os.WriteFile("static/contact.html", []byte("Contact page content"), 0644)

	// Execute tests
	code := m.Run()
	os.Exit(code)
}

// TestHomePage verifies the root route serves the static homepage
func TestHomePage(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homePage)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

// TestAboutPage verifies the about page endpoint resolves smoothly
func TestAboutPage(t *testing.T) {
	req, err := http.NewRequest("GET", "/about", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(aboutPage)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

// TestContactPage verifies the contact page endpoint works
func TestContactPage(t *testing.T) {
	req, err := http.NewRequest("GET", "/contact", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(contactPage)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

// TestSubmitContactHandler verifies the form submission endpoint processes POST data safely
func TestSubmitContactHandler(t *testing.T) {
	formData := url.Values{}
	formData.Set("name", "Jane Doe")
	formData.Set("email", "jane@example.com")
	formData.Set("message", "Booking inquiry")

	req, err := http.NewRequest("POST", "/submit-contact", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(submitContactHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedText := "Thank you! Message received."
	if !strings.Contains(rr.Body.String(), expectedText) {
		t.Errorf("handler returned unexpected body: missing success confirmation text")
	}
}
