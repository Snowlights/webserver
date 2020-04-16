package logic

import (
	"fmt"
	"github.com/Snowlights/pub/adapter"
	corpus "github.com/Snowlights/pub/grpc"
	"log"
	"net/http"
)

func Evaluation(w http.ResponseWriter, r *http.Request){
	data := ParseRequest(r)

	req := &corpus.EvaluationReq{
	}

	for k,v :=range data{
		switch k {
		case "audio":
			req.Audio = v.(string)
		case "text":
			req.Text = v.(string)
		case "cookie":
			req.Cookie = v.(string)
		}
	}
	if req.Audio == "" || req.Text == "" || req.Cookie == ""{
		w.Write([]byte("wrong input args"))
		return
	}

	res := adapter.Evaluation(r.Context(),req)
	if res.Errinfo != nil{
		log.Printf("code error %v\n",res.Errinfo.Msg)
	}

	dataShow := fmt.Sprintf("%v",res)
	w.Write([]byte(dataShow))
	w.WriteHeader(http.StatusOK)
}