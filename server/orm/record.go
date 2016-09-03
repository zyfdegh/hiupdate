package orm

import (
	"encoding/json"
	"log"

	"github.com/zyfdegh/hiupdate/server/dao"
	"github.com/zyfdegh/hiupdate/server/entity"
)

// QueryRecord query record from unqlite with 'id' and return as entity
func QueryRecord(id string) (record *entity.Record, err error) {
	record = &entity.Record{}
	data, err := dao.UnqliteReady.Query([]byte(id))
	if err != nil {
		log.Printf("orm query record error: %v", err)
		return
	}
	err = json.Unmarshal(*data, record)
	if err != nil {
		log.Printf("orm unmarshal json error: %v", err)
		return
	}
	return
}

// UpdateRecord updates a record in unqlite
func UpdateRecord(record *entity.Record) (err error) {
	data, err := json.Marshal(record)
	if err != nil {
		log.Printf("orm marshal json error: %v", err)
		return
	}
	err = dao.UnqliteReady.Update([]byte(record.ID), data)
	if err != nil {
		log.Printf("orm update record error: %v", err)
		return
	}
	return
}
