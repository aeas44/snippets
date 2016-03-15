package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	DataDir string 	`json:"dataDir"`
}

func main() {
	// fmt.Println("Yo")
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)

	fmt.Println(config.DataDir) // "/tmp"
}

