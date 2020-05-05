package dblib


import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	//mysql
	MysqlString string = "root" + ":" + "GaoPeng@123" + "@tcp(" + "192.168.2.5" + ":" + "3306" + ")/" + "yyw-shop"
	//mysql 连接池
	Db *sql.DB
	//err
	err error
)

func init()  {
	Db, err = sql.Open("mysql",MysqlString)
	check(err)
}


func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
