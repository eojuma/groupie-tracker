package main

import (
	"fmt"
	"groupie-tracker/backend"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	// Define custom template functions
	funcMap := template.FuncMap{
		"formatLoc": func(s string) string {
			s = strings.ReplaceAll(s, "_", " ")
			s = strings.ReplaceAll(s, "-", ", ")
			return strings.Title(s)
		},
	}

	// Parse templates with custom functions
	var err error
	backend.Templates, err = template.New("").Funcs(funcMap).ParseGlob("frontend/*.html")
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err)
	}

	// Serve static files (CSS, JS, images)
	fs := http.FileServer(http.Dir("frontend/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Register routes
	http.HandleFunc("/", backend.HomeHandler)
	//http.HandleFunc("/artist/", backend.ArtistHandler)

	port := ":8080"
	fmt.Printf("🎵 Groupie Tracker running at http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}