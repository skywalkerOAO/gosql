package gosql

import (
	"context"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
)

// 初始化连接组

func init() {
	dbOptionMap = make(map[string]string)
	dbPoolObjectMap = make(map[string]*sql.DB)
	redisPoolObjectMap = make(map[string]redis.Conn)
}

// GetDBCon 获取连接字符串
func GetDBCon(NickName string) (*sql.Conn, error) {
	return dbPoolObjectMap[NickName].Conn(context.Background())
}
func GetRedisCon(NickName string) redis.Conn {
	return redisPoolObjectMap[NickName]
}
