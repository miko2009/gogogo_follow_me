package config

import "fmt"

var JsonConfig JsonConfigure

type JsonConfigure struct {
	ErrorMsg  ErrorMsg `mapstructure:"error_msg"`
}


type ErrorMsg struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}
func LoadJsonConfig() error {

	myViper := getViper("message", "json")
	if err := myViper.ReadInConfig(); err != nil {
		return err
	}

	fmt.Println(myViper.GetString("error_msg.msg"))
	if err := myViper.Unmarshal(&JsonConfig); err != nil {
		return err
	}
	fmt.Println(JsonConfig.ErrorMsg)
	return nil
}
