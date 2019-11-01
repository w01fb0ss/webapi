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

package v1

import (
	"github.com/gin-gonic/gin"
)

func GetTests(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "test",
	})
}

func GetTest(c *gin.Context)  {

}

func AddTest(c *gin.Context)  {

}

func EditTest(c *gin.Context)  {

}

func DeleteTest(c *gin.Context)  {

}