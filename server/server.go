package server

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/zyfdegh/hiupdate/server/api"
	"github.com/zyfdegh/hiupdate/server/conf"
)

var port = conf.OptionsReady.Port

// Serve do router mapping and start http server
func Serve() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/record/", handleRecord)
	http.HandleFunc("/report/", handleReport)
	http.HandleFunc("/persons/forgot/", handlePersonsForgot)

	s := &http.Server{Addr: fmt.Sprintf(":%d", port)}
	log.Printf("server start on localhost:%d", port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("start server error: %v", err)
	}
}

func handleRoot(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "static/html/index.html")
}

func handleRecord(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		api.GetRecord(w, req)
	case http.MethodPut:
		api.PutRecord(w, req)
	default:
		io.WriteString(w, "method not allowed")
	}
	return
}

// GET /report?date=20160825
func handleReport(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		api.GetReport(w, req)
	default:
		io.WriteString(w, "method not allowed")
	}
	return
}

func handlePersonsForgot(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		api.GetForgotPersons(w, req)
	default:
		io.WriteString(w, "method not allowed")
	}
	return
}
