package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Token    string `json:"token"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ReadConfig() Config {
	file, err := os.OpenFile("./config.json", os.O_CREATE, 0644)
	//fmt.Println(file)
	if err != nil && !os.IsExist(err) {
		fmt.Println("Generating config.json")
		Data := Config{
			Token:    "YourToken",
			Email:    "YourEmail",
			Password: "YourPassword",
		}
		R, err := json.Marshal(Data)
		if err != nil {
			fmt.Println(err)
		}
		b, _ := prettyprint(R)
		os.WriteFile("config.json", []byte(b), 0666) //Writing file
		os.Exit(0)
	}
	var config Config
	body, _ := ioutil.ReadAll(file)
	if err := json.Unmarshal(body, &config); err != nil {
		fmt.Println(err)
		return Config{}
	}
	return config
}

func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}
