package configs

import (
	"os"

	"github.com/spf13/viper"
)

type Service struct {
	Port       string `yaml:"port"`
	Version    string `yaml:"version"`
	Jwt_secret string `yaml:"jwt_secret"`
}

type Database struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Db_name  string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Db_name  string `yaml:"dbname"`
	Network  string `yaml:"network"`
}

type Node struct {
	NodeID int64 `yaml:"node_id"`
}

type Configs struct {
	Service  *Service  `yaml:"service"`
	Database *Database `yaml:"database"`
	Redis    *Redis    `yaml:"redis"`
	Node     *Node     `yaml:"node"`
}

var Config *Configs

func InitConfig() {
	workdir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workdir + "/configs")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		panic(err)
	}
}
