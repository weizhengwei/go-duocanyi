package main

import (
	"database/sql"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
)


type GetOrgName struct {
	PageSize int
	Esblover bool
	PageNumber int
	ORG_CODE string
	CONSUMER_ID string
	SERVICE_CODE string
}

type Json_Item struct {
	ORG_NAME string `json:"ORG_NAME"`
}

type Json_Data struct {
	Data []Json_Item `json:"data"`
}

type Json_Body struct {
	BODY Json_Data `json:"BODY"`
}

func DealOrgName(s string, db *sql.DB) {
	var getorgcode GetOrgName
	err := json.Unmarshal([]byte(s), &getorgcode)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getorgcode.ORG_CODE)
	var orgname string
	err = db.QueryRow("SELECT ORG_NAME FROM tb_organization WHERE ORG_CODE= ?", getorgcode.ORG_CODE).Scan(&orgname)
	if err != nil {
	    fmt.Println(err)
	    return
	}
	fmt.Println(orgname)
	// var v string
	// rows.Next()
	// rows.Scan(&v)
	// fmt.Printf("%s\n",v)
	
	json_item := Json_Item{ORG_NAME: orgname}
	json_data := Json_Data{}
	json_data.Data = append(json_data.Data, json_item)
	var json_body Json_Body
	json_body.BODY = json_data
	ret, err := json.Marshal(&json_body)
	fmt.Println(string(ret))
}