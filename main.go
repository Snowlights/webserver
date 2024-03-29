package webserver

import (
	"github.com/Snowlights/webserver/logic"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var server *http.Server


func main() {

	// 一个通知退出的chan
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.HandleFunc("/bye", sayBye)
	_ = logic.Register(mux)

	server = &http.Server{
		Addr:         ":9100",
		WriteTimeout: time.Second * 4,
		Handler:      mux,
	}

	go func() {
		// 接收退出信号
		<-quit
		if err := server.Close(); err != nil {
			log.Fatal("Close server:", err)
		}
	}()

	log.Println("Starting httpserver")
	err := server.ListenAndServe()
	if err != nil {
		// 正常退出
		if err == http.ErrServerClosed {
			log.Fatal("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected", err)
		}
	}
	log.Fatal("Server exited")

}

// 关闭http
func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye ,shutdown the server"))     // 没有输出
	err := server.Shutdown(nil)
	if err != nil {
		log.Printf("shutdown the server err")
	}
}

