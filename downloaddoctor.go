package main

import (
	"database/sql"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
)

//医生数据
type ReqDoctor struct {
	Type string
	ORG_CODE string
	MACHINE_ID string
	TIME string
}

func DownloadDoctor(s string, db *sql.DB) {
	fmt.Println("DownloadDoctor")
	var rd ReqDoctor
	err := json.Unmarshal([]byte(s), &rd)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rd)
	if rd.Type == "1" {
		if len(rd.TIME) == 0 {

		}else{

		}
	}else if rd.Type == "0" {
		// rows, err := db.Query("SELECT NAME,SEX,ORG_CODE,MOBILE,BIRTHDAY,CREDENTIAL_CODE,PASSWORD,EMAIL,FLAG,CENSUS_REGISTER_ADDRESS,UPDATETIME$ as UPDATETIME,SCDATE,SOURCEDATA,RECORDWAY,FILENAME FROM tb_medical_technicians WHERE CATEGORY='1' and ORG_CODE= ? and FLAG='1'", "2016041301")
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// defer rows.Close()
		// for rows.Next() {
		// 	var NAME,SEX,ORG_CODE,MOBILE,BIRTHDAY,CREDENTIAL_CODE,PASSWORD,EMAIL,FLAG,CENSUS_REGISTER_ADDRESS,UPDATETIME,SCDATE,SOURCEDATA,RECORDWAY,FILENAME string
		// 	if err := rows.Scan(&NAME,&SEX,&ORG_CODE,&MOBILE,&BIRTHDAY,&CREDENTIAL_CODE,&PASSWORD,&EMAIL,&FLAG,&CENSUS_REGISTER_ADDRESS,&UPDATETIME,&SCDATE,&SOURCEDATA,&RECORDWAY,&FILENAME); err != nil {
		// 		fmt.Println(err)
		// 		return
		// 	}
		// 	fmt.Println(NAME)
		// }
	}
}