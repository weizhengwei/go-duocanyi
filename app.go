package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
	"net/url"
	"encoding/json"
)

const BIND_ADDR = ":9090"

type RequestJson struct {
	SERVICE_CODE string
}

func Home(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Home Page")
	m := req.Method
	if m == "POST" {
		fmt.Println(m)
		fmt.Println(req.URL.Path)
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(body))
	}
	if m == "GET" {
		fmt.Println(m)
		fmt.Println(req.URL.Path)
	}
}

func AllServlet(res http.ResponseWriter, req *http.Request) {
	fmt.Println("servlet page")
	m := req.Method
	if m == "POST" {
		fmt.Println(m)
		fmt.Println(req.URL.Path)
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(body))
		ss := strings.Replace(string(body), "%7B", "{", -1)
		ss = strings.Replace(ss, "%7D", "}", -1)
		ss = strings.Replace(ss, "%3A", ":", -1)
		ss = strings.Replace(ss, "%2C", ",", -1)
		ss = strings.Replace(ss, "%22", "", -1)
		fmt.Println(ss)
		resUri, pErr := url.QueryUnescape(string(body))
	    if pErr != nil {
	    	fmt.Println(pErr)
	        return
	    }
	    fmt.Println(resUri)
	    resUri = resUri[5:]
	    var jq RequestJson
	    err = json.Unmarshal([]byte(resUri), &jq)
	    if err != nil {
	    	fmt.Println(err)
	    	return
	    }
	    fmt.Println(jq.SERVICE_CODE)
	    switch jq.SERVICE_CODE {
	    case "bull.CloudPlatInterface.Get_tb_admindivision_cascade_select":

	    case "bull.ResourcesHZ.SNY_tb_equipment_CRUD":

	    default:


	    }
	}
	if m == "GET" {
		fmt.Println(m)
		fmt.Println(req.URL.Path)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/serviceProxy/servlet/", AllServlet)
	fmt.Println("server listen at", BIND_ADDR)
	log.Fatal(http.ListenAndServe(BIND_ADDR, nil))
}


/*

AutoUploadDataService.java
// 上传医生信息
uploadDoctor
bull.ResourcesHZ.SYN_tb_medical_technicians_CRUD

// 上传医生头像
uploadDoctorImage
bull.ResourcesHZ.SyN_medical_UpImage

// 上传居民信息
uploadPerson
bull.ResourcesHZ.SYN_mpi_personbasics_CRUD

// 上传居民头像
uploadPersonImage
bull.ResourcesHZ.SYN_person_UpImage

// 上传健康档案
uploadDocuments
bull.ResourcesHZ.SYN_mpi_personbasics_archives_CRUD

uploadCheckups
bull.ResourcesHZ.SNY_yh_union_CRUD

uploadReport
bull.ResourcesHZ.SNY_mpi_person_jkbg

// 上传心电检测文件
uploadHeartChart
bull.ResourcesHZ.SYN_yhxd_CRUD

uploadLocation
bull.ResourcesHZ.SYN_tb_equipment_location

AutoUploadService.java


AdminSettingActivity.java
DeviceInfoFragment.java

getOrgName
jsonObj.put("SERVICE_CODE", "bull.CloudPlatInterface.Get_tb_admindivision_cascade_select");
jsonObj.put("ORG_CODE", orgIdEt.getText().toString());
jsonObj.put("pageNumber", 0);
jsonObj.put("pageSize", 10);
jsonObj.put("esblover", false);
jsonObj.put("CONSUMER_ID", "bbeeb31c1e7542a793203cc3bc376840");


upload
bull.ResourcesHZ.SNY_tb_equipment_CRUD
*/