package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string `yaml:"address" env-default:"localhost:8080"`
}

// struct embedding
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	Storagepath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	HttpServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flg := flag.String("config", "", "Path to configuration file")
		flag.Parse()      // Parse command-line flags
		configPath = *flg // Override with command-line flag if provided
		if configPath == "" {
			log.Fatal("Config Path is required")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist at path: %s", configPath)
	}

	var cfg Config
	er := cleanenv.ReadConfig(configPath, &cfg) // reads the configuration file (like local.yaml) from the given path and loads its values into the cfg struct — mapping fields based on their yaml tags.
	if er != nil {
		log.Fatalf("failed to read config: %v", er.Error())
	}
	fmt.Println("Configuration loaded from", configPath)

	return &cfg

}
