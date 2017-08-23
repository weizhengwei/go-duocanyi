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
	type int
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
	type int
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
	type int
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
	type int
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
	type int
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
	type int
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
	type int
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
	type int
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
	type int
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
	type int
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
	type int
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
	type int
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

// System.out.println("map.getData===" + map.get("data"));// 心电=1
// System.out.println("map.getData===" + map.get("xymbdata"));// 血氧脉搏=2
// System.out.println("map.getData===" + map.get("xydata"));// 血压=3
// System.out.println("map.getData===" + map.get("twdata"));// 体温=4
// System.out.println("map.getData===" + map.get("xtdata"));// 血糖=5
// System.out.println("map.getData===" + map.get("nydata"));// 尿常规=6
// System.out.println("map.getData===" + map.get("bmidata"));// BMI=7
// System.out.println("map.getData===" + map.get("nsdata"));// 尿酸=8
// System.out.println("map.getData===" + map.get("dgcdata"));// 胆固醇=9
// System.out.println("map.getData===" + map.get("xhdbdata"));// 血红蛋白=10
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
//kjdata 快检 JKItem type=12
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
	Kjdata []JKItem
	Thxhdbdata []HbA1cItem
}

func UploadCheckups(s string, db *sql.DB) {
	var uploadjson UploadJson
	err := json.Unmarshal([]byte(s), &uploadjson)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(uploadjson)
}