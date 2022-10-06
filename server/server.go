package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	First string
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := Person{
		First: "Derry",
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Printf("encoded bad data: %v \n", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	var p1 []Person
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Printf("encoded bad data: %v \n", err)
	}
	log.Println(p1)
}

func main() {
	// curl http://localhost:8080/encode
	http.HandleFunc("/encode", foo)
	// curl -XGET -H "Content-type: application/json" -d '[{"First": "Derry Renaldy"}, {"First": "Budi Sunarya"}]' 'localhost:8080/decode'
	http.HandleFunc("/decode", bar)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
