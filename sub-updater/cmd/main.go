package main

import (
	"github.com/kmahyyg/phicomm-k2p-clash/sub-updater/config"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	log.Printf("Sub-Updater, Version: %s , Config: %s \n", config.CurrentVer, config.ConfigName)
	// create folder if not exists, require root
	if os.Geteuid() != 0 {
		log.Fatalln("Require root.")
	}
	// read program config
	curCwd, _ := os.Getwd()
	_, err := os.Stat(curCwd + "/" + config.ConfigName)
	if err != nil {
		log.Fatalln(err)
	}
	readInConf, err := ioutil.ReadFile(config.ConfigName)
	if err != nil {
		log.Fatalln(err)
	}
	config.ClientConf = &config.ClientConfig{}
	err = yaml.Unmarshal(readInConf, config.ClientConf)
	if err != nil {
		log.Fatalln(err)
	}
	// check network connectivity
	reqclient := http.DefaultClient
	reqclient.Timeout = 5*time.Second
	resp, err := reqclient.Get(config.ClientConf.CaptivePortal)
	if resp == nil || resp.StatusCode != 204 || err != nil {
		log.Fatalln(err, "Network is not connected.")
	}
	// create folder for config
	fdinfo, err1 := os.Stat(config.ClientConf.ClashConfPath)
	if err1 != nil || !fdinfo.IsDir() {
		err = os.Mkdir(config.ClientConf.ClashConfPath, 0755)
		if err != nil {
			log.Println("Clash Config Path CANNOT be created.")
			log.Fatalln(err)
		}
	}
	fdinfo, err2 := os.Stat(config.ClientConf.ClashCorePath)
	if err2 != nil || fdinfo.IsDir() {
		log.Println("Clash Binary Does NOT exists!")
		log.Fatalln(err2)
	}
	// start download config from isp
	var downloadedConf []byte
	if val, exists := config.ClientConf.NodeProvider[config.ClientConf.UseProvider]; exists {
		respcf, err := reqclient.Get(val)
		if err != nil {
			log.Fatalln(err)
		}
		// read response as []byte
		downloadedConf, err = ioutil.ReadAll(respcf.Body)
		if err != nil {
			log.Fatalln(err)
		}
		// initialize global var
		config.OriISPClashConf = &config.ClashConfig{}
		// start to unmarshal, might encounter issue like 403
		//TODO here
	} else {
		log.Fatalln("Current specified node provider is not existing!")
	}

}
