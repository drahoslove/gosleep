package main

// note build as:
//   go build -ldflags "-H windowsgui"

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/drahoslove/gosleep/sleep"
)

// server settings
const (
	IP   = "0.0.0.0"
	PORT = 6969
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		var indexTmpl = template.Must(template.ParseFiles("index.html"))
		indexTmpl.Execute(res, nil)
	})
	http.HandleFunc("/lock", func(res http.ResponseWriter, req *http.Request) {
		go sleep.LockScreen()
	})
	http.HandleFunc("/sleep", func(res http.ResponseWriter, req *http.Request) {
		go sleep.Sleep()
	})
	http.HandleFunc("/hibernate", func(res http.ResponseWriter, req *http.Request) {
		go sleep.Hibernate()
	})

	fmt.Printf("gosleep server starting: %v:%v\n", IP, PORT)
	http.ListenAndServe(fmt.Sprintf("%v:%v", IP, PORT), nil)
}
