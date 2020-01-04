package util

import (
	"encoding/json"
	"io/ioutil"
)

// Config プロパティ
type Config struct {
	Server ServerConfig `json:"server"`
	Db     DBConfig     `json:"db"`
}

// ServerConfig サーバ
type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// DBConfig DB緒元
type DBConfig struct {
	Address string `json:"address"`
	Cluster string `json:"cluster"`
	User    string `json:"user"`
	Pass    string `json:"pass"`
}

// GetConfig ゲッタ
func GetConfig() Config {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)

	return config
}
