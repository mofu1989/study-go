package main

import (
	"net/http"
	"io"
	"time"
)

func main() {
	http.HandleFunc("/",hello)
	http.HandleFunc("/user/info",getUsername)
	http.ListenAndServe(":8080",nil)
}

func hello(w http.ResponseWriter, r *http.Request){
	cookie := http.Cookie{Name:"username",Value:"liuyh", Expires:time.Now().AddDate(1,0,0)}
	http.SetCookie(w,&cookie)
	io.WriteString(w,"hello world")
}

func getUsername(w http.ResponseWriter, r *http.Request){
	username,_ := r.Cookie("username")
	io.WriteString(w,"hi! "+username.Value)

}
