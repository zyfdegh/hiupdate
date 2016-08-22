package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/zyfdegh/hiupdate/server/conf"
	"github.com/zyfdegh/hiupdate/server/entity"
	"github.com/zyfdegh/hiupdate/server/service"
)

var port = conf.OptionsReady.Port

// Serve do router mapping and start http server
func Serve() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/update/", handleUpdate)

	s := &http.Server{Addr: fmt.Sprintf(":%d", port)}
	log.Printf("server start on localhost:%d", port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("start server error: %v", err)
	}
}

func handleRoot(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "static/html/index.html")
}

func handleUpdate(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		getUpdate(w, req)
	case http.MethodPut:
		putUpdate(w, req)
	default:
		io.WriteString(w, "method not allowed")
	}
	return
}

// GET /update?name="Zhang"
func getUpdate(w http.ResponseWriter, req *http.Request) {
	var name = req.FormValue("name")
	var date = req.FormValue("date")
	fmt.Println(name)
	fmt.Println(date)
	resp, err := service.GetUpdate(name, date)
	if err != nil {
		log.Printf("serve get update request error: %v", err)
		return
	}
	body, err := json.Marshal(resp)
	if err != nil {
		log.Printf("marshal object error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(body))
}

// PUT /update
// Body entity.ReqUpdate
func putUpdate(w http.ResponseWriter, req *http.Request) {
	var reqUpdate = &entity.ReqUpdate{}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("read request error: %v", err)
		return
	}
	err = json.Unmarshal(data, reqUpdate)
	if err != nil {
		log.Printf("unmarshal to object error: %v", err)
		return
	}

	resp, err := service.PutUpdate(reqUpdate)
	if err != nil {
		log.Printf("serve put update request error: %v", err)
		return
	}
	body, err := json.Marshal(resp)
	if err != nil {
		log.Printf("marshal object error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(body))
	return
}
