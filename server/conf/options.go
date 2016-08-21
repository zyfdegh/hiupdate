package conf

import (
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/hashicorp/hcl"
)

const CONFIG_FILE = "./hiupdate.conf"

var (
	once         sync.Once
	OptionsReady = &Options{}
)

type Options struct {
	Port   int    `hcl:"port"`
	DBFile string `hcl:"db"`
}

// load config
func init() {
	once.Do(func() {
		options := NewOptions()
		if err := options.Load(); err != nil {
			log.Printf("load options error: %v, using default", err)
			OptionsReady = &defaultOptions
			return
		}
		OptionsReady = options
	})
}

// New returns new config
func NewOptions() *Options {
	return &Options{}
}

// Load loads options from config file
func (p *Options) Load() error {
	if _, err := os.Stat(CONFIG_FILE); os.IsNotExist(err) {
		log.Printf("stat config file error: %v", err)
		return err
	}

	content, err := ioutil.ReadFile(CONFIG_FILE)
	if err != nil {
		log.Printf("read config file error: %v", err)
		return err
	}

	if err := hcl.Decode(p, string(content)); err != nil {
		log.Printf("decode hcl error: %v", err)
		return err
	}
	return nil
}
