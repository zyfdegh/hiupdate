package conf

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/hcl"
)

const CONFIG_FILE = "./hiupdate.conf"

var Opts = &Options{}

type Options struct {
	Port int    `hcl:"port"`
	DB   string `hcl:"db"`
}

// load config
func init() {
	options := New()
	if err := options.Load(); err != nil {
		log.Printf("load options error: %v, using default", err)
		Opts = &defaultOptions
		return
	}
	Opts = options
	return
}

// New returns new config
func New() *Options {
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
