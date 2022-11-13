package global

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
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

func GetConfig(filePath string) *Config {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	cfg := &Config{}
	if err = yaml.Unmarshal(file, cfg); err != nil {
		log.Fatal(err)
	}
	log.Println(cfg)
	return cfg
}

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
