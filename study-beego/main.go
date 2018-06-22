package main

import (
	"net/http"
	"io"
	"log"
)

func main() {
	http.HandleFunc("/",sayHello)
	http.HandleFunc("/bye",sayBye)
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func sayHello(w http.ResponseWriter,r *http.Request ){
	io.WriteString(w,"Hello world")
}

func sayBye(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,"Bye bye")
}



