package main

import (
	"fmt"
	"net/http"
	"log"
	"io"
	"io/ioutil"
	_"strings"
	"net/url"
	"encoding/json"
	"os"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
    "strconv"
)

const BIND_ADDR = ":9090"
var db *sql.DB
var logger *log.Logger

type RequestJson struct {
	SERVICE_CODE string
}

func Home(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Home Page")
	m := req.Method
	if m == "POST" {
		fmt.Println(m)
		fmt.Println(req.URL.Path)
		// body, err := ioutil.ReadAll(req.Body)
		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	return
		// }
		//fmt.Println(string(body))
	}
	if m == "GET" {
		fmt.Println(m)
		fmt.Println(req.URL.Path)
	}
}

func AllServlet(res http.ResponseWriter, req *http.Request) {
	logger.Println("AllServlet")
	m := req.Method
	if m == "POST" {
		logger.Println(req.URL.Path)
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		urljson := req.URL.Query().Get("json")
		if len(urljson) != 0 {
			var uj RequestJson
		    err := json.Unmarshal([]byte(urljson), &uj)
		    if err != nil {
		    	fmt.Println(err)
		    	return
		    }
		    switch uj.SERVICE_CODE {
		    	case "bull.ResourcesHZ.SyN_medical_UpImage":// 上传医生头像
		    		UploadImageFile(res, req, "./Image/Doctor/")
		    		UploadDoctorImage(urljson, db)
		    	case "bull.ResourcesHZ.SYN_person_UpImage":// 上传居民头像
		    		UploadImageFile(res, req, "./Image/Person/")
		    		UploadPersonImage(urljson, db)
		    	case "bull.ResourcesHZ.SYN_yhxd_CRUD":// 上传心电检测文件

		    	case "bull.ResourcesHZ.SYN_yhfetalhm_CRUD":// 上传胎监文件

		    }
		    return
		}
		// fmt.Println(string(body))

		// ss := strings.Replace(string(body), "%7B", "{", -1)
		// ss = strings.Replace(ss, "%7D", "}", -1)
		// ss = strings.Replace(ss, "%3A", ":", -1)
		// ss = strings.Replace(ss, "%2C", ",", -1)
		// ss = strings.Replace(ss, "%22", "", -1)
		// fmt.Println(ss)
		PostJsonBody, err := url.QueryUnescape(string(body))
	    if err != nil {
	    	fmt.Println(err)
	        return
	    }
	    logger.Println(PostJsonBody)
	    PostJsonBody = PostJsonBody[5:]
	    var jq RequestJson
	    err = json.Unmarshal([]byte(PostJsonBody), &jq)
	    if err != nil {
	    	fmt.Println(err)
	    	return
	    }
	    
	    switch jq.SERVICE_CODE {
	    case "bull.CloudPlatInterface.Get_tb_admindivision_cascade_select"://根据机构ID获取机构名称
	    	fmt.Println(jq.SERVICE_CODE)
	    	DealOrgName(res, PostJsonBody, db)
	    case "bull.ResourcesHZ.SNY_tb_equipment_CRUD"://注册多参仪
	    	DealRegister(PostJsonBody, db)
	    case "bull.ResourcesHZ.Down_tb_medical_technicians"://下拉医生信息
	    	DownloadDoctor(PostJsonBody, db)
	    case "bull.ResourcesHZ.Down_mpi_personbasics"://下拉居民信息

	    case "bull.ResourcesHZ.SYN_tb_medical_technicians_CRUD":// 上传医生信息
	    	UploadDoctor(PostJsonBody, db, logger)
	    case "bull.ResourcesHZ.SYN_mpi_personbasics_CRUD":// 上传居民信息
	    	UploadPerson(PostJsonBody, db)
	    case "bull.ResourcesHZ.SYN_mpi_personbasics_archives_CRUD":// 上传健康档案

	    case "bull.ResourcesHZ.SNY_yh_union_CRUD"://uploadCheckups
	    	UploadCheckups(PostJsonBody, db)
	    case "bull.ResourcesHZ.SNY_mpi_person_jkbg":// 上传健康报告
	    	UploadReport(PostJsonBody, db)
	    case "bull.ResourcesHZ.SNY_yhxd_new_CRUD":// 上传完了心电检测文件，上传的文件信息
	    	UploadHeartChartFileInfo(PostJsonBody, db)
	    case "bull.ResourcesHZ.SNY_yhfetalhm_CRUD":// 上传完了胎监文件，上传的文件信息
	    	UploadFatalChartFileInfo(PostJsonBody, db)
	    case "bull.ResourcesHZ.SNY_tb_equipment_status_CRUD":
	    	fmt.Println("上传设备状态")
	    case "bull.ResourcesHZ.SYN_tb_equipment_location":
	    	fmt.Println("上传gps位置")

	    case "bull.ResourcesHZ.SyN_medical_UpImage":// 上传医生头像
	    case "bull.ResourcesHZ.SYN_person_UpImage":// 上传居民头像
	    case "bull.ResourcesHZ.SYN_yhxd_CRUD":// 上传心电检测文件
	    default:


	    }
	}
	if m == "GET" {
		fmt.Println(m)
		fmt.Println(req.URL.Path)
	}
}

//curl -F "file=@register_device.bat" localhost:9090/upload -v
func Upload(res http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(32 << 20)
	file, handler, err := req.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Fprintf(res, "upload ok")
}


func TestRet(res http.ResponseWriter, req *http.Request) {
	retitem := RetItem{RET_CODE: "100100", SERVICE_CODE: "serviercode", RET_MSG: "ret msg"}
	syshead := SysHead{}
	syshead.RET = append(syshead.RET, retitem)
	syshead.RET_STATUS = "status"
	var ret Result
	ret.SYS_HEAD = syshead
	str, err := json.Marshal(&ret)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(str))
}

type Config struct {
	DB_IP string
	DB_PORT int
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME string
}

func ReadConfigFile() (*Config, error) {
	file, err := os.Open("app.conf")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	cfg := &Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func FileServlet(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	fmt.Println(req.URL.Path)
	fmt.Println(req.RemoteAddr)
	fmt.Println(req.RequestURI)
	req.ParseMultipartForm(32 << 20)
	file, handler, err := req.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println(handler.Filename)
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Fprintf(res, "upload ok")
}

type FileInfo struct {
	SYS_HEAD SysHead
	BODY FileInfoItem
}

type FileInfoItem struct {
	FileIds []string `json:"fileIds"`
}

//心电和胎监文件的上传，心电是两个文件，胎监是3个文件
func MultiFileServlet(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		req.ParseMultipartForm(32 << 20)	//在使用r.MultipartForm前必须先调用ParseMultipartForm方法，参数为最大缓存
		// fmt.Println(r.MultipartForm)
		// fmt.Println(r.MultipartReader())
		var fileinfo FileInfo 
		fii := FileInfoItem{}
		retitem := RetItem{RET_CODE: "000000", SERVICE_CODE: "serviercode", RET_MSG: "ret msg"}
		syshead := SysHead{}
		syshead.RET = append(syshead.RET, retitem)
		syshead.RET_STATUS = "S"
		
		if req.MultipartForm != nil && req.MultipartForm.File != nil {
			fhs := req.MultipartForm.File["file"]		//获取所有上传文件信息
			num := len(fhs)

			fmt.Printf("总文件数：%d 个文件", num)

			//循环对每个文件进行处理
			for _, fheader := range fhs {
				//获取文件名
				filename := fheader.Filename

				//结束文件
				file,err := fheader.Open()
				if err != nil {
					fmt.Println(err)
				}

				//保存文件
				defer file.Close()
				f, err := os.Create(filename)
				defer f.Close()
				io.Copy(f, file)
				fii.FileIds = append(fii.FileIds, filename)
				fmt.Println(filename)
				//获取文件状态信息
				//fstat,_ := f.Stat()

				//打印接收信息
				//fmt.Fprintf(w, "%s  NO.: %d  Size: %d KB  Name：%s\n", time.Now().Format("2006-01-02 15:04:05"), n, fstat.Size()/1024, filename)
				//fmt.Printf("%s  NO.: %d  Size: %d KB  Name：%s\n", time.Now().Format("2006-01-02 15:04:05"), n, fstat.Size()/1024, filename)

			}
		}
		fileinfo.BODY = fii
		fileinfo.SYS_HEAD = syshead
		ret, err := json.Marshal(fileinfo)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(ret))
		res.Write(ret)
	}
}

func main() {
	initlogger()

	var db_connector string
	cfg, err := ReadConfigFile()
	if err != nil {
		db_connector = "root:r00t@tcp(localhost:3306)/med_gwc?charset=utf8"
	}else{
		db_connector += cfg.DB_USERNAME
		db_connector += ":" + cfg.DB_PASSWORD
		db_connector += "@" + "tcp(" + cfg.DB_IP + ":" + strconv.Itoa(cfg.DB_PORT) + ")"
		db_connector += "/" + cfg.DB_NAME + "?charset=utf8"
	}
	fmt.Println(db_connector)
	logger.Println(db_connector)
	//db, _ := sql.Open("mysql", "root:r00t@tcp(localhost:3306)/med_gwc?charset=utf8")
	db, _ = sql.Open("mysql", db_connector)
	
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	http.HandleFunc("/", Home)
	http.HandleFunc("/serviceProxy/servlet/", AllServlet)
	http.HandleFunc("/medhcpfile/file/uploadFetalMonitoring.do", MultiFileServlet)
	http.HandleFunc("/upload", Upload)
	http.HandleFunc("/testret", TestRet)
	fs := http.FileServer(http.Dir("."))
	http.Handle("/test", http.StripPrefix("/test/", fs))
	logger.Println("server listen at", BIND_ADDR)
	log.Fatal(http.ListenAndServe(BIND_ADDR, nil))
}

func initlogger(){
    logfile, err := os.OpenFile("app.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0777)
    if err != nil {
        fmt.Printf("%s\r\n", err.Error())
        os.Exit(-1)
    }
    //defer logfile.Close()
    writers := []io.Writer{
        logfile,
        os.Stdout,
    }
    fileAndStdoutWriter := io.MultiWriter(writers...)
    //logger := log.New(fileAndStdoutWriter, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
    logger = log.New(fileAndStdoutWriter, "", log.Ldate|log.Ltime|log.Lshortfile)//Llongfile
    //logger.Println("hello")
    // logger.Fatal("test")
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



json={"data":{"MACHINE_ID":"D1a00010205","ORG_CODE":"123000","CREATETIME":"2017-08-19 15:19:07","CONTACT_PHONE":"12300001111","ORG_NAME":"机构名","PID":"D1a00010205123000","CONTACT":"请问"},"SERVICE_CODE":"bull.ResourcesHZ.SNY_tb_equipment_CRUD"}
{
    "data": {
        "MACHINE_ID": "D1a00010205",
        "ORG_CODE": "123000",
        "CREATETIME": "2017-08-19 15:19:07",
        "CONTACT_PHONE": "12300001111",
        "ORG_NAME": "机构名",
        "PID": "D1a00010205123000",
        "CONTACT": "请问"
    },
    "SERVICE_CODE": "bull.ResourcesHZ.SNY_tb_equipment_CRUD"
}



json={"pageSize":10,"esblover":false,"pageNumber":0,"ORG_CODE":"123000","CONSUMER_ID":"bbeeb31c1e7542a793203cc3bc376840","SERVICE_CODE":"bull.CloudPlatInterface.Get_tb_admindivision_cascade_select"}
{
    "pageSize": 10,
    "esblover": false,
    "pageNumber": 0,
    "ORG_CODE": "123000",
    "CONSUMER_ID": "bbeeb31c1e7542a793203cc3bc376840",
    "SERVICE_CODE": "bull.CloudPlatInterface.Get_tb_admindivision_cascade_select"
}



JSONArray data = response.getJSONObject("BODY").getJSONArray("data");
final String name = data.length() > 0 ? data.getJSONObject(0).getString("ORG_NAME") : "";
getActivity().runOnUiThread(new Runnable()
{
    
    @Override
    public void run()
    {
        orgNameEt.setText(name);
        orgNameEt.setSelection(orgNameEt.getText().length());
    }
});


JSONObject response = new JSONObject(responseinfo.result);
Log.e("yao", "设备注册response:" + response);

if (response.has("SYS_HEAD"))
{
    JSONObject json = response.getJSONObject("SYS_HEAD");
    if ("S".equals(json.getString("RET_STATUS")))
    {
        JSONObject body = response.getJSONObject("BODY");
        if ("S".equals(body.getString("flag")))
        {
            what = 0;
            device.setDeviceId(body.getString("org_code"));
            device.setManufacturer(body.getString("MANUFACTURER"));
        }
        else
        {
            what = 1;// 设备注册失败:机构编码或机构名称不存在
        }
    }
    else
    {
        what = 2;
    }
}


db, err := sql.Open("mysql", "root:r00t@tcp(localhost:3306)/test")
err := db.Ping()
if err != nil {
	fmt.Println(err)
}

var str string
err = db.QueryRow("select name from xxx where id = 1").Scan(&str)
if err != nil && err != sql.ErrNoRows {
	fmt.Println(err)
}
fmt.Println(str)

for rows.Next(){
	err = rows.Scan(&id, &str)
	if err != nil {
		fmt.Println(err)
	}
	//
}
if err = rows.Err(); err != nil {
	rows.Close()	
	fmt.Println(err)
}

--------------------------------------------
stmt, err := db.Prepare("select id, name from www where id = ?")
if err != nil {
	fmt.Println(err)
}
defer stmt.Close()

rows, err := stmt.Query(1)
if err != nil {
	fmt.Println(err)
}
defer rows.Close()
for rows.Next(){

}
---------------------------------------------------
stmt, err := db.Prepare("
	insert into xxx(name) values(?)")
if err != nil{
	fmt.Println(err)
}
defer stmt.Close()
res, err := stmt.Exec("hello, ddd")
if err != nil {
	fmt.Println(err)
}

lastId, err := res.LastInsertId()

rowCnt, err := res.RowsAffected()

------------------------------------------
//don't do this 
for i := 0; i < 20; i++{
	_, err := db.Query("delect form hello")
}

//use this instead
for i := 0; i < 20; i++{
	_, err := db.Exec("delete from hello")
}




DoctorManagementFragment
JSONObject request = new JSONObject();
try {
    if (noDoctor) {
        request.put("type", "1");
    } else {
        request.put("type", "0");
    }
    request.put("ORG_CODE", HealthBoxApplication.deviceBean.getOrgId());
    request.put("MACHINE_ID", HealthBoxApplication.deviceBean.getDeviceId());
    if (HealthBoxApplication.deviceBean.getDoctorVersion() != 0) {
        String date = DateUtil.format(HealthBoxApplication.deviceBean.getDoctorVersion(),
                "yyyy-MM-dd HH:mm:ss");
        // 有日期才加这个参数
        request.put("TIME", date);
    }
    request.put("SERVICE_CODE", "bull.ResourcesHZ.Down_tb_medical_technicians");
} catch (JSONException e) {
    e.printStackTrace();
}


PersonListFragment
JSONObject request = new JSONObject();
try {
    if (noPerson) {
        request.put("type", "1");
    } else {
        request.put("type", "0");
    }
    request.put("ORG_CODE", HealthBoxApplication.deviceBean.getOrgId());
    request.put("MACHINE_ID", HealthBoxApplication.deviceBean.getDeviceId());
    // request.put("ORG_CODE", "20151029");
    // request.put("MACHINE_ID", "100018");
    if (HealthBoxApplication.deviceBean.getPersonVersion() != 0) {
        String date = DateUtil.format(HealthBoxApplication.deviceBean.getPersonVersion(),
                "yyyy-MM-dd HH:mm:ss");
        // 有日期才加这个参数
        request.put("TIME", date);
    }
    request.put("SERVICE_CODE", "bull.ResourcesHZ.Down_mpi_personbasics");
} catch (JSONException e) {
    e.printStackTrace();
}


JSONObject req = new JSONObject();
try {
    req.put("SERVICE_CODE", "bull.MHPMobileHealth.File");
    req.put("bullfile_id", fileName);
    req.put("bullfile_name", "1.png");
} catch (JSONException e) {
    e.printStackTrace();

*/