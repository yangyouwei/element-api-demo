package dblib


import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	//mysql
	MysqlString string = "root" + ":" + "Yang@1008" + "@tcp(" + "192.168.1.202" + ":" + "3306" + ")/" + "shop"
	//mysql 连接池
	Db *sql.DB
	//err
	err error
)

func init()  {
	Db, err = sql.Open("mysql",MysqlString)
	check(err)

	err = Db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}


func check(err error) {
	if err != nil {
		fmt.Println(err)
		log.Panic(err)
	}
}
