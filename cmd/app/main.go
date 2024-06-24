package main

import (
	"log"
)

func main() {
	err := loop()
	if err != nil {
		log.Fatal(err)
	}
}

func loop() error {
	return nil
}
