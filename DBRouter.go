package gosql

import (
	"context"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
)

func init() {
	dbOptionMap = make(map[string]string)
	dbPoolObjectMap = make(map[string]*sql.DB)
	redisPoolObjectMap = make(map[string]*redis.Pool)
}

/*GetDBCon
  get database connect (from )
  获取数据库连接
*/
func GetDBCon(NickName string) (*sql.Conn, error) {
	return dbPoolObjectMap[NickName].Conn(context.Background())
}

/*GetDB
  get DB string
  获取数据库实例
*/
func GetDB(NickName string) *sql.DB {
	return dbPoolObjectMap[NickName]
}

/*GetRedisCon
  get redis connect
  获取Redis连接字符串
*/
func GetRedisCon(NickName string) redis.Conn {
	return redisPoolObjectMap[NickName].Get()
}
