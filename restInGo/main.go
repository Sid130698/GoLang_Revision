package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloUser)
	http.ListenAndServe(":8080", nil)
}
func helloUser(writer http.ResponseWriter, request *http.Request){
	var  greeting string = "Hello, world!";
	fmt.Fprintln(writer, greeting)
}
