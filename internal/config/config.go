package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	AppPort  int      `json:"appPort"`
	Database Database `json:"database"`
}
type Database struct {
	DatabaseName string `json:"databaseName"`
	HostName     string `json:"hostName"`
	Port         int    `json:"port"`
}

func ReadConfig(path string) (c Config, err error) {
	configFile, err := os.Open(path)
	if err != nil {
		return c, fmt.Errorf("cannot open this file: %v", err)
	}
	defer configFile.Close()

	byteValue, err1 := ioutil.ReadAll(configFile)
	if err1 != nil {
		fmt.Println("error while converting to byte")
		return c, err1
	}
	json.Unmarshal(byteValue, &c)
	return c, nil
}
