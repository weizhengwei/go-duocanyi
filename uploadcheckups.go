package main

import (
	"database/sql"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
)

//血氧脉搏数据
type SPOItem struct {
	PID string
	Type int
	YHBM string
	YHMC string
	YSBM string
	YSMC string
	JCRQ string
	XY string
	MB string
	SJZT string
	MACHINE_ID string
	ORG_CODE string
	ORG_NAME string	
}

//血压数据
type NIBPItem struct{
	PID string
	Type int
	YHBM string
	YHMC string
	YSBM string
	YSMC string
	JCRQ string
	SSY string
	SZY string
	PJY string
	RESULT string
	SJZT string
	MACHINE_ID string
	ORG_CODE string
	ORG_NAME string
}

//体温数据
type TemperatureItem struct {
	PID string
	Type int
	YHBM string
	YHMC string
	YSBM string
	YSMC string
	JCRQ string
	WD string
	SJZT string
	MACHINE_ID string
	ORG_CODE string
	ORG_NAME string
}

//血糖数据
type GluItem struct {
	PID string
	Type int
	YHBM string
	YHMC string
	YSBM string
	YSMC string
	JCRQ string
	XT string
	SJZT string
	MACHINE_ID string
	ORG_CODE string
	ORG_NAME string
}

//尿常规数据
type UrineItem struct {
	PID string
	Type int
	YHBM string
	YHMC string
	YSBM string
	YSMC string
	JCRQ string
	PH string
	YXSY string
	DBZ string
	PTT string
	DHS string
	NDF string
	TT string
	BZ string
	BXB string
	WSSC string
	HXB string
	SJZT string
	MACHINE_ID string
	ORG_CODE string
	ORG_NAME string
}

type BmiItem struct {
	PID string
	Type int
	YHBM string
	YHMC string
	YSBM string
	YSMC string
	JCRQ string
	SG string
	TZ string
	YW string
	BMI string
	SJZT string
	MACHINE_ID string
	ORG_CODE string
	ORG_NAME string
}

//尿酸
type UaItem struct {
	PID string
	Type int
	YHBM string
	YHMC string
	YSBM string
	YSMC string
	JCRQ string	
	NS string
	SJZT string
	MACHINE_ID string
	ORG_CODE string
	ORG_NAME string
}

//胆固醇
type CholItem struct {
	PID string
	Type int
	YHBM string
	YHMC string
	YSBM string
	YSMC string
	JCRQ string
	DGC string
	SJZT string
	MACHINE_ID string
	ORG_CODE string
	ORG_NAME string
}

//血红蛋白
type BfItem struct {
	PID string
	Type int
	YHBM string
	YHMC string
	YSBM string
	YSMC string
	JCRQ string	
	CHOL string
	HDL string
	TG string
	LDL string
	SJZT string
	MACHINE_ID string
	ORG_CODE string
	ORG_NAME string
}

type KJItem struct {
	PID string
	Type int
	YHBM string
	YHMC string
	YSBM string
	YSMC string
	JCRQ string
	SHORT_NAME string
	PRO_VALUE string
	PRO_TYPE string
	SJZT string
	MACHINE_ID string
	ORG_CODE string
	ORG_NAME string
}

type HbA1cItem struct {
	PID string
	Type int
	YHBM string
	YHMC string
	YSBM string
	YSMC string
	JCRQ string
	NGSP string
	IFCC string
	SJZT string
	MACHINE_ID string
	ORG_CODE string
	ORG_NAME string
}

type HbItem struct {
	PID string
	Type int
	YHBM string
	YHMC string
	YSBM string
	YSMC string
	JCRQ string
	HB string
	HCT string
	SJZT string
	MACHINE_ID string
	ORG_CODE string
	ORG_NAME string
}

// System.out.print8ln("map.getData===" + map.get("data"));// 心电=1
// System.out.print8ln("map.getData===" + map.get("xymbdata"));// 血氧脉搏=2
// System.out.print8ln("map.getData===" + map.get("xydata"));// 血压=3
// System.out.print8ln("map.getData===" + map.get("twdata"));// 体温=4
// System.out.print8ln("map.getData===" + map.get("xtdata"));// 血糖=5
// System.out.print8ln("map.getData===" + map.get("nydata"));// 尿常规=6
// System.out.print8ln("map.getData===" + map.get("bmidata"));// BMI=7
// System.out.print8ln("map.getData===" + map.get("nsdata"));// 尿酸=8
// System.out.print8ln("map.getData===" + map.get("dgcdata"));// 胆固醇=9
// System.out.print8ln("map.getData===" + map.get("xhdbdata"));// 血红蛋白=10
//xzdata 血脂 bfArray
//kjdata 快检 kjArray
//thxhdbdata 糖化血红蛋白 hba1cArray

//血氧脉搏数据 SPOItem type=2
//血压数据 NIBPItem type=3
//体温数据 TemperatureItem type=4
//血糖数据 GluItem type=5
//尿常规数据 UrineItem type=6
//BMI BmiItem type=7
//尿酸 UaItem type=8
//胆固醇 CholItem type=9
//血红蛋白 HbItem type=10
//xzdata 血脂 BfItem type=11
//kjdata 快检 KJItem type=12
//thxhdbdata 糖化血红蛋白 HbA1cItem type=13

type UploadJson struct {
	Xymbdata []SPOItem
	Xydata []NIBPItem
	Twdata []TemperatureItem
	Xtdata []GluItem
	Nydata []UrineItem
	Bmidata []BmiItem
	Nsdata []UaItem
	Dgcdata []CholItem
	Xhdbdata []HbItem
	Xzdata []BfItem
	Kjdata []KJItem
	Thxhdbdata []HbA1cItem
}

/*
json={
    "xymbdata": [
        {
            "MACHINE_ID": "100251",
            "YHBM": "370811197402045542",
            "YHMC": "李淑梅",
            "ORG_NAME": "喻屯镇西王楼村卫生室",
            "PID": "100251C1503565817980",
            "SJZT": "1",
            "XY": "98",
            "JCRQ": "2017-08-24 17:10:17",
            "YSBM": "370811197402265537",
            "ORG_CODE": "37081103053",
            "YSMC": "王德生",
            "type": 2,
            "MB": "70"
        }
    ],
    "SERVICE_CODE": "bull.ResourcesHZ.SNY_yh_union_CRUD",
    "xydata": [
        {
            "YHBM": "370811197402045542",
            "YHMC": "李淑梅",
            "ORG_NAME": "喻屯镇西王楼村卫生\n室",
            "PID": "100251C1503565958793",
            "YSMC": "王德生",
            "type": 3,
            "SZY": "71",
            "MACHINE_ID": "100251",
            "SSY": "104",
            "RESULT": "成人模式",
            "SJZT": "1",
            "JCRQ": "2017-08-24 17:12:38",
            "YSBM": "370811197402265537",
            "ORG_CODE": "37081103053",
            "PJY": "82"
        }
    ]
}
*/

func UploadCheckups(s string, db *sql.DB) {
	var uploadjson UploadJson
	err := json.Unmarshal([]byte(s), &uploadjson)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(uploadjson)
	if len(uploadjson.Xymbdata) != 0 {
		upload_spo_data(uploadjson.Xymbdata, db)
	}
}

//
func upload_spo_data(data []SPOItem, db *sql.DB) {
	fmt.Println("upload_spo_data")
	fmt.Println(data)
    tx,_ := db.Begin()
    for _, v := range(data) {
        //每次循环用的都是tx内部的连接，没有新建连接，效率高
        tx.Exec(`INSERT INTO yhxy01(YHBM,YHMC,XY,MB,SJZT,JCRQ,SHLY,RESULT,PID,MACHINE_ID,ORG_CODE,ORG_NAME,YSBM,YSMC) 
        	values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`, v.YHBM, v.YHMC, v.XY, v.MB, v.SJZT, v.JCRQ, "0", "", v.PID, v.MACHINE_ID, 
        	v.ORG_CODE, v.ORG_NAME, v.YSBM, v.YSMC)
    }
    //最后释放tx内部的连接
    tx.Commit()
}

func upload_nibp_data(data []NIBPItem, db *sql.DB) {
	fmt.Println("upload_nibp_data")
	// fmt.Println(data)
	// tx,_ := db.Begin()
 //    for _, v := range(data) {
 //        //每次循环用的都是tx内部的连接，没有新建连接，效率高
 //        tx.Exec(`INSERT INTO yhxy01(YHBM,YHMC,XY,MB,SJZT,JCRQ,SHLY,RESULT,PID,MACHINE_ID,ORG_CODE,ORG_NAME,YSBM,YSMC) 
 //        	values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`, v.YHBM, v.YHMC, v.XY, v.MB, v.SJZT, v.JCRQ, "0", "", v.PID, v.MACHINE_ID, 
 //        	v.ORG_CODE, v.ORG_NAME, v.YSBM, v.YSMC)
 //    }
 //    //最后释放tx内部的连接
 //    tx.Commit()
}

func upload_tmperature_data(data []TemperatureItem, db *sql.DB) {

}

func upload_glu_data(data []GluItem, db *sql.DB) {

}

func upload_urine_data(data []UrineItem, db *sql.DB) {

}

func upload_bmi_data(data []BmiItem, db *sql.DB) {

}

func upload_chol_data(data []CholItem, db *sql.DB) {

}

func upload_hb_data(data []HbItem, db *sql.DB) {

}

func upload_bf_data(data []BfItem, db *sql.DB) {

}

func upload_kj_data(data []KJItem, db *sql.DB) {

}

func upload_hba1c_data(data []HbA1cItem, db *sql.DB) {

}
