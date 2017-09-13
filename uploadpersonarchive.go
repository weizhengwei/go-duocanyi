package main

import (
	"database/sql"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
    "log"
    "github.com/satori/go.uuid"
    "time"
)

//医生数据
type JsonPersonArchive struct {
	Data []JsonPersonArchiveItem
}

type JsonPersonArchiveItem struct {
	PID string
	PERSON_ID string
	CZLX string
	JOBTYPE string
	YLFYPAY string
	LXRMC string
	LXRTEL string
	WORK_PLACE_NAME string	
	BLOODTYPE string
	RH string
	MARRIAGE_STATUS_CODE string
	DEGREE string
	RELIGION string
	GUARDIAN string
	YWGMS string
	CJQK string
	DISEASE string
	YCBS string
	BLS string
	SHHJ string
	SCDATE string
	CJYSBM string
	XGYSBM string
	YSBM string
	ORG_CODE string
	ORG_NAME string
}

func UploadPersonArchive(s string, db *sql.DB, logger *log.Logger) {
	logger.Println("UploadPersonArchive")
	var jpa JsonPersonArchive
	err := json.Unmarshal([]byte(s), &jpa)
	if err != nil {
		fmt.Println(err)
		return
	}
	logger.Println(jpa)
	tx,_ := db.Begin()
	time := gettime()
	uuid := getuuid()
    for _, v := range(jpa.Data) {
    	_, err := tx.Exec(`update mpi_person_archives set PERSON_ID='?',YWGMS='?',DISEASE_JB='?',DISEASE_SS='?',
    		DISEASE_WS='?',DISEASE_SX='?',FAMILY_FATHER='?',FAMILY_MOTHER='?',FAMILY_BROTHER='?',FAMILY_CHILDREN='?',YCBS='?',CJQK='?',LIVE_CF='?',LIVE_RL='?',LIVE_WATER='?',LIVE_WC='?',LIVE_QC='?',BLS='?',DISEASE='?',LIVE='?',UPDATE_TIEM=now() where PERSON_ID='?'`,
	    	v.PERSON_ID, v.NAME, v.GENDER, v.BIRTHDAY, "", v.NATION, v.ADDRESS, v.IDENTITY_CARDS, v.TELEPHONE_NUMBER, time, 
	    	time, "1", v.ORG_CODE, v.PID, v.MACHINE_ID, v.RECORDWAY, v.ORG_NAME, v.CJYSBM, "1", uuid, "1", "0", 
	    	v.NAME, v.GENDER, v.BIRTHDAY, "", v.NATION, v.ADDRESS, v.TELEPHONE_NUMBER, time, "1", v.ORG_CODE, v.MACHINE_ID, 
	    	v.RECORDWAY, v.ORG_NAME, v.CJYSBM, "1")
	    if err != nil {
		    fmt.Println(err)
	    	break
    	}
    }
    tx.Commit()
}

func getuuid() string {
	u1 := uuid.NewV4()
	//fmt.Printf("UUIDv4: %s\n", u1)
	s := u1.String()
	ss := s[0:8] + s[9:13] + s[14:18] + s[19:23]+s[24:]
	return ss
}

func gettime() string {
	var datetime = time.Now()
	layout := "2006-01-02 15:04:05"
	return datetime.Format(layout)
}
/*
// 上传健康档案
private void uploadDocuments() {
    final List<PersonBean> persons = HealthDAO.getInstatnce().queryDocumentsForUpload();
    if (persons.size() == 0) {
        return;
    }
    JSONObject jsonObjectForDocument = new JSONObject();
    JSONArray jsonArrayForDocument = new JSONArray();
    try {
        jsonObjectForDocument.put("SERVICE_CODE", "bull.ResourcesHZ.SYN_mpi_personbasics_archives_CRUD");

        for (PersonBean person : persons) {
            jsonArrayForDocument.put(HealthBoxJsonApi.toDocumentJson(person, HealthBoxApplication.deviceBean));
        }
        jsonObjectForDocument.put("data", jsonArrayForDocument);
    } catch (JSONException e) {
        Log.e("yao", e.toString());
    }
    Log.e("yao", "jsonObjectForDocument--" + jsonObjectForDocument);

    FormBody formBody = new FormBody.Builder().add("json", jsonObjectForDocument.toString()).build();
    UploadBean uploadBean = new UploadBean(UploadBean.DOCUMENT, getString(R.string.person_doc), false);
    list.add(uploadBean);
    OkHttpUtil.postParams(client, AppUtils.getUrl(), formBody, new UploadCallback(uploadBean) {

        @Override
        public void onBullSuccess(JSONObject response) {
            for (PersonBean person : persons) {
                HealthDAO.getInstatnce().updateIncrementFlag(SDDataBaseHelper.T_PERSON, person.getId(), 4);
            }
        }

    });
}
public static JSONObject toDocumentJson(PersonBean person, DeviceBean device) throws JSONException {
	// 注册参数
	JSONObject params = new JSONObject();
	params.put("PID", person.getIdCard());
	params.put("PERSON_ID", person.getIdCard());
	if (person.getJBXX() != null) {
		JSONObject jsonJBXX = new JSONObject(person.getJBXX());
		params.put("CZLX", AppUtils.getJson(jsonJBXX, "addtype"));
		params.put("JOBTYPE", AppUtils.getJson(jsonJBXX, "occupation"));
		params.put("YLFYPAY", AppUtils.getJson(jsonJBXX, "payment"));
		params.put("LXRMC", AppUtils.getJson(jsonJBXX, "contactname"));
		params.put("LXRTEL", AppUtils.getJson(jsonJBXX, "contactphone"));
		params.put("WORK_PLACE_NAME", AppUtils.getJson(jsonJBXX, "workplace"));
		params.put("BLOODTYPE", AppUtils.getJson(jsonJBXX, "bloodtype"));
		params.put("RH", AppUtils.getJson(jsonJBXX, "rhtype"));
		params.put("MARRIAGE_STATUS_CODE", AppUtils.getJson(jsonJBXX, "marriage"));
		params.put("DEGREE", AppUtils.getJson(jsonJBXX, "education"));
		params.put("RELIGION", AppUtils.getJson(jsonJBXX, "religion"));
		params.put("GUARDIAN", AppUtils.getJson(jsonJBXX, "guardian"));
	}
	params.put("YWGMS", person.getYWGM());
	params.put("CJQK", person.getCJQK());
	params.put("DISEASE", person.getJWBS());
	params.put("YCBS", person.getYCBS());
	params.put("BLS", person.getBLS());
	params.put("SHHJ", person.getSHHJ());
	params.put("SCDATE", DateUtil.getDateTime(person.getCreateTime()));
	// params.put("XGDATE", DateUtil.getDateTime(person.getUpdateTime()));
	params.put("CJYSBM", person.getCreateDoctor());
	params.put("XGYSBM", person.getUpdateDoctor());
	params.put("YSBM", person.getUpdateDoctor());
	params.put("ORG_CODE", device.getOrgId());
	params.put("ORG_NAME", device.getOrgName());
	return params;
}
json={
    "data": [
        {
            "CJYSBM": "130523198608251631",
            "ORG_NAME": "Frankfurt",
            "PID": "110101196011116995",
            "PERSON_ID": "110101196011116995",
            "YSBM": "",
            "ORG_CODE": "10086",
            "XGYSBM": "",
            "SCDATE": "2017-09-13 16:58:41"
        }
    ],
    "SERVICE_CODE": "bull.ResourcesHZ.SYN_mpi_personbasics_archives_CRUD"
}
*/