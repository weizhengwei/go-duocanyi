package main

import (
	"database/sql"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
)

//健康报告数据
type JsonReport struct {
	Data []JsonReportItem
}

type JsonReportItem struct {
	PID string
	BGMC string
	YSBH string
	YSMC string
	USERBH string
	USERMC string
	ORG_CODE string
	ORG_NAME string	
	YHXD string
	YHXY01 string
	YHXY string
	YHXT string
	YHNY string
	YHBMI string
	YHNS string
	YHDGC string
	YHXHDB string
	YHXZ string
	JKPJ string
	JKZD string
}

func UploadReport(s string, db *sql.DB) {
	var jr JsonReport
	err := json.Unmarshal([]byte(s), &jr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jr)
}