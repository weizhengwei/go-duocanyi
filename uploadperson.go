package main

import (
	"database/sql"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
)

//医生数据
type JsonPerson struct {
	Data []JsonPersonItem
}

type JsonPersonItem struct {
	PID string
	GLOBAL_ID string
	IDENTITY_CARDS string
	HCARDNO string
	NAME string
	GENDER string
	BIRTHDAY string
	ADDRESS string	
	TELEPHONE_NUMBER string
	NATION string
	EMAIL string
	RECORDWAY string
	SCDATE string
	CJYSBM string
	XGYSBM string
	FLAG string
	YSBM string
	MACHINE_ID string
	ORG_CODE string
	ORG_NAME string
}

func UploadPerson(s string, db *sql.DB) {
	fmt.Println("UploadPerson")
	var jp JsonPerson
	err := json.Unmarshal([]byte(s), &jp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jp)
}

/*
// 上传居民信息
private void uploadPerson() {
    final List<PersonBean> persons = HealthDAO.getInstatnce().queryPersonsForUpload();
    if (persons.size() == 0) {
        return;
    }
    JSONObject jsonObjectForPerson = new JSONObject();
    JSONArray jsonArrayForPerson = new JSONArray();
    try {
        jsonObjectForPerson.put("SERVICE_CODE", "bull.ResourcesHZ.SYN_mpi_personbasics_CRUD");

        for (PersonBean person : persons) {
            jsonArrayForPerson.put(HealthBoxJsonApi.toPersonJson(person, HealthBoxApplication.deviceBean));
        }
        jsonObjectForPerson.put("data", jsonArrayForPerson);
    } catch (JSONException e) {
        Log.e("yao", e.toString());
    }
    Log.e(TAG, "jsonObjectForPerson--" + jsonObjectForPerson);

    FormBody formBody = new FormBody.Builder().add("json", jsonObjectForPerson.toString()).build();
    UploadBean uploadBean = new UploadBean(UploadBean.PERSON, getString(R.string.person_data), false);
    list.add(uploadBean);
    OkHttpUtil.postParams(client, AppUtils.getUrl(), formBody, new UploadCallback(uploadBean) {

        @Override
        public void onBullSuccess(JSONObject response) {
            for (PersonBean person : persons) {
                HealthDAO.getInstatnce().updateIncrementFlag(SDDataBaseHelper.T_PERSON, person.getId(), 1);
            }
        }
    });
}

public static JSONObject toPersonJson(PersonBean person, DeviceBean device) throws JSONException {
	// 注册参数
	JSONObject params = new JSONObject();
	params.put("PID", person.getIdCard());
	params.put("GLOBAL_ID", person.getIdCard());
	params.put("IDENTITY_CARDS", person.getIdCard());
	params.put("HCARDNO", person.getHealthCard());
	params.put("NAME", person.getName());
	params.put("ORG_CODE", person.getOrgId());
	params.put("GENDER", person.getGender() == 1 ? "1" : "2");
	params.put("BIRTHDAY", person.getBirthday());
	params.put("ADDRESS", person.getAddr());
	params.put("TELEPHONE_NUMBER", person.getPhone());
	params.put("NATION", person.getNation());
	params.put("EMAIL", person.getMail());
	params.put("RECORDWAY", String.valueOf(person.getRecordWay()));
	params.put("SCDATE", DateUtil.getDateTime(person.getCreateTime()));
	// params.put("XGDATE", DateUtil.getDateTime(person.getUpdateTime()));
	params.put("CJYSBM", person.getCreateDoctor());
	params.put("XGYSBM", person.getUpdateDoctor());
	params.put("FLAG", String.valueOf(person.getStatus()));
	params.put("YSBM", person.getUpdateDoctor());
	params.put("MACHINE_ID", device.getDeviceId());
	params.put("ORG_CODE", device.getOrgId());
	params.put("ORG_NAME", device.getOrgName());
	return params;
}
*/