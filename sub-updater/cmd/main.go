package main

import (
	"github.com/kmahyyg/phicomm-k2p-clash/sub-updater/config"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
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
	reqclient.Timeout = 30*time.Second
	resp, err := reqclient.Get(config.ClientConf.CaptivePortal)
	if resp == nil || resp.StatusCode != 204 || err != nil {
		log.Fatalln(err, "Network is not connected.")
	}
	// create folder for config
	fdinfo, err1 := os.Stat(config.ClientConf.ClashConfPath)
	if fdinfo == nil|| err1 != nil || !fdinfo.IsDir() {
		err = os.Mkdir(config.ClientConf.ClashConfPath, 0755)
		if err != nil {
			log.Println("Clash Config Path CANNOT be created.")
			log.Fatalln(err)
		}
	}
	fdinfo, err2 := os.Stat(config.ClientConf.ClashCorePath)
	if fdinfo == nil || err2 != nil || fdinfo.IsDir() {
		log.Println("Clash Binary Does NOT exists!")
		log.Println("Core binary can be downloaded from: ", config.ClientConf.CoreDwnldURL)
		log.Fatalln(err2)
	}
	runtime.GC()
	// start download config from isp
	var downloadedConf []byte
	log.Println("Currently use: ", config.ClientConf.UseProvider)
	if val, exists := config.ClientConf.NodeProvider[config.ClientConf.UseProvider]; exists {
		respcf, err := reqclient.Get(val)
		if err != nil {
			log.Fatalln(err)
		}
		if respcf.StatusCode != 200 || respcf == nil {
			log.Printf("HTTP Error %d " ,respcf.StatusCode)
			log.Fatalln("Request Config From ISP is NOT successfully finished.")
		}
		// read response as []byte
		downloadedConf, err = ioutil.ReadAll(respcf.Body)
		if err != nil {
			log.Fatalln(err)
		}
		// initialize global var
		config.OriISPClashConf = &config.ClashConfig{}
		// start to unmarshal, might encounter issue like 403
		err = yaml.Unmarshal(downloadedConf, config.OriISPClashConf)
		if err != nil {
			log.Fatalln(err)
		}
		// dns might be empty in some way
		if config.OriISPClashConf.DNS == nil {
			config.OriISPClashConf.DNS = &config.ClashDNS{}
		}
		// detect if mmdb exists
		fdinfo, err = os.Stat(config.ClientConf.ClashConfPath + "/Country.mmdb")
		if err != nil || fdinfo ==nil ||fdinfo.IsDir() {
			// download mmdb
			respmmdb, err := reqclient.Get(config.ClientConf.MmdbDwnldURL)
			log.Println("If unsuccessful, please download mmdb from ", config.ClientConf.MmdbDwnldURL)
			if err == nil && respmmdb.StatusCode == 200 {
				countryMMDB, _ := ioutil.ReadAll(respmmdb.Body)
				err = ioutil.WriteFile(config.ClientConf.ClashConfPath + "/Country.mmdb", countryMMDB, 0644)
				if err != nil {
					log.Fatalln(err)
				}
			} else {
				log.Fatalln(err)
			}
		}
		// detect clash-dashboard
		cdashboardPath := config.ClientConf.ClashConfPath + "/" + config.ClientConf.OriginalClashConf.Controller.ExternalUI
		fdinfo, err = os.Stat(cdashboardPath)
		if fdinfo == nil || err != nil || !fdinfo.IsDir() {
			log.Println("Please Download Clash-DashBoard from " + config.ClientConf.DashboardDwnldURL)
			log.Println("After that, extract to " + cdashboardPath + "/")
			log.Fatalln(err)
		}
		runtime.GC()
		// manipulate config
		err = ManipulateClashConf(config.ClientConf, config.OriISPClashConf)
		if err != nil {
			log.Fatalln(err)
		}
		// write to disk
		modifiedConf, err := yaml.Marshal(config.OriISPClashConf)
		if err != nil {
			log.Fatalln(err)
		}
		err = ioutil.WriteFile(config.ClientConf.ClashConfPath + "/config.yaml", modifiedConf, 0644)
		if err != nil {
			log.Fatalln(err)
		}
		runtime.GC()
		// start up clash
		RunClash()
	} else {
		log.Fatalln("Current specified node provider is not existing!")
	}

}


func RunClash(){
	proc := exec.Command(config.ClientConf.ClashCorePath, "-d", config.ClientConf.ClashConfPath)
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr
	err := proc.Start()
	if err != nil {
		log.Fatalln(err)
	}
}

func ManipulateClashConf(subconf *config.ClientConfig, ispconf *config.ClashConfig) error {
	// append rule first
	ispconf.NodeNRoute.Rule = append(subconf.Rules2Insert, ispconf.NodeNRoute.Rule...)
	// modify inbound first
	ispconf.Inbound = subconf.OriginalClashConf.Inbound
	// modify general
	ispconf.General = subconf.OriginalClashConf.General
	// then controller
	ispconf.Controller = subconf.OriginalClashConf.Controller
	// then dns
	data, err := yaml.Marshal(subconf.OriginalClashConf.DNS)
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(data, ispconf.DNS)
	if err != nil {
		log.Fatalln(err)
	}
	// then return
	return nil
}

