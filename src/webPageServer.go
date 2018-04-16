package main

import (
	"net/http"
	"strconv"
)

func startHttpServer(u *url){
	http.Handle("/", http.FileServer(http.Dir("www/")))
	http.ListenAndServe(":"+strconv.Itoa(u.port) , nil)
}
