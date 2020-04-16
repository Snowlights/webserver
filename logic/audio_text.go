package logic

import (
	"fmt"
	"github.com/Snowlights/pub/adapter"
	corpus "github.com/Snowlights/pub/grpc"
	"log"
	"net/http"
)

func AudioToText(w http.ResponseWriter, r *http.Request){
	data := ParseRequest(r)

	req := &corpus.TransAudioToTextReq{
	}

	for k,v := range data{
		switch k {
		case "audio":
			req.Audio = v.(string)
		case "cookie":
			req.Cookie = v.(string)
		}
	}
	if req.Audio == "" || req.Cookie == ""{
		w.Write([]byte("wrong input args"))
		return
	}

	res := adapter.TransAudioToText(r.Context(),req)
	if res.Errinfo != nil{
		log.Printf("code error %v\n",res.Errinfo.Msg)
	}

	dataShow := fmt.Sprintf("%v",res)
	w.Write([]byte(dataShow))
	w.WriteHeader(http.StatusOK)

}
