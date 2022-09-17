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

func OpenTransaction(DB *sql.DB) (*sql.Tx, error) {
	return DB.Begin()
}
func TExec(tx *sql.Tx, params string, args ...interface{}) (int, error) {
	res, err := tx.Exec(params, args...)
	if err != nil {
		tx.Rollback()
		fmt.Printf("执行%s失败，错误：%v\n", params, err)
		return 0, err
	}
	RowsAffected, _ := res.RowsAffected()
	fmt.Printf("执行成功：%v行\n", RowsAffected)
	return int(RowsAffected), err
}
func SubmitTransaction(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Printf("提交失败，事务已回滚，错误：%v\n", err)
		return err
	}
	return nil
}
