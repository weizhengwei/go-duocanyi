package main

import (
	"database/sql"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
)

type HeartChartInfo struct {
	Data HeartChartInfoItem
}

type HeartChartInfoItem struct {
	PID string
	Type int
	YHBM string
	YHMC string
	YSBM string
	YSMC string
	JCRQ string
	SJZT string
	MACHINE_ID string
	ORG_CODE string
	ORG_NAME string
	Bullfile_id string
	Xml_bullfile_id string
}

type FatalChartInfo struct {
	PID string
	Type int
	YHBM string
	YHMC string
	YSBM string
	YSMC string
	JCRQ string
	SJZT string
	MACHINE_ID string
	ORG_CODE string
	ORG_NAME string
	AverateFhr string
	MonitorStartTime string
	MonitorEndTime string
	CheckGestationalAge string
}

func UploadHeartChartFileInfo(s string, db *sql.DB) {
	fmt.Println("UploadHeartChartFileInfo")
	var hci HeartChartInfo
	err := json.Unmarshal([]byte(s), &hci)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(hci)
	var empi, doctorempi string
	row := db.QueryRow("SELECT EMPI FROM mpi_personbasics WHERE IDENTITY_CARDS = ?", hci.Data.YHBM)
	err = row.Scan(&empi)
	fmt.Println(empi)

	row = db.QueryRow("SELECT DOCTOR_EMPI FROM tb_medical_technicians WHERE CREDENTIAL_CODE = ?", hci.Data.YSBM)
	err = row.Scan(&doctorempi)
	fmt.Println(doctorempi)
	

	_, err = db.Exec(`INSERT INTO yhxd(YHBM,YHMC,SJZT,JCRQ,SHLY,RESULT,PID,MACHINE_ID,FILENAME,XMLFILENAME,ORG_CODE,ORG_NAME,YSBM,YSMC,EMPI,DOCTOR_EMPI) 
        	values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`, hci.Data.YHBM, hci.Data.YHMC, hci.Data.SJZT, hci.Data.JCRQ, "0", "", hci.Data.PID, hci.Data.MACHINE_ID,
        	hci.Data.Bullfile_id, hci.Data.Xml_bullfile_id, hci.Data.ORG_CODE, hci.Data.ORG_NAME, hci.Data.YSBM, hci.Data.YSMC, empi, doctorempi)
	if err != nil {
		fmt.Println(err)
	}
}

func UploadFatalChartFileInfo(s string, db *sql.DB) {
	fmt.Println("UploadHeartChartFileInfo")
}

