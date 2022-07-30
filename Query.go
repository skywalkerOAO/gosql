package gosql

import (
	"context"
	"database/sql"
	"fmt"
)

// Query 返回查询所有行
func Query(con *sql.Conn, params string, args ...interface{}) []map[string]interface{} {
	row, err := con.QueryContext(context.Background(), params, args...)
	if err != nil {
		fmt.Println("Error time is：" + timeStr())
		fmt.Println("Error SQL is：" + params)
		fmt.Println("Database return Error is：" + err.Error())
		return nil
	}
	columns, err := row.Columns()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	tableData := make([]map[string]interface{}, 0)
	if err != nil {
		return tableData
	}
	count := len(columns)
	if count == 0 {
		fmt.Println(tableData)
		return tableData
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
	return tableData
}
