package mhttp

import (
	// "errors"
	"fmt"
	"log"
	"net/http"
	"os"
	// "path/filepath"
	"sync"
)

type IServer interface {
	InitializeFileServer() error
	InitializeHandlerFunctions() error
	ListenAndServe()
}

type ServerConfig struct {
	StaticDir string
	URL       string
	Port      string
	counter   int
	mutex     sync.Mutex
}

func NewServer(staticDir, url, port string) *ServerConfig {
	return &ServerConfig{
		StaticDir: staticDir,
		URL:       url,
		Port:      port,
	}
}

func (s *ServerConfig) InitializeFileServer() error {
	fmt.Println("Initializing static file server...")

	if err := validateFolder(s.StaticDir); err != nil {
		return fmt.Errorf("couldn't validate static folder: %w", err)
	}

	fs := http.FileServer(http.Dir(s.StaticDir))
	http.Handle("/", fs)

	fmt.Println("Static file handler initialized.")
	return nil
}

func (s *ServerConfig) InitializeHandlerFunctions() error {
	fmt.Println("Initializing dynamic route handlers...")

	if err := validateFolder(s.StaticDir); err != nil {
		return fmt.Errorf("couldn't validate static folder: %w", err)
	}

	http.HandleFunc("/", s.serveFile)
	http.HandleFunc("/increment", s.incrementCounter)
	http.HandleFunc("/hi", hiHandler)

	fmt.Println("Dynamic handlers initialized.")
	return nil
}

func (s *ServerConfig) ListenAndServe() {
	address := fmt.Sprintf("%s:%s", s.URL, s.Port)
	fmt.Printf("Server listening at http://%s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func validateFolder(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}
	return nil
}
