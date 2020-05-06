package main

import (
	db "github.com/yangyouwei/element-api-demo/dblib"
	. "github.com/yangyouwei/element-api-demo/router"
)

func main() {
	defer db.Db.Close()
	router := InitRouter()
	//router.Use(gin.Logger())
	//router.Use(utils.CrossDomain())
	//router.Use(utils.Validate())
	router.Run(":8080")
}

