package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	QQ   string
	Addr string
	Port int
	Url  string
}

var AppConfig Config

func Init(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("config read error: %v", err)

	}
	err = yaml.Unmarshal(file, &AppConfig)
	if err != nil {
		return fmt.Errorf("unmarshal error: %v", err)
	}
	return nil
}
