package main

import (
	"net/http"
	"strconv"
	"github.com/csvikram/hello-world-golang/service"
)

func main() {

	http.HandleFunc("/isprime", isPrime)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		println("terminating")
		println(err.Error())
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("hello world"))
}

func isPrime(w http.ResponseWriter, r *http.Request)  {

	values := r.URL.Query()
	x :=values.Get("number")
	number,err := strconv.Atoi(x)
	if err == nil {
		w.Write([]byte(strconv.FormatBool(service.IsPrime(number))))
		return
	}
	w.Write([]byte("Invalid Value"))
}
