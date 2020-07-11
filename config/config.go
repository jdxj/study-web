package config

import (
	"encoding/json"
	"os"
)

const configPath = "config.json"

var (
	config *Config
)

func init() {
	file, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cfg := &Config{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		panic(err)
	}
	config = cfg
}

type Config struct {
	DB     *DB     `json:"db"`
	Server *Server `json:"server"`
}

type DB struct {
	Host string `json:"host"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
}

type Server struct {
	Addr string `json:"addr"`
}

func GetDB() *DB {
	return config.DB
}

func GetServer() *Server {
	return config.Server
}
