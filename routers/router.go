/*
 * Copyright (c) 2019.
 * Author:Alex	xiaoshenlong1910@gmail.com

   .--,       .--,
 ( (  \.---./  ) )
  '.__/o   o\__.'
     {=  ^  =}
      >  -  <
     /       \
    //       \\
   //|   .   |\\
   "'\       /'"_.-~^`'-.
      \  _  /--'         `
    ___)( )(___
   (((__) (__)))    高山仰止,景行行止.虽不能至,心向往之。

*/

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/w01fb0ss/common/middleware/jwt"
	"github.com/w01fb0ss/common/middleware/log"
	"github.com/w01fb0ss/common/routers/api"
	v1 "github.com/w01fb0ss/common/routers/api/v1"
)

func InitRouter() *gin.Engine {
	// gin.SetMode(config.Conf.RunMode)
	r := gin.New()
	l := logrus.New()
	r.Use(log.Logger(l))
	r.Use(gin.Recovery())
	r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tests", v1.GetTests)
		apiv1.GET("/tests/:id", v1.GetTest)
		apiv1.POST("/tests", v1.AddTest)
		apiv1.PUT("/tests/:id", v1.EditTest)
		apiv1.DELETE("/tests/:id", v1.DeleteTest)
	}
	return r
}
