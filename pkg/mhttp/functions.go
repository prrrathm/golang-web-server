package mhttp

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
)

// serveFile serves the index.html or the requested file within the static directory
func (s *ServerConfig) serveFile(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(s.StaticDir, r.URL.Path)
	http.ServeFile(w, r, path)
}

// incrementCounter increases the counter and sends the current value as a response
func (s *ServerConfig) incrementCounter(w http.ResponseWriter, r *http.Request) {
	s.mutex.Lock()
	s.counter++
	fmt.Fprint(w, strconv.Itoa(s.counter))
	s.mutex.Unlock()
}

// hiHandler responds with a simple greeting
func hiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi")
}
