package main

import (
	"database/sql"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
)

//医生数据
type JsonDocument struct {
	Data []JsonDocumentItem
}

type JsonDocumentItem struct {
	PID string
	PERSON_ID string
	CZLX string
	JOBTYPE string
	YLFYPAY string
	LXRMC string
	LXRTEL string
	WORK_PLACE_NAME string
	BLOODTYPE string
	RH string
	MARRIAGE_STATUS_CODE string
	DEGREE string
	RELIGION string
	GUARDIAN string
	YWGMS string
	CJQK string
	DISEASE string
	YCBS string
	BLS string
	SHHJ string
	SCDATE string
	CJYSBM string
	XGYSBM string
	ORG_CODE string
	ORG_NAME string
}

func UploadDocument(s string, db *sql.DB) {
	var jd JsonDocument
	err := json.Unmarshal([]byte(s), &jd)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jd)
}