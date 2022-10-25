package gosql

import (
	"database/sql"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

var dbOptionMap map[string]string
var dbPoolObjectMap map[string]*sql.DB
var redisPoolObjectMap map[string]*redis.Pool

/*DBRegister
 * Before you get connection,you need regist the database first.
 * 在你获取连接之前，你需要先对数据库进行注册。
 */
func DBRegister(DrvName string, Server string, DBName string, User string, Password string, Port int, NickName string) {
	switch DrvName {
	case "mssql":
		dbOptionMap[NickName] = fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;port=%d;encrypt=disable", Server, User, Password, DBName, Port)
		db, _ := sql.Open("mssql", dbOptionMap[NickName])
		dbPoolObjectMap[NickName] = db
		break
	case "mysql":
		dbOptionMap[NickName] = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", User, Password, Server, Port, DBName)
		db, _ := sql.Open("mysql", dbOptionMap[NickName])
		dbPoolObjectMap[NickName] = db
		break
	case "redis":
		address := fmt.Sprintf("%s:%d", Server, Port)
		db := &redis.Pool{
			MaxIdle:     10,
			MaxActive:   0,
			IdleTimeout: 100,
			Dial: func() (redis.Conn, error) {
				if User == "" && Password == "" {
					return redis.Dial("tcp", address)
				}
				if User == "" && Password != "" {
					return redis.Dial("tcp", address, redis.DialPassword(Password))
				}
				return redis.Dial("tcp", address, redis.DialUsername(User), redis.DialPassword(Password))
			},
		}
		redisPoolObjectMap[NickName] = db
		break
	case "oracle":
		dbOptionMap[NickName] = fmt.Sprintf(`user="%s" password = "%s" connectString="%s:%d/%s"`, User, Password, Server, Port, DBName)
		db, _ := sql.Open("godror", dbOptionMap[NickName])
		dbPoolObjectMap[NickName] = db
		break
	default:
		log.Fatal(fmt.Sprintf("unknown database type %s", DrvName))
	}

}
