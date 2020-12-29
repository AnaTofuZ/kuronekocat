package main

import (
	"log"
	"os"

	"github.com/anatofuz/kuronekocat"
)

func main() {
	os.Exit(run())
}

func run() int {
	err := kuronekocat.NewKuronekoCmd().Execute()
	if err != nil {
		log.Println(err)
		return 1
	}
	return 0
}
