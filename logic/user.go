package logic

import (
	"fmt"
	"github.com/Snowlights/pub/adapter"
	corpus "github.com/Snowlights/pub/grpc"
	"log"
	"net/http"
)

func UserLogin(w http.ResponseWriter, r *http.Request){
	data := ParseRequest(r)

	req := &corpus.LoginUserReq{}
	for k,v := range data{
		switch k {
		case "e_mail":
			req.EMail = v.(string)
		case "user_password":
			req.UserPassword = v.(string)
		case "phone":
			req.Phone = v.(string)
		case "code":
			req.Code = v.(string)
		}
	}

	if req.EMail == "" || req.Phone == ""{
		w.Write([]byte("wrong input args"))
		return
	}

	res := adapter.LoginUser(r.Context(),req)
	if res.Errinfo != nil{
		log.Printf("code error %v\n",res.Errinfo.Msg)
	}

	dataShow := fmt.Sprintf("%v",res)
	w.Write([]byte(dataShow))
	w.WriteHeader(http.StatusOK)
	return
}

func UserLogOut(w http.ResponseWriter, r *http.Request){
	data := ParseRequest(r)

	req := &corpus.LogoutUserInfoReq{
	}

	for k,v := range data{
		switch k {
		case "cookie":
			req.Cookie = v.(string)
		}
	}
	if req.Cookie == ""{
		w.Write([]byte("wrong input args"))
		return
	}

	res := adapter.LogoutUserInfo(r.Context(),req)
	if res.Errinfo != nil{
		log.Printf("code error %v\n",res.Errinfo.Msg)
	}

	dataShow := fmt.Sprintf("%v",res)
	w.Write([]byte(dataShow))
	w.WriteHeader(http.StatusOK)
	return
}

func UserUpdateInfo(w http.ResponseWriter, r *http.Request){
	data := ParseRequest(r)

	req := &corpus.UpdateUserInfoReq{
	}

	for k,v := range data{
		switch k {
		case "id":
			req.UserId = v.(int64)
		case "user_name":
			req.UserName = v.(string)
		case "e_mail":
			req.EMail = v.(string)
		case "password":
			req.Password = v.(string)
		case "description":
			req.Description = v.(string)
		case "cookie":
			req.Cookie = v.(string)
		}
	}

	if req.UserId == 0 || req.Cookie == ""{
		w.Write([]byte("wrong input args"))
		return
	}

	res := adapter.UpdateUserInfo(r.Context(),req)
	if res.Errinfo != nil{
		log.Printf("code error %v\n",res.Errinfo.Msg)
	}

	dataShow := fmt.Sprintf("%v",res)
	w.Write([]byte(dataShow))
	w.WriteHeader(http.StatusOK)
	return
}

func UserUpdatePhone(w http.ResponseWriter, r *http.Request){
	data := ParseRequest(r)

	req := &corpus.UpdateUserPhoneReq{
	}

	for k,v := range data{
		switch k {
		case "phone":
			req.Phone = v.(string)
		case "code":
			req.Code = v.(string)
		case "cookie":
			req.Cookie = v.(string)
		}
	}

	if req.Phone == "" || req.Code == "" || req.Cookie == ""{
		w.Write([]byte("wrong input args"))
		return
	}


	res := adapter.UpdateUserPhone(r.Context(),req)
	if res.Errinfo != nil{
		log.Printf("code error %v\n",res.Errinfo.Msg)
	}

	dataShow := fmt.Sprintf("%v",res)
	w.Write([]byte(dataShow))
	w.WriteHeader(http.StatusOK)
	return
}

func UserDelInfo(w http.ResponseWriter, r *http.Request){
	data := ParseRequest(r)

	req := &corpus.DelUserInfoReq{
	}

	for k,v := range data{
		switch k {
		case "id":
			req.UserId = v.(int64)
		case "cookie":
			req.Cookie = v.(string)
		}
	}

	if req.UserId == 0 || req.Cookie == ""{
		w.Write([]byte("wrong input args"))
		return
	}

	res := adapter.DelUserInfo(r.Context(),req)
	if res.Errinfo != nil{
		log.Printf("code error %v\n",res.Errinfo.Msg)
	}

	dataShow := fmt.Sprintf("%v",res)
	w.Write([]byte(dataShow))
	w.WriteHeader(http.StatusOK)
	return
}

func UserListInfo(w http.ResponseWriter, r *http.Request){
	data := ParseRequest(r)

	req := &corpus.ListUserInfoReq{
	}

	for k,v := range data{
		switch k {
		case "offset":
			req.Offset = v.(int64)
		case "limit":
			req.Limit = v.(int64)
		case "cookie":
			req.Cookie = v.(string)
		case "e_mail":
			req.EMail = v.(string)
		case "user_name":
			req.UserName = v.(string)
		case "phone":
			req.Phone = v.(string)
		}
	}
	if req.Limit == 0 || req.Cookie == "" {
		w.Write([]byte("wrong input args"))
		return
	}

	res := adapter.ListUserInfo(r.Context(),req)
	if res.Errinfo != nil{
		log.Printf("code error %v\n",res.Errinfo.Msg)
	}

	dataShow := fmt.Sprintf("%v",res)
	w.Write([]byte(dataShow))
	w.WriteHeader(http.StatusOK)
	return
}

func ListImageByUserCookde(w http.ResponseWriter, r *http.Request){
	data := ParseRequest(r)

	req := &corpus.ListImageByUserCookieReq{
	}

	for k,v := range data{
		switch k {
		case "limit":
			req.Limit = v.(int64)
		case "offset":
			req.Offset = v.(int64)
		case "cookie":
			req.Cookie = v.(string)
		}
	}

	if req.Limit == 0 || req.Cookie == ""{
		w.Write([]byte("wrong input args"))
		return
	}

	res := adapter.ListImageByUserCookie(r.Context(),req)
	if res.Errinfo != nil{
		log.Printf("code error %v\n",res.Errinfo.Msg)
	}

	dataShow := fmt.Sprintf("%v",res)
	w.Write([]byte(dataShow))
	w.WriteHeader(http.StatusOK)
	return
}