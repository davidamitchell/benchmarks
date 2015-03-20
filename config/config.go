package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	ServerURL string     `json:"admin_url"`
	Database  string     `json:"dbpath"`
}

var Configuration Config

func init() {
	configFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Printf("There is a problem with the file: %v\n", err)
	}
	json.Unmarshal(configFile, &Configuration)
}
