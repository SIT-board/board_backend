package global

import (
	"sync"
)

type ServerConfig struct {
	ListenIp   string `yaml:"listen-ip"`
	ListenPort int    `yaml:"listen-port"`
}

type ImageConfig struct {
	RawDirectory       string `yaml:"raw-directory"`
	ThumbnailDirectory string `yaml:"thumbnail-directory"`
}

type Config struct {
	Server ServerConfig `yaml:"server"`
	Image  ImageConfig  `yaml:"image"`
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
		Image: ImageConfig{
			RawDirectory:       "images/raw",
			ThumbnailDirectory: "images/thumbnail",
		},
	}
}
