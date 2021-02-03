package config

import "fmt"

var Config Configure

type AppConfig struct {
	Name string
	Port int
}

type MysqlConfig struct {
	Host             string
	UserName         string
	Password         string
	DataBase         string
	Port             int
	MaxLifeTime      int
	MaxIdleConns     int
	ShareDbNumber    int
	ShareTableNumber int
}

type RedisConfig struct {
	Host     string
	Port     int
	UserName string
	Password string
	TimeOut  int64
}

/* 公共配置end */

type Configure struct {
	App    AppConfig
	Master MysqlConfig
	Slave  MysqlConfig
}


func LoadTomlConfig() error {

	myViper := getViper("config", "toml")
	if err := myViper.ReadInConfig(); err != nil {
		//	logger.ErrorWithoutContext("load config err:%s", err.Error())
		return err
	}

	if err := myViper.Unmarshal(&Config); err != nil {
		//	logger.ErrorWithoutContext("load config err:%s", err.Error())
		return err
	}

	return nil
}


func GetPort() string {
	return fmt.Sprintf(":%d", Config.App.Port)
}
