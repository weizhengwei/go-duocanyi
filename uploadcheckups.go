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
	MB string
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

//血脂
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

//糖化血红蛋白
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

func UploadCheckups(s string, db *sql.DB) {
	var uploadjson UploadJson
	err := json.Unmarshal([]byte(s), &uploadjson)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(uploadjson)

	if len(uploadjson.Xymbdata) != 0 {
		upload_spo_data(uploadjson.Xymbdata, db)//血氧
	}
	if len(uploadjson.Xydata) != 0 {
		upload_nibp_data(uploadjson.Xydata, db)//血压
	}
	if len(uploadjson.Twdata) != 0 {
		upload_tmperature_data(uploadjson.Twdata, db)//体温
	}
	if len(uploadjson.Xtdata) != 0 {
		upload_glu_data(uploadjson.Xtdata, db)//血糖数据
	}
	if len(uploadjson.Nydata) != 0 {
		upload_urine_data(uploadjson.Nydata, db)//尿常规数据
	}
	if len(uploadjson.Bmidata) != 0 {
		upload_bmi_data(uploadjson.Bmidata, db)//BMI数据
	}
	if len(uploadjson.Nsdata) != 0 {
		upload_ua_data(uploadjson.Nsdata, db)//尿酸数据
	}
	if len(uploadjson.Dgcdata) != 0 {
		upload_chol_data(uploadjson.Dgcdata, db)//胆固醇
	}
	if len(uploadjson.Xhdbdata) != 0 {
		upload_hb_data(uploadjson.Xhdbdata, db)//血红蛋白
	}
	if len(uploadjson.Xzdata) != 0 {
		upload_bf_data(uploadjson.Xzdata, db)//血脂
	}
	if len(uploadjson.Kjdata) != 0 {
		upload_kj_data(uploadjson.Kjdata, db)//快检
	}
	if len(uploadjson.hxhdbdata) != 0 {
		upload_hba1c_data(uploadjson.hxhdbdata, db)//糖化血红蛋白
	}
}

//血氧(血氧，脉率) upload_spo_data
//血压 upload_nibp_data
//体温 upload_tmperature_data
//血糖数据 upload_glu_data
//尿常规数据 upload_urine_data
//BMI数据 upload_bmi_data
//尿酸数据 upload_ua_data
//胆固醇 upload_chol_data
//血红蛋白 upload_hb_data
//血脂 upload_bf_data
//快检 upload_kj_data
//糖化血红蛋白 upload_hba1c_data


//血氧(血氧，脉率)
func upload_spo_data(data []SPOItem, db *sql.DB) {
	fmt.Println("upload_spo_data 上传血氧数据")
	fmt.Printf("%v\n", data)
    tx,_ := db.Begin()
    for _, v := range(data) {
        
        tx.Exec(`INSERT INTO yhxy01(YHBM,YHMC,XY,MB,SJZT,JCRQ,SHLY,RESULT,PID,MACHINE_ID,ORG_CODE,ORG_NAME,YSBM,YSMC) 
        	values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`, v.YHBM, v.YHMC, v.XY, v.MB, v.SJZT, v.JCRQ, "0", "", v.PID, v.MACHINE_ID, 
        	v.ORG_CODE, v.ORG_NAME, v.YSBM, v.YSMC)
    }
    
    tx.Commit()
}

//血压(收缩压，舒张压，平均压，脉搏)
func upload_nibp_data(data []NIBPItem, db *sql.DB) {
	fmt.Println("upload_nibp_data 上传血压数据")
	fmt.Printf("%v\n", data)
	tx,_ := db.Begin()
    for _, v := range(data) {
        
        _, err := tx.Exec(`INSERT INTO yhxy(YHBM,YHMC,SSY,SZY,PJY,SJZT,JCRQ,SHLY,RESULT,PID,MACHINE_ID,ORG_CODE,ORG_NAME,YSBM,YSMC) 
        	values(?,?,?,?,?,?,?,?,?,?,?,?,?,?, ?)`, v.YHBM, v.YHMC, v.SSY, v.SZY, v.PJY, v.SJZT, v.JCRQ, "0", v.RESULT, v.PID, v.MACHINE_ID, 
        	v.ORG_CODE, v.ORG_NAME, v.YSBM, v.YSMC)
        if err != nil {
        	fmt.Println(err)
        }
    }
    
    err := tx.Commit()
    if err != nil {
    	fmt.Println(err)
    }
}

//体温(温度)
func upload_tmperature_data(data []TemperatureItem, db *sql.DB) {
	fmt.Println("upload_tmperature_data 上传体温数据")
	fmt.Printf("%v\n", data)
    tx,_ := db.Begin()
    for _, v := range(data) {
        
        tx.Exec(`INSERT INTO yhtw(YHBM,YHMC,WD,SJZT,JCRQ,SHLY,RESULT,PID,MACHINE_ID,ORG_CODE,ORG_NAME,YSBM,YSMC) 
        	values(?,?,?,?,?,?,?,?,?,?,?,?,?)`, v.YHBM, v.YHMC, v.WD, v.SJZT, v.JCRQ, "0", "", v.PID, v.MACHINE_ID, 
        	v.ORG_CODE, v.ORG_NAME, v.YSBM, v.YSMC)
    }
    
    tx.Commit()
}

//血糖数据(血糖)
func upload_glu_data(data []GluItem, db *sql.DB) {
	fmt.Println("upload_glu_data 上传血糖数据")
	fmt.Printf("%v\n", data)
    tx,_ := db.Begin()
    for _, v := range(data) {
        
        tx.Exec(`INSERT INTO yhxt(YHBM,YHMC,XT,SJZT,JCRQ,SHLY,RESULT,PID,MACHINE_ID,ORG_CODE,ORG_NAME,YSBM,YSMC) 
        	values(?,?,?,?,?,?,?,?,?,?,?,?,?)`, v.YHBM, v.YHMC, v.XT, v.SJZT, v.JCRQ, "0", "", v.PID, v.MACHINE_ID, 
        	v.ORG_CODE, v.ORG_NAME, v.YSBM, v.YSMC)
    }
    
    tx.Commit()
}

//尿常规数据(白细胞,亚硝酸盐,尿胆原,蛋白质,PH,比重,酮体,胆红素,葡萄糖,维生素C)
func upload_urine_data(data []UrineItem, db *sql.DB) {
	fmt.Println("upload_urine_data 上传尿常规数据")
	fmt.Printf("%v\n", data)
    tx,_ := db.Begin()
    for _, v := range(data) {
        
        tx.Exec(`INSERT INTO yhny(YHBM,YHMC,PH,YXSY,DBZ,PTT,DHS,NDF,TT,BZ,BXB,WSSC,HXB,SJZT,JCRQ,SHLY,RESULT,PID,MACHINE_ID,ORG_CODE,ORG_NAME,YSBM,YSMC) 
        	values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`, v.YHBM, v.YHMC, v.PH, v.YXSY, v.DBZ, v.PTT, 
        	v.DHS, v.NDF, v.TT, v.BZ, v.BXB, v.WSSC, v.HXB, v.SJZT, v.JCRQ, "0", "", v.PID, v.MACHINE_ID, v.ORG_CODE, 
        	v.ORG_NAME, v.YSBM, v.YSMC)
    }
    
    tx.Commit()
}

//BMI数据(身高,体重,腰围,BMI)
func upload_bmi_data(data []BmiItem, db *sql.DB) {
	fmt.Println("upload_bmi_data 上传BMI数据")
	fmt.Printf("%v\n", data)
    tx,_ := db.Begin()
    for _, v := range(data) {
        
        tx.Exec(`INSERT INTO yhbmi(YHBM,YHMC,SG,TZ,YW,BMI,SJZT,JCRQ,SHLY,RESULT,PID,MACHINE_ID,ORG_CODE,ORG_NAME,YSBM,YSMC) 
        	values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`, v.YHBM, v.YHMC, v.SG, v.TZ, v.YW, v.BMI, v.SJZT, v.JCRQ, "0", "", v.PID,
        	v.MACHINE_ID, v.ORG_CODE, v.ORG_NAME, v.YSBM, v.YSMC)
    }
    
    tx.Commit()
}

//尿酸数据(尿酸)
func upload_ua_data(data []UaItem, db *sql.DB) {
	fmt.Println("upload_ua_data 上传尿酸数据")
	fmt.Printf("%v\n", data)
    tx,_ := db.Begin()
    for _, v := range(data) {
        
        tx.Exec(`INSERT INTO yhns(YHBM,YHMC,NS,SJZT,JCRQ,SHLY,RESULT,PID,MACHINE_ID,ORG_CODE,ORG_NAME,YSBM,YSMC) 
        	values(?,?,?,?,?,?,?,?,?,?,?,?,?)`, v.YHBM, v.YHMC, v.NS, v.SJZT, v.JCRQ, "0", "", v.PID, v.MACHINE_ID, 
        	v.ORG_CODE, v.ORG_NAME, v.YSBM, v.YSMC)
    }
    
    tx.Commit()
}

//胆固醇(胆固醇)
func upload_chol_data(data []CholItem, db *sql.DB) {
	fmt.Println("upload_chol_data 上传胆固醇数据")
	fmt.Printf("%v\n", data)
    tx,_ := db.Begin()
    for _, v := range(data) {
        
        tx.Exec(`INSERT INTO yhdgc(YHBM,YHMC,DGC,SJZT,JCRQ,SHLY,RESULT,PID,MACHINE_ID,ORG_CODE,ORG_NAME,YSBM,YSMC) 
        	values(?,?,?,?,?,?,?,?,?,?,?,?,?)`, v.YHBM, v.YHMC, v.DGC, v.SJZT, v.JCRQ, "0", "", v.PID, v.MACHINE_ID, 
        	v.ORG_CODE, v.ORG_NAME, v.YSBM, v.YSMC)
    }
    
    tx.Commit()
}

//血红蛋白(HB,HCT)
func upload_hb_data(data []HbItem, db *sql.DB) {
	fmt.Println("upload_hb_data 上传血红蛋白数据")
	fmt.Printf("%v\n", data)
    tx,_ := db.Begin()
    for _, v := range(data) {
        
        tx.Exec(`INSERT INTO yhxhdb(YHBM,YHMC,HB,HCT,SJZT,JCRQ,SHLY,RESULT,PID,MACHINE_ID,ORG_CODE,ORG_NAME,YSBM,YSMC) 
        	values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`, v.YHBM, v.YHMC, v.HB, v.HCT, v.SJZT, v.JCRQ, "0", "", v.PID, v.MACHINE_ID, 
        	v.ORG_CODE, v.ORG_NAME, v.YSBM, v.YSMC)
    }
    
    tx.Commit()
}

//血脂(总胆固醇,高密度脂蛋白,甘油三酯,低密度脂蛋白)
func upload_bf_data(data []BfItem, db *sql.DB) {
	fmt.Println("upload_bf_data 上传血脂数据")
	fmt.Printf("%v\n", data)
    tx,_ := db.Begin()
    for _, v := range(data) {
        
        tx.Exec(`INSERT INTO yhxz(YHBM,YHMC,CHOL,HDL,TG,LDL,SJZT,JCRQ,SHLY,RESULT,PID,MACHINE_ID,ORG_CODE,ORG_NAME,YSBM,YSMC) 
        	values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`, v.YHBM, v.YHMC, v.CHOL, v.HDL, v.TG, v.LDL, v.SJZT, v.JCRQ, "0", "", v.PID, v.MACHINE_ID, 
        	v.ORG_CODE, v.ORG_NAME, v.YSBM, v.YSMC)
    }
    
    tx.Commit()
}

//SHORT_NAME PRO_VALUE PRO_TYPE
func upload_kj_data(data []KJItem, db *sql.DB) {
	fmt.Println("upload_hba1c_data 上传糖化血红蛋白数据")
	fmt.Printf("%v\n", data)
    tx,_ := db.Begin()
    for _, v := range(data) {
        
        tx.Exec(`INSERT INTO yhkj(YHBM,YHMC,SHORT_NAME,PRO_VALUE,PRO_TYPE,SJZT,JCRQ,SHLY,RESULT,PID,MACHINE_ID,ORG_CODE,ORG_NAME,YSBM,YSMC) 
        	values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`, v.YHBM, v.YHMC, v.SHORT_NAME, v.PRO_VALUE, v.PRO_TYPE, v.SJZT, v.JCRQ, "0", "", v.PID, v.MACHINE_ID, 
        	v.ORG_CODE, v.ORG_NAME, v.YSBM, v.YSMC)
    }
    
    tx.Commit()
}

//糖化血红蛋白 NGSP,IFCC
func upload_hba1c_data(data []HbA1cItem, db *sql.DB) {
	fmt.Println("upload_hba1c_data 上传糖化血红蛋白数据")
	fmt.Printf("%v\n", data)
    tx,_ := db.Begin()
    for _, v := range(data) {
        
        tx.Exec(`INSERT INTO yhdhxhdb(YHBM,YHMC,NGSP,IFCC,SJZT,JCRQ,SHLY,RESULT,PID,MACHINE_ID,ORG_CODE,ORG_NAME,YSBM,YSMC) 
        	values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`, v.YHBM, v.YHMC, v.NGSP, v.IFCC, v.SJZT, v.JCRQ, "0", "", v.PID, v.MACHINE_ID, 
        	v.ORG_CODE, v.ORG_NAME, v.YSBM, v.YSMC)
    }
    
    tx.Commit()
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