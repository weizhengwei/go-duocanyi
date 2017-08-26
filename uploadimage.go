package main

import (
	"database/sql"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
    _"encoding/json"
)



func UploadDoctorImage(s string, db *sql.DB) {
	fmt.Println("UploadDoctorImage")
}

func UploadPersonImage(s string, db *sql.DB) {
	fmt.Println("UploadPersonImage")
}