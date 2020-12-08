package config

import (
	yaml "gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

const (
	CurrentVer string = "v0.0.1"
)

var ClientConf ClientConfig

type ClientConfig struct {

}