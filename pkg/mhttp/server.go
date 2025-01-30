package mhttp

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

type IServer interface {
	InitializeFileServer() error
	InitializeHandlerFunctions() error
	ListenAndServe() error
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
		return fmt.Errorf("static folder validation failed: %w", err)
	}

	fs := http.FileServer(http.Dir(s.StaticDir))
	http.Handle("/", fs)

	fmt.Println("Static file handler initialized.")
	return nil
}

func (s *ServerConfig) InitializeHandlerFunctions() error {
	fmt.Println("Initializing dynamic route handlers...")

	if err := validateFolder(s.StaticDir); err != nil {
		return fmt.Errorf("static folder validation failed: %w", err)
	}

	http.HandleFunc("/", s.serveFile)
	http.HandleFunc("/increment", s.incrementCounter)
	http.HandleFunc("/hi", hiHandler)

	fmt.Println("Dynamic handlers initialized.")
	return nil
}

func (s *ServerConfig) ListenAndServe() error {
	address := fmt.Sprintf("%s:%s", s.URL, s.Port)
	fmt.Printf("Server listening at http://%s\n", address)

	if err := http.ListenAndServe(address, nil); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}

func validateFolder(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("directory does not exist: %s", path)
		}
		return fmt.Errorf("error accessing directory %s: %w", path, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}
	return nil
}
