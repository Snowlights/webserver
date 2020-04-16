package logic

import (
	"fmt"
	"github.com/Snowlights/pub/adapter"
	corpus "github.com/Snowlights/pub/grpc"
	"log"
	"net/http"
)

func  SendMessage(w http.ResponseWriter, r *http.Request) {
	data := ParseRequest(r)

	req := &corpus.SendMessageReq{
	}

	for k,v := range data{
		switch k {
		case "phone":
			req.Phone = v.(string)
		}
	}
	if req.Phone == "" || len(req.Phone) != 11{
		w.Write([]byte("wrong input args"))
		return
	}

	fmt.Printf("%v",req)
	res := adapter.SendMessage(r.Context(),req)

	if res.Errinfo != nil{
		log.Printf("code error %v\n",res.Errinfo.Msg)
	}
	fmt.Printf("%v",res.Data.Code)
	dataShow := fmt.Sprintf("%v",res)
	w.Write([]byte(dataShow))
	w.WriteHeader(http.StatusOK)
	return
}
