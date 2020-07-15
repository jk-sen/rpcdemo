package connector

import (
    "api.com/rpc/package/seting"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

var Db *gorm.DB
var err error

// 数据库连接
func MysqlConnect() {

    args := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", seting.MysqlSet.User, seting.MysqlSet.Password, seting.MysqlSet.Host, seting.MysqlSet.Port, seting.MysqlSet.Name)
    Db, err = gorm.Open("mysql", args)

    if err != nil {
        panic(err)
    }

    Db.DB().SetMaxIdleConns(10)
    Db.DB().SetMaxOpenConns(100)
    Db.SingularTable(true)
}
