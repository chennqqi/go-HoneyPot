package config

import (
	"fmt"
	"log"
	"os"

	"github.com/chennqqi/goutils/yamlconfig"
)

// Config is the struct for all configurable data
type Config struct {
	DB     Database   `json:"db" yaml:"db"`
	TCP    TCP        `json:"tcp" yaml:"tcp"`
	Http   RemoteHttp `json:"http" yaml:"http"`
	Report string     `json:"report" yaml:"report"`
}

// Database is the config struct for the database
type Database struct {
	Driver   string `json:"driver" yaml:"driver"`
	Host     string `json:"host" yaml:"host"`
	Name     string `json:"name" yaml:"name"`
	Username string `json:"user" yaml:"user"`
	Password string `json:"pass" yaml:"pass"`
	Port     string `json:"port" yaml:"port"`
}

// Database is the config struct for the database
type RemoteHttp struct {
	Uri string `json:"uri" yaml:"uri"`
}

// TCP is the config struct for the tcp server
type TCP struct {
	Ports []string `json:"ports" yaml:"ports"`
}

// Read reads the configuration file and returns a struct of it
func Read() (Config, error) {
	var cfg Config
	err := yamlconfig.Load(&cfg, "")
	if os.IsNotExist(err) {
		fmt.Println("No config found, make default")
		yamlconfig.Save(cfg, "")
		return cfg, err
	} else if err != nil {
		log.Fatalf("Could not read config: %v", err)
	}
	return cfg, nil
}
