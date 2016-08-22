package service

import (
	"log"
	"net/http"
	"strings"

	"github.com/zyfdegh/hiupdate/server/entity"
	"github.com/zyfdegh/hiupdate/server/orm"
	"github.com/zyfdegh/hiupdate/server/util"
)

func GetYesterdayRecord(name string) (resp *entity.RespGetUpdate, err error) {
	return GetUpdate(name, util.GetYesterday())
}

func checkReqGetUpdate(name, date string) (errmsg string) {
	if len(strings.TrimSpace(name)) == 0 {
		return "name not set"
	}
	if len(strings.TrimSpace(date)) == 0 {
		return "name not set"
	}
	return ""
}

func GetUpdate(name, date string) (resp *entity.RespGetUpdate, err error) {
	resp = &entity.RespGetUpdate{}
	errmsg := checkReqGetUpdate(name, date)
	if errmsg != "" {
		resp.Success = false
		resp.ErrNo = http.StatusBadRequest
		resp.Errmsg = errmsg
		return
	}
	id := util.GenerateID(name, date)
	record, err := orm.QueryRecord(id)
	if err != nil {
		log.Printf("service query record error: %v", err)
		resp.Success = false
		resp.ErrNo = http.StatusInternalServerError
		resp.Errmsg = "service query record error"
		return
	}
	resp.Success = true
	resp.Data = *record
	return
}

func checkReqPutUpdate(req *entity.ReqUpdate) string {
	if len(strings.TrimSpace(req.Name)) == 0 {
		return "name not set"
	}
	if len(strings.TrimSpace(req.Todo)) == 0 {
		return "todo not set"
	}
	if len(strings.TrimSpace(req.Issue)) == 0 {
		return "issue not set"
	}
	return ""
}

func PutUpdate(req *entity.ReqUpdate) (resp *entity.RespPutUpdate, err error) {
	resp = &entity.RespPutUpdate{}

	errmsg := checkReqPutUpdate(req)
	if errmsg != "" {
		resp.Success = false
		resp.ErrNo = http.StatusBadRequest
		resp.Errmsg = errmsg
		return
	}

	record := &entity.Record{}
	record.Person.Name = req.Name
	record.Content.Done = req.Done
	record.Content.Todo = req.Todo
	record.Content.Issue = req.Issue
	date := util.GetToday()
	record.Date = date
	record.ID = util.GenerateID(req.Name, date)

	err = orm.UpdateRecord(record)
	if err != nil {
		log.Printf("service record update error: %v", err)
		resp.Success = false
		resp.ErrNo = http.StatusInternalServerError
		resp.Errmsg = "service record update error"
		return
	}

	resp.Success = true
	resp.Data = *record
	return
}
