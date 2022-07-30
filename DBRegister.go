package gosql

import (
	"database/sql"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 数据库Source连接
var dbOptionMap map[string]string
var dbPoolObjectMap map[string]*sql.DB
var redisPoolObjectMap map[string]redis.Conn

// DBRegister  数据库注册方法
func DBRegister(DrvName string, Server string, DBName string, User string, Password string, Port int, NickName string) {
	switch DrvName {
	case "mssql":
		dbOptionMap[NickName] = fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;port=%d;encrypt=disable", Server, User, Password, DBName, Port)
		db, _ := sql.Open("mssql", dbOptionMap[NickName])
		dbPoolObjectMap[NickName] = db
	case "mysql":
		dbOptionMap[NickName] = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", User, Password, Server, Port, DBName)
		db, _ := sql.Open("mysql", dbOptionMap[NickName])
		dbPoolObjectMap[NickName] = db
	case "redis":
		address := fmt.Sprintf("%s:%d", Server, Port)
		c, err := redis.Dial("tcp", address)
		if err != nil {
			fmt.Println("Connect to redis error", err)
			return
		} else {
			redisPoolObjectMap[NickName] = c
		}
		//case "oracle":
		//	dbOptionMap[NickName] = fmt.Sprintf(`user="%s" password = "%s" connectString="%s:%d/%s"`, User, Password, Server, Port, DBName)
		//	db, _ := sql.Open("godror", dbOptionMap[NickName])
		//	dbPoolObjectMap[NickName] = db
	}

}
