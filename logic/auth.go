package logic

import (
	"fmt"
	"github.com/Snowlights/pub/adapter"
	corpus "github.com/Snowlights/pub/grpc"
	"log"
	"net/http"
)

func AddAuth(w http.ResponseWriter, r *http.Request){
	data := ParseRequest(r)

	req := &corpus.AddAuthReq{
	}

	for k,v := range data{
		switch k {
		case "auth_code":
			req.AuthCode = v.(string)
		case "auth_description":
			req.AuthDescription = v.(string)
		case "service_name":
			req.ServiceName = v.(string)
		case "cookie":
			req.Cookie = v.(string)
		}
	}
	if req.AuthCode == "" || req.Cookie == ""{
		w.Write([]byte("wrong input args"))
		return
	}

	res := adapter.AddAuth(r.Context(),req)
	if res.Errinfo != nil{
		log.Printf("code error %v\n",res.Errinfo.Msg)
	}

	dataShow := fmt.Sprintf("%v",res)
	w.Write([]byte(dataShow))
	w.WriteHeader(http.StatusOK)
}

func DelAuth(w http.ResponseWriter, r *http.Request){
	data := ParseRequest(r)

	req := &corpus.DelAuthReq{
	}

	for k,v := range data{
		switch k {
		case "id":
			req.Id = v.(int64)
		case "auth_code":
			req.AuthCode = v.(string)
		case "cookie":
			req.Cookie = v.(string)
		}
	}

	if (req.AuthCode == "" && req.Id == 0) || req.Cookie == ""{
		w.Write([]byte("wrong input args"))
		return
	}

	res := adapter.DelAuth(r.Context(),req)
	if res.Errinfo != nil{
		log.Printf("code error %v\n",res.Errinfo.Msg)
	}

	dataShow := fmt.Sprintf("%v",res)
	w.Write([]byte(dataShow))
	w.WriteHeader(http.StatusOK)
}

func UpdateAuth(w http.ResponseWriter, r *http.Request){
	data := ParseRequest(r)

	req := &corpus.UpdateAuthReq{
	}

	for k,v := range data{
		switch k {
		case "id":
			req.Id = v.(int64)
		case "auth_code":
			req.AuthCode = v.(string)
		case "auth_description":
			req.AuthDescription = v.(string)
		case "service_name":
			req.ServiceName = v.(string)
		case "cookie":
			req.Cookie = v.(string)
		}
	}

	if (req.Id == 0 || req.AuthCode == "") || req.Cookie == ""{
		w.Write([]byte("wrong input args"))
		return
	}

	res := adapter.UpdateAuth(r.Context(),req)
	if res.Errinfo != nil{
		log.Printf("code error %v\n",res.Errinfo.Msg)
	}

	dataShow := fmt.Sprintf("%v",res)
	w.Write([]byte(dataShow))
	w.WriteHeader(http.StatusOK)
}

func ListAuth(w http.ResponseWriter, r *http.Request){
	data := ParseRequest(r)

	req := &corpus.ListAuthReq{
	}

	for k,v := range data{
		switch k {
		case "offset":
			req.Offset = v.(int64)
		case "limit":
			req.Limit = v.(int64)
		}
	}

	if req.Limit == 0{
		w.Write([]byte("wrong input args"))
		return
	}

	res := adapter.ListAuth(r.Context(),req)
	if res.Errinfo != nil{
		log.Printf("code error %v\n",res.Errinfo.Msg)
	}

	dataShow := fmt.Sprintf("%v",res)
	w.Write([]byte(dataShow))
	w.WriteHeader(http.StatusOK)
}

func AddUserAuth(w http.ResponseWriter, r *http.Request){
	data := ParseRequest(r)

	req := &corpus.AddUserAuthReq{
	}

	for k,v := range data{
		switch k {
		case "id":
			req.UserId = v.(int64)
		case "auth_code":
			req.AuthCode = v.(string)
		case "cookie":
			req.Cookie = v.(string)
		}
	}

	if req.UserId == 0 || req.AuthCode == "" || req.Cookie == ""{
		w.Write([]byte("wrong input args"))
		return
	}

	res := adapter.AddUserAuth(r.Context(),req)
	if res.Errinfo != nil{
		log.Printf("code error %v\n",res.Errinfo.Msg)
	}

	dataShow := fmt.Sprintf("%v",res)
	w.Write([]byte(dataShow))
	w.WriteHeader(http.StatusOK)
}

func DelUserAuth(w http.ResponseWriter, r *http.Request){
	data := ParseRequest(r)

	req := &corpus.DelUserAuthReq{}

	for k,v := range data{
		switch k {
		case "id":
			req.UserId = v.(int64)
		case "auth_code":
			req.AuthCode = v.(string)
		case "cookie":
			req.Cookie = v.(string)
		}
	}

	if req.UserId == 0 || req.AuthCode == "" || req.Cookie == ""{
		w.Write([]byte("wrong input args"))
		return
	}

	res := adapter.DelUserAuth(r.Context(),req)
	if res.Errinfo != nil{
		log.Printf("code error %v\n",res.Errinfo.Msg)
	}

	dataShow := fmt.Sprintf("%v",res)
	w.Write([]byte(dataShow))
	w.WriteHeader(http.StatusOK)
}

func ListUserAuth(w http.ResponseWriter, r *http.Request){
	data := ParseRequest(r)

	req := &corpus.ListUserAuthReq{
	}

	for k,v := range data{
		switch k {
		case "user_id":
			req.UserId = v.(int64)
		case "limit":
			req.Limit = v.(int64)
		case "offset":
			req.Offset = v.(int64)
		case "cookie":
			req.Cookie = v.(string)
		}
	}

	if req.UserId == 0 || req.Limit == 0 || req.Cookie == ""{
		w.Write([]byte("wrong input args"))
		return
	}

	res := adapter.ListUserAuth(r.Context(),req)
	if res.Errinfo != nil{
		log.Printf("code error %v\n",res.Errinfo.Msg)
	}

	dataShow := fmt.Sprintf("%v",res)
	w.Write([]byte(dataShow))
	w.WriteHeader(http.StatusOK)
}