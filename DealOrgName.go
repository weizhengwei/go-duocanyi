package main

import (
	"database/sql"
	"fmt"
	"net/http"
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

func DealOrgName(res http.ResponseWriter, s string, db *sql.DB) {
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
	res.Write(ret)
}



/*
private void getOrgName() {
    JSONObject jsonObj = new JSONObject();
    try {
        jsonObj.put("SERVICE_CODE", "bull.CloudPlatInterface.Get_tb_admindivision_cascade_select");
        jsonObj.put("ORG_CODE", orgIdEt.getText().toString());
        jsonObj.put("pageNumber", 0);
        jsonObj.put("pageSize", 10);
        jsonObj.put("esblover", false);
        jsonObj.put("CONSUMER_ID", "bbeeb31c1e7542a793203cc3bc376840");
    } catch (JSONException e) {
        Log.e("yao", e.toString());
    }

    RequestParams params = new RequestParams();
    params.addBodyParameter("json", jsonObj.toString());
    Log.e("yao", "获取设备名称 : " + jsonObj.toString());
    HttpUtils http = new HttpUtils();
    http.send(HttpMethod.POST, AppUtils.getUrl(), params, new BullRequestCallBack() {

        @Override
        public void onBullSuccess(JSONObject response) {
            try {
                JSONArray data = response.getJSONObject("BODY").getJSONArray("data");
                final String name = data.size() > 0 ? data.getJSONObject(0).getString("ORG_NAME") : "";

                Activity activity = getActivity();

                if (activity != null) {
                    activity.runOnUiThread(new Runnable() {

                        @Override
                        public void run() {
                            orgNameEt.setText(name);
                            orgNameEt.setSelection(orgNameEt.getText().length());
                        }
                    });
                }
            } catch (JSONException e) {
                Log.e("yao", e.toString());
            }
        }

        @Override
        public void onBullFailure(JSONObject response) {

        }

        @Override
        public void onFailure(HttpException httpexception, String s) {
        }
    });
}
*/