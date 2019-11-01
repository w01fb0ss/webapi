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

package main

import (
	"fmt"
	"os"

	"github.com/w01fb0ss/common/config"
	"github.com/w01fb0ss/common/models"
	"github.com/w01fb0ss/common/routers"
)

func init() {
	if err := config.InitConfig(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if err := models.InitDB(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func main() {
	router := routers.InitRouter()
	router.Run(fmt.Sprintf(":%d", config.Conf.Server.Port))
}
