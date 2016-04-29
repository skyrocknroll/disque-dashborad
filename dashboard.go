package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"html/template"
	"io"
	"github.com/skyrocknroll/disque-dashborad/disque"
)

func main() {
	r := mux.NewRouter()
	assets := http.StripPrefix("/", http.FileServer(http.Dir("assets/")))
	r.HandleFunc("/", index)
	r.Handle("/assets", assets)
	fmt.Println("Listening on 0.0.0.0:9090")
	http.ListenAndServe("0.0.0.0:9090", r)

}
func index(w http.ResponseWriter, r *http.Request) {
	pool := disque.GetPool()
	client, err := pool.Get()
	defer client.Close()
	hello, err := client.Hello()
	fmt.Println(hello.NodeId)
	if err != nil {
		io.WriteString(w, err.Error())
	}
	t, err := template.ParseFiles("assets/index.html")
	if err != nil {
		io.WriteString(w, err.Error())
	}

	t.Execute(w, hello)
}
