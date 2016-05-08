package main

import (
	"fmt"
	"github.com/EverythingMe/go-disque/disque"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/skyrocknroll/disque-dashborad/utils"
	"html/template"
	"io"
	"net/http"
)

var disquePool *disque.Pool
var t *template.Template

func init() {
	var err error
	disquePool = disque.NewPool(disque.DialFunc(dial), "127.0.0.1:7711")
	t, err = template.ParseGlob("assets/*.html")
	if err != nil {
		fmt.Println(err.Error())
	}
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/node-info", nodeInfo)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.Handle("/", r)
	fmt.Println("Listening on 0.0.0.0:9090")
	http.ListenAndServe("0.0.0.0:9090", r)

}
func dial(addr string) (redis.Conn, error) {
	return redis.Dial("tcp", addr)
}

func index(w http.ResponseWriter, r *http.Request) {

	client, err := disquePool.Get()

	defer client.Close()
	hello, err := client.Hello()
	fmt.Println(hello.NodeId)
	if err != nil {
		io.WriteString(w, err.Error())
	}

	if err != nil {
		io.WriteString(w, err.Error())
	}
	//tdata := utils.TemplateMeta{
	//	"Home",
	//	hello,
	//
	//}
	t.ExecuteTemplate(w, "Dashboard", hello)
	//t.ExecuteTemplate(w, "layout", hello)
}

func nodeInfo(w http.ResponseWriter, r *http.Request) {
	addr := r.URL.Query()["addr"][0]
	client, err := redis.Dial("tcp", addr)
	defer client.Close()
	data, err := redis.String(client.Do("INFO"))
	if err != nil {
		fmt.Println(err.Error())
	}
	//t, err := template.ParseGlob("assets/*.html")
	if err != nil {
		io.WriteString(w, err.Error())
	}
	parsedData := utils.ParseInfoCommandResponse(data)

	//io.WriteString(w, data)
	t.ExecuteTemplate(w, "nodeInfo", parsedData)

}

func queueInfo(w http.ResponseWriter, r *http.Request) {

}
