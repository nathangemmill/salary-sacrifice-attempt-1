package main

import (
	"html/template"
	"net/http"
)

//structure which identifies the {{.}} value from affordable.html
type CanIAffordIt struct {
	CanIAffordIt string
}

//w is to write and r is to read the html file. * is a pointer
func salsac(w http.ResponseWriter, r *http.Request) {
	//t is template
	var t = template.Must(template.ParseFiles("salsac.html"))
	t.Execute(w, nil)
}

func affordable(w http.ResponseWriter, r *http.Request) {
	t := template.New("CanIAffordIt")
	t, _ = t.Parse("NA FAM. YOU BROKE.")
}

func main() {
	//nil defined the domain name that you want it to run on. nil is localhost.
	http.HandleFunc("/", salsac)
	http.HandleFunc("/affordable", affordable)
	http.ListenAndServe(":8888", nil)
}
