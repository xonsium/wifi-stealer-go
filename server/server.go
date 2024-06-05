package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var port string = ":6969"

func index(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	r := make(map[string]interface{})
	err := decoder.Decode(&r)
	if err != nil {
		panic(err)
	}
	jsonString, err := json.Marshal(r)
	if err != nil {
		log.Fatalf("encoding into json failed %s\n", err)
	}
	fmt.Println(string(jsonString))
}

func main() {
	http.HandleFunc("/", index)
	log.Println("started server at PORT " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
