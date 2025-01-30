package mhttp

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// serveFile serves files from the static directory or returns 404 if file not found
func (s *ServerConfig) serveFile(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(s.StaticDir, r.URL.Path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.Error(w, "404 - File not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "500 - Internal server error", http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, r, path)
}

// incrementCounter increases the counter and sends the current value as a response
func (s *ServerConfig) incrementCounter(w http.ResponseWriter, r *http.Request) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.counter++
	if _, err := fmt.Fprint(w, strconv.Itoa(s.counter)); err != nil {
		http.Error(w, "500 - Failed to write response", http.StatusInternalServerError)
	}
}

// hiHandler responds with a simple greeting
func hiHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprint(w, "Hi"); err != nil {
		http.Error(w, "500 - Failed to write response", http.StatusInternalServerError)
	}
}
