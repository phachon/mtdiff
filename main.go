package main

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"log"
	"strings"
)

// mtdiff --master="dsn" --slave="dsn" --table="table"
// dsn = [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
// root:admin@127.0.0.1:3306/test
var (
	defaultDsn = "[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]"
	master = flag.String("master", "", "please input master mysql dsn: "+defaultDsn)
	slave = flag.String("slave", "", "please input slave mysql dsn: "+defaultDsn)
	table = flag.String("table", "*", "please input table name")
)

var (
	masterHandle *mysqlHandle
	slaveHandle *mysqlHandle
)

func init()  {
	flag.Parse()
}

func main() {

	var err error

	masterHandle, err = NewMysqlHandle(*master)
	if err != nil {
		printError(err.Error())
	}
	defer masterHandle.Close()
	slaveHandle, err = NewMysqlHandle(*slave)
	if err != nil {
		printError(err.Error())
	}
	defer slaveHandle.Close()

	var handleTables = []string{}
	masterTables, err := masterHandle.ShowTables()
	if err != nil {
		printError(err.Error())
	}

	// match table
	// '*' all
	if *table == "*" {
		handleTables = masterTables
	}else if strings.Index(*table, "*") == 0 {
		// '*test' match suffix = test
		suffix := strings.TrimLeft(*table, "*")
		for _, masterTable := range masterTables {
			if strings.HasSuffix(masterTable, suffix) {
				handleTables = append(handleTables, masterTable)
			}
		}

	}else if strings.Index(*table, "*") == (len(*table) - 1) {
		// 'test*' match prefix = test
		prefix := strings.TrimRight(*table, "*")
		for _, masterTable := range masterTables {
			if strings.HasPrefix(masterTable, prefix) {
				handleTables = append(handleTables, masterTable)
			}
		}
	}else {
		// 'test' match test
		for _, masterTable := range masterTables {
			if masterTable == *table {
				handleTables = append(handleTables, masterTable)
				break
			}
		}
	}

	slaveTables, err := slaveHandle.ShowTables()
	if err != nil {
		printError(err.Error())
	}
	slaveTablesMap := map[string]string{}
	for _, slaveTable := range slaveTables {
		slaveTablesMap[slaveTable] = slaveTable
	}

	for _, handleTable := range handleTables {
		tableName, ok := slaveTablesMap[handleTable]
		if !ok {
			// create table
			createSql, err := masterHandle.ShowCreateTable(tableName)
			if err != nil {
				printError(err.Error())
			}
			err = slaveHandle.CreateTable(tableName, createSql)
			if err != nil {
				printError(err.Error())
			}
		}else {
			// todo check table column

		}
	}

}

func printError(msg string)  {
	log.Println(msg)
	os.Exit(100)
}