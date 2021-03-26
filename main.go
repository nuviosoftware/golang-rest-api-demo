package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		d, _ := ioutil.ReadAll(req.Body)
		log.Printf("Body is %s", d)
	})
	http.ListenAndServe(":9090", nil)
}
