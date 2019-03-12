package mssql

import (
	"fmt"
	"database/sql"
	"github.com/astaxie/beego/config"
	_"github.com/denisenkom/go-mssqldb"
)

var (
	MssqlDb *sql.DB
	err error
)

func init() {
	cfg, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
		return
	}
	server := cfg.String("mssql::host")
	server_port,_ := cfg.Int("mssql::port")
	mssql_db := cfg.String("mssql::dbname")
	ms_user := cfg.String("mssql::user")
	ms_pwd := cfg.String("mssql::password")
	connString := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%d;encrypt=disable", server, mssql_db, ms_user, ms_pwd, server_port)
	fmt.Printf(" connString:%s\n", connString)
	MssqlDb, err = sql.Open("mssql", connString)
	if err != nil {
		panic(err)
	}
}