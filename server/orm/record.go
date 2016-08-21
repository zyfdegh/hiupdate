package orm

import (
	"encoding/json"
	"log"

	"github.com/zyfdegh/hiupdate/server/dao"
	"github.com/zyfdegh/hiupdate/server/entity"
)

func RecordUpdate(record *entity.Record) (err error) {
	data, err := json.Marshal(record)
	if err != nil {
		log.Printf("marshal json error: %v", err)
		return
	}
	err = dao.UnqliteReady.Update([]byte(record.ID), data)
	if err != nil {
		log.Printf("dao update error: %v", err)
		return
	}
	return
}
