package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var conf *Config

func init() {
	confPath := "/Users/zhangbicheng/PycharmProjects/msc/msc-backend/conf/conf.toml"
	if _, err := toml.DecodeFile(confPath, &conf); err != nil {
		fmt.Println(err)
	}
}

func GetConfig() *Config {
	return conf
}

// Config ： 配置信息
type Config struct {
	RdConf      redisConfig   `toml:"redis"`
	PgConf      pgConfig      `toml:"postgres"`
	CrawlerConf crawlerConfig `toml:"crawler"`
}

type crawlerConfig struct {
	Host string `toml:"host"`
}

// RedisConfig : redis配置信息
type redisConfig struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

// PgConfig : Postgres配置信息
type pgConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"username"`
	DBName   string `toml:"dbname"`
	Password string `toml:"password"`
}
