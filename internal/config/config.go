package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type Config struct {
	Env      string         `yaml:"env" env-default:"local"`
	GRPC     GRPCConfig     `yaml:"grpc"`
	Postgres PostgresConfig `yaml:"postgres"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}
type PostgresConfig struct {
	Url string `yaml:"url"`
}

func MustLoad() *Config {

	path := fetchConfigPath()

	if path == "" {
		panic("config path is empty")
	}

	return MustLoadPath(path)
}

func MustLoadPath(path string) *Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file doesn't exist" + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	res = os.Getenv("CONFIG_PATH")

	if res == "" {
		if err := godotenv.Load(".env"); err != nil {
			panic("failed to load .env file: " + err.Error())
		}
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
