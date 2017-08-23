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

	var ID, AREA_CODE string
	err = db.QueryRow("select ID$ as OID,AREA_CODE from tb_organization where ORG_CODE=? and ORG_NAME=?", 
		registerjson.Data.ORG_CODE, registerjson.Data.ORG_NAME).Scan(&ID, &AREA_CODE)
	if err != nil {//机构代码和名称不存在
		fmt.Println(err)
		return
	}
	fmt.Println(ID, AREA_CODE)
}