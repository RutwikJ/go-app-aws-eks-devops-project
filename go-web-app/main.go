package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Restrict root path so it doesn't accidentally catch broken links
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "static/home.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/contact.html")
}

func submitContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Read and print the form inputs straight to your console terminal
	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")
	fmt.Printf("Message From: %s (%s)\nNote: %s\n---\n", name, email, message)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h3>Thank you! Message received.</h3><a href='/'>Return Home</a>")
}

func main() {
	// Route mappings matching your static folder layouts
	http.HandleFunc("/", homeHandler) // Changed from "/home" to "/" for default loading
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)
	http.HandleFunc("/submit-contact", submitContactHandler) // Added to process the form

	fmt.Println("Server starting smoothly at http://localhost:8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
