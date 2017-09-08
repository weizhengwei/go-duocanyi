package main

import (
	"database/sql"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
    "github.com/satori/go.uuid"
    "time"
)

//医生数据
type JsonDoctor struct {
	Data []JsonDoctorItem
}

type JsonDoctorItem struct {
	PID string
	CREDENTIAL_CODE string
	NAME string
	ORG_CODE string
	AGENCY_NAME string
	SEX string
	BIRTHDAY string
	CENSUS_REGISTER_ADDRESS string	
	MOBILE string
	EMAIL string
	PASSWORD string
	RECORDWAY string
	MACHINE_ID string
	FLAG string
	SCDATE string
}

func UploadDoctor(s string, db *sql.DB) {
	fmt.Println("UploadDoctor")
	var jd JsonDoctor
	err := json.Unmarshal([]byte(s), &jd)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jd)
	
	tx,_ := db.Begin()
    for _, v := range(jd.Data) {
        _, err := tx.Exec(`INSERT INTO tb_medical_technicians(CREATETIME$,UPDATETIME$,CURRENTSTATE,
        	CREDENTIAL_CODE,NAME,SEX,BIRTHDAY,MOBILE,AGENCY_NAME,MACHINE_ID,PID,CENSUS_REGISTER_ADDRESS,
        	ORG_CODE,FLAG,PASSWORD,SCDATE,RECORDWAY,DOCTOR_EMPI) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
        	gettime(), gettime(), "1", v.CREDENTIAL_CODE, v.NAME, v.SEX, v.BIRTHDAY, v.MOBILE, v.AGENCY_NAME, v.MACHINE_ID, v.PID, v.CENSUS_REGISTER_ADDRESS, v.ORG_CODE, v.FLAG, v.PASSWORD, v.SCDATE, v.RECORDWAY, getuuid())
        if err != nil {
        	fmt.Println(err)
        	break
        }
    }
    tx.Commit()
}

func getuuid() string {
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)
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
// 上传医生信息
private void uploadDoctor() {
    final List<DoctorBean> doctors = HealthDAO.getInstatnce().queryDoctorsForUpload();
    if (doctors.size() == 0) {
        return;
    }
    JSONObject jsonObjectForDoctor = new JSONObject();
    JSONArray jsonArrayForDoctor = new JSONArray();
    try {
        jsonObjectForDoctor.put("SERVICE_CODE", "bull.ResourcesHZ.SYN_tb_medical_technicians_CRUD");
        for (DoctorBean doctor : doctors) {
            jsonArrayForDoctor.put(HealthBoxJsonApi.toDoctorJson(doctor,
                    HealthBoxApplication.deviceBean.getDeviceId(), HealthBoxApplication.deviceBean.getOrgId(),
                    HealthBoxApplication.deviceBean.getOrgName()));
        }
        jsonObjectForDoctor.put("data", jsonArrayForDoctor);
    } catch (JSONException e) {
        Log.e("error", e.toString());
    }
    Log.e(TAG, "jsonObjectDoctor--" + jsonObjectForDoctor);

    FormBody formBody = new FormBody.Builder().add("json", jsonObjectForDoctor.toString()).build();
    UploadBean uploadBean = new UploadBean(UploadBean.DOCTOR, getString(R.string.doctor_data), false);
    list.add(uploadBean);
    OkHttpUtil.postParams(client, AppUtils.getUrl(), formBody, new UploadCallback(uploadBean) {

        @Override
        public void onBullSuccess(JSONObject response) {
            for (DoctorBean doctor : doctors) {
                HealthDAO.getInstatnce().updateFlag(SDDataBaseHelper.T_DOCTOR, doctor.getId(),
                        doctor.getSyncFlag() + 1);
            }
        }

    });
}

public static JSONObject toDoctorJson(DoctorBean doctor, String deviceId, String orgId, String orgName)
		throws JSONException {
	JSONObject params = new JSONObject();

	params.put("PID", doctor.getIdCard());
	params.put("CREDENTIAL_CODE", doctor.getIdCard());
	params.put("NAME", doctor.getName());
	params.put("ORG_CODE", orgId);
	params.put("AGENCY_NAME", orgName);
	params.put("SEX", doctor.getGender() == 1 ? "1" : "2");
	params.put("BIRTHDAY", doctor.getBirthday());
	params.put("CENSUS_REGISTER_ADDRESS", doctor.getAddr());
	params.put("MOBILE", doctor.getPhone());
	params.put("EMAIL", doctor.getMail());
	params.put("PASSWORD", doctor.getPassword());
	params.put("RECORDWAY", String.valueOf(doctor.getRecordWay()));
	params.put("MACHINE_ID", deviceId);
	params.put("FLAG", String.valueOf(doctor.getStatus()));
	params.put("SCDATE", DateUtil.getDateTime(doctor.getCreateTime()));
	// params.put("XGDATE", DateUtil.getDateTime(doctor.getUpdateTime()));
	return params;
}
*/