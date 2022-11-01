package gosql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func Exec(con *sql.Conn, params string, args ...interface{}) (bool, error) {
	_, err := con.ExecContext(context.Background(), params, args...)
	if err != nil {
		dberr := fmt.Sprintf("Database return error is：%s \n", err.Error())
		errtime := fmt.Sprintf("Error time is：%s \n", timeStr())
		errsql := fmt.Sprintf("Error SQL is：%s \n", params)
		e := errors.New(dberr + errtime + errsql)
		return false, e
	}
	//if res != nil {
	//	RowsAffected, Err := res.RowsAffected()
	//	if Err == nil {
	//		fmt.Println("Handle：" + params)
	//		rows := fmt.Sprintln(RowsAffected, "RowsAffected")
	//		fmt.Println(rows)
	//	}
	//}
	return true, nil
}

func OpenTransaction(DB *sql.DB) (*sql.Tx, error) {
	return DB.Begin()
}
func TExec(tx *sql.Tx, params string, args ...interface{}) (int, error) {
	res, err := tx.Exec(params, args...)
	if err != nil {
		tx.Rollback()
		fmt.Printf("FAILED %s ,ERROR：%v\n", params, err)
		return 0, err
	}
	RowsAffected, _ := res.RowsAffected()
	fmt.Printf("SUCCESS %v ROWS \n", RowsAffected)
	return int(RowsAffected), err
}
func SubmitTransaction(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Printf("SUBMIT FAILED , Transaction RollBacked ,ERROR：%v\n", err)
		return err
	}
	return nil
}
