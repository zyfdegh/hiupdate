package conf

import (
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/hashicorp/hcl"
)

// ConfigFile is default config file path
const ConfigFile = "./hiupdate.conf"

var (
	once sync.Once

	// OptionsReady is a loaded object containing fields in config file.
	// Using defaults if not set.
	OptionsReady = &Options{}
)

// Options is the structure for config file.
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

// NewOptions returns new config
func NewOptions() *Options {
	return &Options{}
}

// Load loads options from config file
func (p *Options) Load() error {
	if _, err := os.Stat(ConfigFile); os.IsNotExist(err) {
		log.Printf("stat config file error: %v", err)
		return err
	}

	content, err := ioutil.ReadFile(ConfigFile)
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
