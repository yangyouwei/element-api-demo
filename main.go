package main

import (
	db "github.com/yangyouwei/element-api-demo/dblib"
	. "github.com/yangyouwei/element-api-demo/router"
)

func main() {
	defer db.Db.Close()
	router := InitRouter()
	router.Run(":8080")
}
