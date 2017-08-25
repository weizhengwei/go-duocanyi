package main

import (
	"database/sql"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
)

//医生数据
type JsonPerson struct {
	Data []JsonPersonItem
}

type JsonPersonItem struct {
	PID string
	GLOBAL_ID string
	IDENTITY_CARDS string
	HCARDNO string
	NAME string
	GENDER string
	BIRTHDAY string
	ADDRESS string	
	TELEPHONE_NUMBER string
	NATION string
	EMAIL string
	RECORDWAY string
	SCDATE string
	CJYSBM string
	XGYSBM string
	FLAG string
	YSBM string
	MACHINE_ID string
	ORG_CODE string
	ORG_NAME string
}

func UploadPerson(s string, db *sql.DB) {
	var jp JsonPerson
	err := json.Unmarshal([]byte(s), &jp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jp)
}