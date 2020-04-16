package logic

import (
	"fmt"
	"github.com/Snowlights/pub/adapter"
	corpus "github.com/Snowlights/pub/grpc"
	"log"
	"net/http"
)

func AddAdminUser(w http.ResponseWriter, r *http.Request){
	data := ParseRequest(r)
	req := &corpus.AddAdminUserReq{}

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

	res := adapter.AddAdminUser(r.Context(),req)
	if res.Errinfo != nil{
		log.Printf("code error %v\n",res.Errinfo.Msg)
	}

	dataShow := fmt.Sprintf("%v",res)
	w.Write([]byte(dataShow))
	w.WriteHeader(http.StatusOK)
}

func DelAdminUser(w http.ResponseWriter, r *http.Request){
	data := ParseRequest(r)

	req := &corpus.DelAdminUserReq{}

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

	res := adapter.DelAdminUser(r.Context(),req)
	if res.Errinfo != nil{
		log.Printf("code error %v\n",res.Errinfo.Msg)
	}

	dataShow := fmt.Sprintf("%v",res)
	w.Write([]byte(dataShow))
	w.WriteHeader(http.StatusOK)
}

func ListAdminUser(w http.ResponseWriter, r *http.Request){
	data := ParseRequest(r)

	req := &corpus.ListAdminUserReq{
	}
	for k,v := range data{
		switch k {
		case "offset":
			req.Offset = v.(int64)
		case "limit":
			req.Limit = v.(int64)
		case "cookie":
			req.Cookie = v.(string)
		}
	}
	if req.Limit == 0 ||req.Cookie == ""{
		w.Write([]byte("wrong input args"))
		return
	}
	res := adapter.ListAdminUser(r.Context(),req)
	if res.Errinfo != nil{
		log.Printf("code error %v\n",res.Errinfo.Msg)
	}

	dataShow := fmt.Sprintf("%v",res)
	w.Write([]byte(dataShow))
	w.WriteHeader(http.StatusOK)
}