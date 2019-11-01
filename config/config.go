/*
 * Copyright (c) 2019.
 * Author:Alex	xiaoshenlong1910@gmail.com

 * code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
              ┏┓      ┏┓
            ┏┛┻━━━┛┻┓
            ┃      ☃      ┃
            ┃  ┳┛  ┗┳  ┃
            ┃      ┻      ┃
            ┗━┓      ┏━┛
                ┃      ┗━━━┓
                ┃  神兽保佑    ┣┓
                ┃　永无BUG！   ┏┛
                ┗┓┓┏━┳┓┏┛
                  ┃┫┫  ┃┫┫
                  ┗┻┛  ┗┻┛
*/

package config

import (
	"os"

	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
	AppName string
	RunMode string
	App
	Server
	DB
	Log
	// ...
}

type App struct {
	JwtSecret string
}

type Server struct {
	Port         int
}

type DB struct {
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type Log struct {
	Type  string
	Level string
}

func InitConfig() (err error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	pwd, _ := os.Getwd()
	v.AddConfigPath(pwd)
	if err = v.ReadInConfig(); err != nil {
		return
	}
	if err = v.Unmarshal(&Conf); err != nil {
		return
	}
	return
}
