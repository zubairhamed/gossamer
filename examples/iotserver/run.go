package main

import (
	"github.com/zubairhamed/gossamer/store"
	"github.com/zubairhamed/gossamer/webapp/server"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		panic(err)
	}

	s := server.NewServer(cfg.Host, cfg.Port)

	s.UseStore(store.NewMongoStore(cfg.MongoDatabase.Host, cfg.MongoDatabase.Database))
	s.Start()
}

func parseConfig() (*AppConfig, error) {
	cfgPath, err := filepath.Abs("./config.yml")
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	var cfg *AppConfig = &AppConfig{}
	err = yaml.Unmarshal(buf, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

type AppConfig struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`

	MongoDatabase struct {
		Host     string `yaml:"host"`
		Database string `yaml:"database"`
	} `yaml:"datasource-mongo"`
}
