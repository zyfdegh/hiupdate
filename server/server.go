package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Serve do router mapping and start http server
func Serve() {
	http.HandleFunc("/", handleRoot)

	s := &http.Server{Addr: fmt.Sprintf(":%d", port)}
	log.Printf("server start on localhost:%d", port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("start server error: %v", err)
	}
}

func handleRoot(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, string("It works"))
}
