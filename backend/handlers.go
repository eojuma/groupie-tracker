package backend

import (
	"html/template"
	"log"
	"net/http"
	//"strconv"
	//"strings"
)

// Templates holds all parsed templates
var Templates *template.Template

// InitTemplates parses all HTML templates with custom functions
// func InitTemplates(pattern string) error {
// 	funcMap := template.FuncMap{
// 		"formatLoc": FormatLocation,
// 	}
// 	var err error
// 	Templates, err = template.New("").Funcs(funcMap).ParseGlob(pattern)
// 	return err
// }

// HomeHandler serves the home page with optional search
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	artists, err := GetArtistsWithRelations()
	if err != nil {
		log.Printf("Error fetching artists: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	query := r.URL.Query().Get("query")
	filtered := SearchArtists(query, artists)

	data := PageData{
		Artists: filtered,
		Query:   query,
		Total:   len(filtered),
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := Templates.ExecuteTemplate(w, "index.html", data); err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// ArtistHandler serves the artist detail page
// func ArtistHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		ErrorHandler(w, r, http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Extract ID from URL: /artist/1
// 	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/artist/"), "/")
// 	if len(parts) == 0 || parts[0] == "" {
// 		ErrorHandler(w, r, http.StatusNotFound)
// 		return
// 	}

// 	id, err := strconv.Atoi(parts[0])
// 	if err != nil || id < 1 {
// 		ErrorHandler(w, r, http.StatusNotFound)
// 		return
// 	}

// 	artists, err := GetRelations()
// 	if err != nil {
// 		log.Printf("Error fetching artists: %v", err)
// 		ErrorHandler(w, r, http.StatusInternalServerError)
// 		return
// 	}

// 	// Find the artist by ID
// 	var artist *ArtistDetail
// 	for i := range artists {
// 		if artists[i].ID == id {
// 			artist = &artists[i]
// 			break
// 		}
// 	}

// 	if artist == nil {
// 		ErrorHandler(w, r, http.StatusNotFound)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	if err := Templates.ExecuteTemplate(w, "artist.html", artist); err != nil {
// 		log.Printf("Error rendering template: %v", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 	}
// }

// ErrorHandler renders a custom error page
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	messages := map[int]string{
		http.StatusNotFound:            "The page you're looking for doesn't exist.",
		http.StatusInternalServerError: "Something went wrong on our end. Please try again later.",
		http.StatusMethodNotAllowed:    "This request method is not allowed.",
		http.StatusBadRequest:          "Bad request.",
	}

	msg, ok := messages[status]
	if !ok {
		msg = "An unexpected error occurred."
	}

	data := ErrorData{
		StatusCode: status,
		Message:    msg,
	}

	if err := Templates.ExecuteTemplate(w, "error.html", data); err != nil {
		log.Printf("Error rendering error template: %v", err)
		http.Error(w, msg, status)
	}
}