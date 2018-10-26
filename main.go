package main

import (
	"flag"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"log"
)

// mtdiff --master="dsn" --slave="dsn" --table="table"
// dsn = [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
// root:admin@127.0.0.1:3306/test
var (
	defaultDsn = "[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]"
	master = flag.String("master", "", "please input master mysql dsn: "+defaultDsn)
	slave = flag.String("slave", "", "please input slave mysql dsn: "+defaultDsn)
	table = flag.String("table", "", "please input table name")
)

var (
	masterConn *sql.DB
	slaveConn *sql.DB
)

func init()  {
	flag.Parse()
}

func main() {

	var err error

	// master mysql conn
	masterConn, err = sql.Open("mysql", *master)
	if err != nil {
		printError("mysql master conn failed, "+err.Error())
	}
	err = masterConn.Ping()
	if err != nil {
		printError("mysql master conn failed, "+err.Error())
	}

	// slave mysql conn
	slaveConn, err = sql.Open("mysql", *slave)
	if err != nil {
		printError("mysql slave conn failed, "+err.Error())
	}
	err = slaveConn.Ping()
	if err != nil {
		printError("mysql slave conn failed, "+err.Error())
	}

	// list table
	//row, err := masterConn.Query("SHOW TABLES;")
	//if
	//row, err := masterConn.Query("select * from test.test_account;")
	//if err != nil {
	//	printError(err.Error())
	//}
	//fmt.Println(row.Columns())

}

func printError(msg string)  {
	log.Println(msg)
	os.Exit(100)
}