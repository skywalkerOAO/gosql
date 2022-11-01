package gosql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// Query 返回查询所有行
func Query(con *sql.Conn, params string, args ...interface{}) ([]map[string]interface{}, error) {
	row, err := con.QueryContext(context.Background(), params, args...)
	if err != nil {

		dberr := fmt.Sprintf("Database return error is：%s \n", err.Error())
		errtime := fmt.Sprintf("Error time is：%s \n", timeStr())
		errsql := fmt.Sprintf("Error SQL is：%s \n", params)
		e := errors.New(dberr + errtime + errsql)
		return nil, e
	}
	columns, err := row.Columns()
	if err != nil {
		return nil, err
	}
	tableData := make([]map[string]interface{}, 0)
	if err != nil {
		return tableData, nil
	}
	count := len(columns)
	if count == 0 {
		return tableData, nil
	}
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for row.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		row.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	return tableData, nil
}
