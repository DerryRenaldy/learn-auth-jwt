package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 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
type Person struct {
	First string
}

func main() {
	p1 := Person{
		First: "Derry",
	}

	p2 := Person{
		First: "Budi",
	}

	xp := []Person{p1, p2}

	bs, err := json.Marshal(xp)
	if err != nil {
		log.Panicf("got panic when marshaling: %v \n", err)
		return
	}

	fmt.Printf("JSON string: %v \n", string(bs))

	var xp2 []Person

	err = json.Unmarshal(bs, &xp2)
	if err != nil {
		log.Panicf("got panic when unmarshaling: %v \n", err)
		return
	}

	fmt.Printf("From JSON into golang struct: %v \n", xp2)

}
