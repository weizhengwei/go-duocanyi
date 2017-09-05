package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"io"
	"os"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
)

type ImageFileInfo struct {
	Data ImageFileInfoItem
}

type ImageFileInfoItem struct {
	PID string
	NAME string
}

func UploadImageFile(res http.ResponseWriter, req *http.Request, basepath string) {
	req.ParseMultipartForm(32 << 20)
	file, handler, err := req.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	f, err := os.OpenFile(basepath + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Fprintf(res, "upload ok")
}

func UploadDoctorImage(s string, db *sql.DB) {
	fmt.Println("UploadDoctorImage")
	var ifi ImageFileInfo
	err := json.Unmarshal([]byte(s), &ifi)
    if err != nil {
    	fmt.Println(err)
    	return
    }
    fmt.Println("PID", ifi.Data.PID, "NAME", ifi.Data.NAME)
    tx,_ := db.Begin()
    tx.Exec("UPDATE tb_medical_technicians set FILENAME=?, UPDATETIME$=? where PID=?","", "", ifi.Data.PID)
    tx.Commit()
}

func UploadPersonImage(s string, db *sql.DB) {
	fmt.Println("UploadPersonImage")
	var ifi ImageFileInfo
	err := json.Unmarshal([]byte(s), &ifi)
    if err != nil {
    	fmt.Println(err)
    	return
    }
    fmt.Println("PID", ifi.Data.PID, "NAME", ifi.Data.NAME)
    tx,_ := db.Begin()
    tx.Exec("UPDATE mpi_personbasics set FILENAME=?, CHANGED_TIME=? where PID=?","", "", ifi.Data.PID)
    tx.Commit()
}