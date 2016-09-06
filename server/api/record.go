package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/zyfdegh/hiupdate/server/entity"
	"github.com/zyfdegh/hiupdate/server/service"
)

// GetRecord serves GET /record?name=ZhangSan&date=20160906
func GetRecord(w http.ResponseWriter, req *http.Request) {
	var name = req.FormValue("name")
	var date = req.FormValue("date")
	fmt.Println(name)
	fmt.Println(date)
	resp, err := service.GetRecord(name, date)
	if err != nil {
		log.Printf("serve get record request error: %v", err)
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

// PutRecord servers PUT /record
// Body entity.ReqRecord
func PutRecord(w http.ResponseWriter, req *http.Request) {
	var reqRecord = &entity.ReqRecord{}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("read request error: %v", err)
		return
	}
	err = json.Unmarshal(data, reqRecord)
	if err != nil {
		log.Printf("unmarshal to object error: %v", err)
		return
	}

	fmt.Println("/**********")
	fmt.Printf("Name: %s\n", reqRecord.Name)
	fmt.Printf("Done: %s\n", reqRecord.Done)
	fmt.Printf("Todo: %s\n", reqRecord.Todo)
	fmt.Printf("Issue: %s\n", reqRecord.Issue)
	fmt.Println("**********/")

	resp, err := service.PutRecord(reqRecord)
	if err != nil {
		log.Printf("serve put record request error: %v", err)
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
