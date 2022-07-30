package gosql

import (
	"context"
	"database/sql"
	"fmt"
)

func Exec(con *sql.Conn, params string, args ...interface{}) bool {
	res, err := con.ExecContext(context.Background(), params, args...)
	if err != nil {
		fmt.Println("Error time is：" + timeStr())
		fmt.Println("Error SQL is：" + params)
		fmt.Println("Database return Error is：" + err.Error())
		return false
	}
	if res != nil {
		RowsAffected, Err := res.RowsAffected()
		if Err == nil {
			fmt.Println("Handle：" + params)
			rows := fmt.Sprintln(RowsAffected, "RowsAffected")
			fmt.Println(rows)
		}
	}

	return true
}
