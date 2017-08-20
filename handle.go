package main

import (
	"database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

type CustomDB struct{
	db *sql.DB
}

func (x *CustomDB)Init() error{
	fmt.Println("db init")
	var err error
	x.db, err = sql.Open("mysql", "root:r00t@localhost:3306/med_gwc?charset=utf8")
	fmt.Println("Init", err)
	return err
}

func DealOrgName() {
	
}

func DealRegister() {
	
}

func DealDoctor(json string) {
	
}

func DealDoctorImage(json string) {
	
}

func DealPerson(json string) {
	
}

func DealPersonImage(json string) {
	
}

func DealDocuments(json string) {
	
}

func DealCheckups(json string) {
	
}

func DealReoprt(json string) {
	
}

func DealHeartChart() {
	
}