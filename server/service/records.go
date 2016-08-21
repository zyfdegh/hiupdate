package service

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/zyfdegh/hiupdate/server/entity"
	"github.com/zyfdegh/hiupdate/server/orm"
	"github.com/zyfdegh/hiupdate/server/util"
)

func checkReqUpdate(req *entity.ReqUpdate) (errmsg string) {
	if len(strings.TrimSpace(req.Name)) == 0 {
		return "name not set"
	}
	if len(strings.TrimSpace(req.Todo)) == 0 {
		return "todo not set"
	}
	if len(strings.TrimSpace(req.Issue)) == 0 {
		return "issue not set"
	}
	return
}

func DoUpdate(req *entity.ReqUpdate) (resp *entity.RespUpdate, err error) {
	util.PrintPretty(req, "req")
	resp = &entity.RespUpdate{}

	errmsg := checkReqUpdate(req)
	if errmsg != "" {
		resp.Success = false
		resp.ErrNo = http.StatusBadRequest
		resp.Errmsg = errmsg
		return
	}

	record := &entity.Record{}
	record.Person.Name = req.Name
	record.Content.Todo = req.Todo
	record.Content.Issue = req.Issue
	date := util.GetDate(time.Now())
	record.Date = date
	record.ID = util.GenerateID(req.Name, date)

	err = orm.RecordUpdate(record)
	if err != nil {
		log.Printf("record update error: %v", err)
		resp.Success = false
		resp.ErrNo = http.StatusInternalServerError
		resp.Errmsg = "record update error"
		return
	}

	resp.Success = true
	resp.Data = *req
	return
}
