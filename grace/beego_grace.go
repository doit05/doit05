package main

import (
	"github.com/astaxie/beego/grace"
	"log"
	"net/http"
	"os"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello WORLD!"))
	w.Write([]byte("ospid:" + strconv.Itoa(os.Getpid())))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler)

	err := grace.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Println(err)
	}
	log.Println("Server on 8080 stopped")
	os.Exit(0)
}
