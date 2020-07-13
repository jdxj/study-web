package config

import (
	"encoding/json"
	"fmt"
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
	Logger *Logger `json:"logger"`
}

// DB config
type DB struct {
	Host string `json:"host"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
}

func GetDB() *DB {
	return config.DB
}

// Server config
type Server struct {
	Addr string `json:"addr"`
}

func GetServer() *Server {
	return config.Server
}

// Logger config
type Logger struct {
	FileName string `json:"filename"`
	Level    int    `json:"level"`
	MaxLines int    `json:"maxlines"`
	MaxSize  int    `json:"maxsize"`
	Daily    bool   `json:"daily"`
	MaxDays  int    `json:"maxdays"`
	Color    bool   `json:"color"`
}

func (l *Logger) String() string {
	data, _ := json.Marshal(l)
	return fmt.Sprintf("%s", data)
}

func GetLogger() *Logger {
	return config.Logger
}
