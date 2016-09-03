package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/zyfdegh/hiupdate/server/service"
)

// GetReport serves GET /report?date=20160825
func GetReport(w http.ResponseWriter, req *http.Request) {
	var date = req.FormValue("date")
	report, err := service.GetReportText(date)
	fmt.Println(report)
	if err != nil {
		log.Printf("serve get report error: %v", err)
		return
	}
	body, err := json.Marshal(report)
	if err != nil {
		log.Printf("marshal object error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(body))
	return
}
