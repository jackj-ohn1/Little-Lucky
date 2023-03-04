package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var Conf conf

type conf struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Param    string `yaml:"param"`
		Database string `yaml:"database"`
	} `yaml:database`
	Email struct {
		Host      string `yaml:"host"`
		Port      int    `yaml:"port"`
		Sender    string `yaml:"sender"`
		SecretKey string `yaml:"secretKey"`
	}
	Token struct {
		Issuer string `yaml:"issuer"`
	} `yaml:"token"`
}

type Config struct {
	name string
}

func (cfg *Config) init() {
	if cfg.name != "" {
		viper.SetConfigFile(cfg.name)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("读取配置文件失败:", err)
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		log.Fatal("配置文件初始化失败:", err)
	}
}

func (cfg *Config) watch() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Fatalf("Config file changed: %s", in.Name)
	})
	return
}

func Run(name string) {
	cfg := &Config{name}
	cfg.init()
	cfg.watch()
}
