package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/zyfdegh/hiupdate/server/service"
)

// GetForgotPersons serves GET /persons/forgot?date="20160903"
func GetForgotPersons(w http.ResponseWriter, req *http.Request) {
	var date = req.FormValue("date")
	resp, err := service.GetForgotPersons(date)
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
