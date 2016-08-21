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
	if req.Method != http.MethodPut {
		io.WriteString(w, "method not allowed")
		return
	}
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

	respUpdate, err := service.DoUpdate(reqUpdate)
	if err != nil {
		log.Printf("serve update request error: %v", err)
		return
	}
	body, err := json.Marshal(respUpdate)
	if err != nil {
		log.Printf("marshal object error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(body))
	return
}
