package dao

import (
	"log"
	"sync"

	"github.com/svkior/unqlitego"
	"github.com/zyfdegh/hiupdate/server/conf"
)

var (
	// UnqliteReady is the connected unqlite database instance
	UnqliteReady *Unqlite
	once         sync.Once
)

// Unqlite contains unqlite database instance
type Unqlite struct {
	db *unqlitego.Database
}

func init() {
	once.Do(func() {
		UnqliteReady = NewUnqlite()
		// New
		db, err := unqlitego.NewDatabase(conf.OptionsReady.DBFile)
		// defer db.Close()
		if err != nil {
			log.Fatalf("new unqlite db error: %v", err)
			return
		}
		UnqliteReady.db = db
	})
}

// NewUnqlite returns new unqlite struct
func NewUnqlite() *Unqlite {
	return &Unqlite{}
}

// Insert store data to unqlite, same as update
func (u *Unqlite) Insert(k []byte, v []byte) error {
	return u.Update(k, v)
}

// Query query data from unqlite
func (u *Unqlite) Query(k []byte) (*[]byte, error) {
	bytes, err := u.db.Fetch(k)
	if err != nil {
		log.Printf("unqlite fetch error: %v", err)
		return nil, err
	}
	return &bytes, nil
}

// Update store data to unqlite
func (u *Unqlite) Update(k []byte, v []byte) error {
	err := u.db.Store(k, v)
	if err != nil {
		log.Printf("unqlite store error: %v", err)
		return err
	}
	u.commit()
	return nil
}

// Delete remove data from unqlite
func (u *Unqlite) Delete(k []byte, v []byte) error {
	err := u.db.Delete(k)
	if err != nil {
		log.Printf("unqlite delete error: %v", err)
		return err
	}
	u.commit()
	return nil
}

func (u *Unqlite) commit() error {
	err := u.db.Commit()
	if err != nil {
		log.Printf("unqlite commit error: %v", err)
		return err
	}
	return nil
}
