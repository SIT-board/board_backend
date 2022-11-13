package global

import (
	"sync"
)

type ServerConfig struct {
	ListenIp   string `yaml:"listen-ip"`
	ListenPort int    `yaml:"listen-port"`
}

type FileConfig struct {
	Directory string `yaml:"directory"`
	Host      string `yaml:"host"`
}

type Config struct {
	Server ServerConfig `yaml:"server"`
	File   FileConfig   `yaml:"file"`
}

var (
	config     Config
	configOnce sync.Once
)

func GetDefaultConfig() *Config {
	return &Config{
		Server: ServerConfig{
			ListenIp:   "127.0.0.1",
			ListenPort: 8080,
		},
		File: FileConfig{
			Directory: "files",
			Host:      "http://localhost:8080",
		},
	}
}
