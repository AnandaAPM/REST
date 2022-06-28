package models

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/libra9z/orm"
)

func init() {
	// if beego.BConfig.RunMode == "dev" {
	orm.RegisterDataBase("default", "sqlserver", "sqlserver://mobile:m0b1l31234@10.10.0.23:1433")
	// } else if beego.BConfig.RunMode == "prod" {
	// 	orm.RegisterDataBase("default", "postgres", "postgresql://kong:kong@10.244.1.204/postgres?sslmode=disable")
	// }
	orm.Debug = true
}
