package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type requestHandler struct{}

func (h *requestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if *fileName != "" {
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", *fileName))
	}
	w.Header().Set("Content-type", *contentType)

	for {
		data := make([]byte, 512)
		n, err := os.Stdin.Read(data)
		if err != nil && err != io.EOF {
			log.Printf("Error reading input: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if n == 0 {
			break
		}
		_, writeErr := w.Write(data[:n])
		if writeErr != nil {
			log.Printf("Error writing response: %s\n", writeErr)
			return
		}
	}
}

var (
	port        = flag.Int("p", 8080, "Port (default: 8080)")
	address     = flag.String("a", "0.0.0.0", "Address (default: 0.0.0.0)")
	fileName    = flag.String("f", "", "Set filename header")
	contentType = flag.String("c", "application/octet-stream", "Set content-type header (default: application/octet-stream)")
)

func main() {
	flag.Parse()

	http.Handle("/", &requestHandler{})

	addr := fmt.Sprintf("%s:%d", *address, *port)
	log.Printf("Server starting on %s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %s\n", err)
	}
}
