package logic

import (
	"fmt"
	"log"
	"net/http"
)

type MyHandler struct{}

func ParseRequest(r *http.Request) map[string]interface{}{
	var data = make(map[string]interface{})
	err := r.ParseForm()
	if err != nil{
		log.Fatal(err)
	}

	for k,v := range r.Form {
		data[k] = v[0]
	}

	return data
}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%v\n",r.Form)
	for k,v := range r.Form {
		fmt.Printf("%v %v\n", k, v)
	}
	w.Write([]byte("this is version 3"))

}

func Register(mux *http.ServeMux) bool{
	mux.Handle("/", &MyHandler{})
	mux.HandleFunc("/sendmessage/",SendMessage)

	mux.HandleFunc("/user/login/",UserLogin)
	mux.HandleFunc("/user/loginout/",UserLogOut)
	mux.HandleFunc("/user/list/",UserListInfo)
	mux.HandleFunc("/user/del/",UserDelInfo)
	mux.HandleFunc("/user/update/phone/",UserUpdatePhone)
	mux.HandleFunc("/user/update/",UserUpdateInfo)

	mux.HandleFunc("/admin/user/add/",AddAdminUser)
	mux.HandleFunc("/admin/user/del/",DelAdminUser)
	mux.HandleFunc("/admin/user/list/",ListAdminUser)

	mux.HandleFunc("/auth/add/",AddAuth)
	mux.HandleFunc("/auth/del/",DelAuth)
	mux.HandleFunc("/auth/update/",UpdateAuth)
	mux.HandleFunc("/auth/list/",ListAuth)
	mux.HandleFunc("/auth/user/add/",AddUserAuth)
	mux.HandleFunc("/auth/user/del/",DelUserAuth)
	mux.HandleFunc("/auth/user/list/",ListUserAuth)

	mux.HandleFunc("/keyword/",KeyWord)
	mux.HandleFunc("/evaluation/",Evaluation)
	mux.HandleFunc("/recognize/image/",RecognizePicture)
	mux.HandleFunc("/recognize/audio2text/",AudioToText)
	mux.HandleFunc("/recognize/age/",RecognizeAge)

	mux.HandleFunc("/user/picture/list/",ListImageByUserCookde)

	return true
}