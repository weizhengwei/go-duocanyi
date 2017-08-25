package main

import (
	"database/sql"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
)

//医生数据
type JsonDoctor struct {
	Data []JsonDoctorItem
}

type JsonDoctorItem struct {
	PID string
	CREDENTIAL_CODE string
	NAME string
	ORG_CODE string
	AGENCY_NAME string
	SEX string
	BIRTHDAY string
	CENSUS_REGISTER_ADDRESS string	
	MOBILE string
	EMAIL string
	PASSWDWRD string
	RECORDWAY string
	MACHINE_ID string
	FLAG string
	SCDATE string
}

func UploadDoctor(s string, db *sql.DB) {
	var jd JsonDoctor
	err := json.Unmarshal([]byte(s), &jd)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jd)
}