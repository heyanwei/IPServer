package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type ConfigData struct {
	Url  string `yaml:"url"`
	Name string `yaml:"name"`
}

func (c *ConfigData) getConf() *ConfigData {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Println("Open config file error: " + err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Println("config file Format error: " + err.Error())
	}
	return c
}

func main() {
	var c ConfigData
	conf := c.getConf()

	if len(conf.Url) < 5 {
		log.Println("config url error")
		return
	}
	if len(conf.Name) < 2 {
		log.Println("config Name error")
		return
	}

	log.Println("name: " + conf.Name + ", url: " + conf.Url)

	for {
		report(conf.Name, conf.Url)
		time.Sleep(time.Duration(60) * time.Second)
	}
}

func report(name string, url string) {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader("device="+name))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
}
