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
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

func MustLoad() *Config {

	path := fetchConfigPath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file doesn't exist" + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DBName = fetchPostgresEnv()

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

func fetchPostgresEnv() (user, password, dbname string) {

	if err := godotenv.Load(".env"); err != nil {
		panic("failed to load .env file: " + err.Error())
	}

	user = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname = os.Getenv("POSTGRES_DB")

	return user, password, dbname
}
