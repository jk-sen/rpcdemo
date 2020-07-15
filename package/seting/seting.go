package seting

import "os"
import "gopkg.in/ini.v1"

// 初始化连接器中所有的参数
var iniFile *ini.File

// MySQL 设置
type Mysql struct {
    Type     string
    User     string
    Password string
    Host     string
    Name     string
    Port     int
}

var MysqlSet = &Mysql{}

// redis 设置
type Redis struct {
    Host string
    Port string
    Auth string
}

var RedisSet = &Redis{}

func Init() {
    dir, _ := os.Getwd()

    var err error
    iniFile, err = ini.Load(dir + "/app.ini")

    if err != nil {
        panic(err)
    }
    LoadMysqlConf()
    LoadRedisConf()

}

func LoadRedisConf() {
    redisSection := iniFile.Section("redis")

    RedisSet.Host = redisSection.Key("HOST").String()
    RedisSet.Port = redisSection.Key("PORT").String()
    RedisSet.Auth = redisSection.Key("PORT").String()
}

func LoadMysqlConf() {
    mysqlSection := iniFile.Section("mysql")

    port, err := mysqlSection.Key("PORT").Int()
    if err != nil {
        panic("mysql端口转化错误")
    }

    MysqlSet.Port = port
    MysqlSet.Host = mysqlSection.Key("HOST").String()
    MysqlSet.User = mysqlSection.Key("USER").String()
    MysqlSet.Password = mysqlSection.Key("PASSWORD").String()
    MysqlSet.Name = mysqlSection.Key("DB").String()

}
