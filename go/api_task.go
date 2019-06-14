/*
 * 挣闲钱
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"encoding/binary"
	"encoding/json"
	//"encoding/binary"
  	"fmt"
    "log"
	"net/http"
	"net/url"
    "strings"
	//"errors"
	"strconv"
    //"github.com/codegangsta/negroni"
	//"github.com/boltdb/bolt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//"io/ioutil"
	//"reflect"
)

type ErrorResponse struct {
    Error string `json:"error"`
}

func itob(v int) []byte {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(v))
    return b
}

//完成
func GetDelivery(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:HUANG@123@tcp(127.0.0.1:3306)/?charset=utf8")
	if err != nil {
			log.Fatal(err)
	}
	defer db.Close()

	taskId := strings.Split(r.URL.Path, "/")[3]
	_, err = strconv.Atoi(taskId)
	if err != nil {
		reponse := ErrorResponse{"Wrong TaskId"}
		JsonResponse(reponse, w, http.StatusNotFound)
		return
	}
	
	query, err := db.Query("select * from mytest.task where taskId=" + taskId)
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()

	v, err := getJSON(query)
	if err != nil {
		log.Fatal(err)
	}

	if string(v) == "[]" {
		reponse := ErrorResponse{"taskId Not Exists"}
		JsonResponse(reponse, w, http.StatusNotFound)
		return
	}
	v = v[1:len(v)-1]
	str := strings.Replace(string(v), "taskId\":\"", "taskId\":", -1)
	str = strings.Replace(str, "\",\"taskTitle", ",\"taskTitle", -1)
	str = strings.Replace(str, "userId\":\"", "userId\":", -1)
	str = strings.Replace(str, "\"}", "}", -1)
	v = []byte(str)
	fmt.Printf(string(v))
	
	var task Task
	
	_ = json.Unmarshal(v, &task)


	query, err = db.Query("select * from mytest.delivery where deliveryId=" + taskId)
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()

	v, err = getJSON(query)
	if err != nil {
		log.Fatal(err)
	}
	if string(v) == "[]" {
		reponse := ErrorResponse{"Not a DeliveryTask"}
		JsonResponse(reponse, w, http.StatusNotFound)
		return
	}
	v = v[1:len(v)-1]
	str = strings.Replace(string(v), "deliveryId\":\"", "deliveryId\":", -1)
	str = strings.Replace(str, "\"}", "}", -1)
	v = []byte(str)
	fmt.Printf(string(v))
	
	var delivery Delivery
	
	_ = json.Unmarshal(v, &delivery)

	deliveryTask := DeliveryTask{
		Task: task,
		Delivery: delivery,
	}

	JsonResponse(deliveryTask, w, http.StatusOK)

}
//完成
func GetQuestionare(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:HUANG@123@tcp(127.0.0.1:3306)/?charset=utf8")
	if err != nil {
			log.Fatal(err)
	}
	defer db.Close()

	taskId := strings.Split(r.URL.Path, "/")[3]
	_, err = strconv.Atoi(taskId)
	if err != nil {
		reponse := ErrorResponse{"Wrong TaskId"}
		JsonResponse(reponse, w, http.StatusNotFound)
		return
	}
	
	query, err := db.Query("select * from mytest.task where taskId=" + taskId)
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()

	v, err := getJSON(query)
	if err != nil {
		log.Fatal(err)
	}

	if string(v) == "[]" {
		reponse := ErrorResponse{"taskId Not Exists"}
		JsonResponse(reponse, w, http.StatusNotFound)
		return
	}
	v = v[1:len(v)-1]
	
	str := strings.Replace(string(v), "taskId\":\"", "taskId\":", -1)
	str = strings.Replace(str, "\",\"taskTitle", ",\"taskTitle", -1)
	str = strings.Replace(str, "userId\":\"", "userId\":", -1)
	str = strings.Replace(str, "\"}", "}", -1)
	v = []byte(str)
	fmt.Printf(string(v))
	
	var task Task
	
	_ = json.Unmarshal(v, &task)

	query, err = db.Query("select * from mytest.questionare where queryid=" + taskId)
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()

	v, err = getJSON(query)
	if err != nil {
		log.Fatal(err)
	}
	if string(v) == "[]" {
		reponse := ErrorResponse{"Not a QuestionareTask"}
		JsonResponse(reponse, w, http.StatusNotFound)
		return
	}
	v = v[1:len(v)-1]
	str = strings.Replace(string(v), "queryid\":\"", "queryid\":", -1)
	str = strings.Replace(str, "\"}", "}", -1)
	v = []byte(str)
	fmt.Printf(string(v))
	
	var questionare Questionare
	
	_ = json.Unmarshal(v, &questionare)

	questionareTask := QuestionareTask{
		Task: task,
		Questionare: questionare,
	}


	JsonResponse(questionareTask, w, http.StatusOK)
}
//完成
func QAcceptPage(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:HUANG@123@tcp(127.0.0.1:3306)/?charset=utf8")
	if err != nil {
			log.Fatal(err)
	}
	defer db.Close()

	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Fatal(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	page := m["page"][0]
	userId := m["userId"][0]
	IdIndex, err:= strconv.Atoi(page)
	p := "10"
	if IdIndex == 0 {
		p = "3"
	} else {
		IdIndex = (IdIndex - 1)* 10
	}
	Id := strconv.Itoa(IdIndex)

	query, err := db.Query("select * from mytest.usertask where userId = " + userId + " limit " + Id + "," + p)
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()

	v, err := getJSON(query)
	if err != nil {
		log.Fatal(err)
	}

	if string(v) == "[]" {
		reponse := ErrorResponse{"Page is out of index"}
		JsonResponse(reponse, w, http.StatusNotFound)
		return
	}

	var tasks Tasks
	v = []byte("{\"contents\":" + string(v) + "}")
	str := strings.Replace(string(v), "taskId\":\"", "taskId\":", -1)
	str = strings.Replace(str, "\",\"", ",\"", -1)
	str = strings.Replace(str, "userId\":\"", "userId\":", -1)
	str = strings.Replace(str, "\"}", "}", -1)
	v = []byte(str)

	_ = json.Unmarshal(v, &tasks)
	
	//fmt.Printf(string(tasks.Contents[0].TaskId))

	var acceptasks Tasks

	for i, task := range(tasks.Contents) {
		fmt.Printf(string(i))
		query_, err_ := db.Query("select * from mytest.task where taskId = " + strconv.Itoa(task.TaskId))
		if err_ != nil {
			log.Fatal(err_)
		}
		defer query_.Close()

		v_, err_ := getJSON(query_)
		if err_ != nil {
			log.Fatal(err_)
		}

		v_ = v_[1:len(v_)-1]
	
		str_ := strings.Replace(string(v_), "taskId\":\"", "taskId\":", -1)
		str_ = strings.Replace(str_, "\",\"taskTitle", ",\"taskTitle", -1)
		str_ = strings.Replace(str_, "userId\":\"", "userId\":", -1)
		str_ = strings.Replace(str_, "\"}", "}", -1)
		v_ = []byte(str_)
		fmt.Printf(string(v_))
		
		var acceptask Task
		
		_ = json.Unmarshal(v_, &acceptask)
	    acceptasks.Contents = append(acceptasks.Contents, acceptask)
	}

	JsonResponse(acceptasks, w, http.StatusOK)
}
//完成
func QPublishPage(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:HUANG@123@tcp(127.0.0.1:3306)/?charset=utf8")
	if err != nil {
			log.Fatal(err)
	}
	defer db.Close()

	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Fatal(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	page := m["page"][0]
	userId := m["userId"][0]
	IdIndex, err:= strconv.Atoi(page)
	p := "10"
	if IdIndex == 0 {
		p = "3"
	} else {
		IdIndex = (IdIndex - 1)* 10
	}
	Id := strconv.Itoa(IdIndex)

	query, err := db.Query("select * from mytest.task where userId = " + userId + " limit " + Id + "," + p)
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()

	v, err := getJSON(query)
	if err != nil {
		log.Fatal(err)
	}

	if string(v) == "[]" {
		reponse := ErrorResponse{"Page is out of index"}
		JsonResponse(reponse, w, http.StatusNotFound)
		return
	}

	var tasks Tasks
	v = []byte("{\"contents\":" + string(v) + "}")
	str := strings.Replace(string(v), "taskId\":\"", "taskId\":", -1)
	str = strings.Replace(str, "\",\"taskTitle", ",\"taskTitle", -1)
	str = strings.Replace(str, "userId\":\"", "userId\":", -1)
	str = strings.Replace(str, "\"}", "}", -1)
	v = []byte(str)

	_ = json.Unmarshal(v, &tasks)

	JsonResponse(tasks, w, http.StatusOK)
}
//完成
func QueryPageD(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:HUANG@123@tcp(127.0.0.1:3306)/?charset=utf8")
	if err != nil {
			log.Fatal(err)
	}
	defer db.Close()

	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Fatal(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	page := m["page"][0]
	userId := m["userId"][0]
	IdIndex, err:= strconv.Atoi(page)
	p := 10
	if IdIndex == 0 {
		p = 3
	} else {
		IdIndex = (IdIndex - 1)* 10
	}
	Id := strconv.Itoa(IdIndex)

	var tasks Tasks
	var hasCount bool
	hasCount = false
	for p > 0 {
		Id = strconv.Itoa(IdIndex)
		query, err := db.Query("select * from mytest.task where userId <> " + userId + " and taskType <> 'questionare' limit " + Id + ", 1")
		if err != nil {
			log.Fatal(err)
		}
		defer query.Close()

		v, err := getJSON(query)
		if err != nil {
			log.Fatal(err)
		}
	
		if string(v) == "[]"  && hasCount == false{
			reponse := ErrorResponse{"Page is out of index"}
			JsonResponse(reponse, w, http.StatusNotFound)
			return
		} else if string(v) == "[]" {
			break
		}
		hasCount = true
		v = v[1:len(v)-1]
		str := strings.Replace(string(v), "taskId\":\"", "taskId\":", -1)
		str = strings.Replace(str, "\",\"taskTitle", ",\"taskTitle", -1)
		str = strings.Replace(str, "userId\":\"", "userId\":", -1)
		str = strings.Replace(str, "\"}", "}", -1)
		v = []byte(str)
		fmt.Printf(string(v))
		var task Task
		_ = json.Unmarshal(v, &task)

		query, err = db.Query("select * from mytest.usertask where userId = " + userId + " and taskId = " + strconv.Itoa(task.TaskId))
		if err != nil {
			log.Fatal(err)
		}
		defer query.Close()

		v, err = getJSON(query)
		if err != nil {
			log.Fatal(err)
		}
	
		if string(v) == "[]" {
			tasks.Contents = append(tasks.Contents, task)
			p = p - 1	
		}
		IdIndex = IdIndex + 1
	}

	JsonResponse(tasks, w, http.StatusOK)
}
//完成
func QueryPageQ(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:HUANG@123@tcp(127.0.0.1:3306)/?charset=utf8")
	if err != nil {
			log.Fatal(err)
	}
	defer db.Close()

	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Fatal(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	page := m["page"][0]
	userId := m["userId"][0]
	IdIndex, err:= strconv.Atoi(page)
	p := 10
	if IdIndex == 0 {
		p = 3
	} else {
		IdIndex = (IdIndex - 1)* 10
	}
	Id := strconv.Itoa(IdIndex)

	var tasks Tasks
	var hasCount bool
	hasCount = false
	for p > 0 {
		Id = strconv.Itoa(IdIndex)
		query, err := db.Query("select * from mytest.task where userId <> " + userId + " and taskType <> 'delivery' limit " + Id + ", 1")
		if err != nil {
			log.Fatal(err)
		}
		defer query.Close()

		v, err := getJSON(query)
		if err != nil {
			log.Fatal(err)
		}
	
		if string(v) == "[]"  && hasCount == false{
			reponse := ErrorResponse{"Page is out of index"}
			JsonResponse(reponse, w, http.StatusNotFound)
			return
		} else if string(v) == "[]" {
			break
		}
		hasCount = true
		v = v[1:len(v)-1]
		str := strings.Replace(string(v), "taskId\":\"", "taskId\":", -1)
		str = strings.Replace(str, "\",\"taskTitle", ",\"taskTitle", -1)
		str = strings.Replace(str, "userId\":\"", "userId\":", -1)
		str = strings.Replace(str, "\"}", "}", -1)
		v = []byte(str)
		fmt.Printf(string(v))
		var task Task
		_ = json.Unmarshal(v, &task)

		query, err = db.Query("select * from mytest.usertask where userId = " + userId + " and taskId = " + strconv.Itoa(task.TaskId))
		if err != nil {
			log.Fatal(err)
		}
		defer query.Close()

		v, err = getJSON(query)
		if err != nil {
			log.Fatal(err)
		}
	
		if string(v) == "[]" {
			tasks.Contents = append(tasks.Contents, task)
			p = p - 1	
		}
		IdIndex = IdIndex + 1
	}

	JsonResponse(tasks, w, http.StatusOK)
}
//完成
func QueryTitle(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:HUANG@123@tcp(127.0.0.1:3306)/?charset=utf8")
	if err != nil {
			log.Fatal(err)
	}
	defer db.Close()

	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Fatal(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	title := m["title"][0]
	userId := m["userId"][0]
	//IdIndex, err:= strconv.Atoi(page)
	//IdIndex = (IdIndex - 1)* 10
	//Id := strconv.Itoa(IdIndex)

	query, err := db.Query("select * from mytest.task where userId <> " + userId + " and taskTitle = '" + title + "'")
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()

	v, err := getJSON(query)
	if err != nil {
		log.Fatal(err)
	}

	if string(v) == "[]" {
		reponse := ErrorResponse{"Task is not found"}
		JsonResponse(reponse, w, http.StatusNotFound)
		return
	}

	var tasks Tasks
	v = []byte("{\"contents\":" + string(v) + "}")
	str := strings.Replace(string(v), "taskId\":\"", "taskId\":", -1)
	str = strings.Replace(str, "\",\"taskTitle", ",\"taskTitle", -1)
	str = strings.Replace(str, "userId\":\"", "userId\":", -1)
	str = strings.Replace(str, "\"}", "}", -1)
	v = []byte(str)

	_ = json.Unmarshal(v, &tasks)

	var finalTasks Tasks
	for _, task := range(tasks.Contents) {
		query_, err_ := db.Query("select * from mytest.usertask where userId = " + userId + " and taskId = " + strconv.Itoa(task.TaskId))
		if err_ != nil {
			log.Fatal(err_)
		}
		defer query_.Close()
	
		v_, err_ := getJSON(query_)
		if err_ != nil {
			log.Fatal(err_)
		}
	
		if string(v_) == "[]" {
			finalTasks.Contents = append(finalTasks.Contents, task)
		}

		if len(finalTasks.Contents) == 0 {
			reponse := ErrorResponse{"Task is not found"}
			JsonResponse(reponse, w, http.StatusNotFound)
			return
		}
	} 

	JsonResponse(finalTasks, w, http.StatusOK)
}

func getJSON(rows *sql.Rows) ([]byte, error) {
	columns, err := rows.Columns()
	if err != nil {
		return []byte(""), err
	}
	  count := len(columns)
  
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	  }
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return []byte(""), err
	  }
	return jsonData, nil 
  }

  func JsonResponse(response interface{}, w http.ResponseWriter, code int) {
    json, err := json.Marshal(response)
    if err != nil {
        log.Fatal(err)
        return
    }


    w.Header().Set("Access-Control-Allow-Methods","PUT,POST,GET,DELETE,OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,Content-Type,Authorization")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
    w.Write(json)
}