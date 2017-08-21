package main

import (
	"database/sql"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
)

type RegisterJsonData struct {
	PID string
	MACHINE_ID string
	TYPE string
	ORG_NAME string
	ORG_CODE string
	CONTACT string
	CONTACT_PHONE string
	CREATETIME string
}

type RegisterJson struct {
	Data RegisterJsonData
}

func DealRegister(s string, db *sql.DB) {
	var registerjson RegisterJson
	err := json.Unmarshal([]byte(s), &registerjson)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(registerjson)
}