package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port  string
	Mysql Mysql
}

type Mysql struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func Load(path string) Config {

	godotenv.Load(path + "/.env")

	conf := viper.New()

	conf.AutomaticEnv()
    
	cfg := Config{
		Port: conf.GetString("PORT"),
		Mysql: Mysql{
			Host:     conf.GetString("MYSQL_HOST"),
			Port:     conf.GetString("MYSQL_PORT"),
			User:     conf.GetString("MYSQL_USER"),
			Password: conf.GetString("MYSQL_PASSWORD"),
			DBName:   conf.GetString("MYSQL_DB"),
		},
	}

	return cfg
}