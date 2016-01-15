package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	listen = flag.String("listen", "", "Address to listen for broadcast events, e.g. 127.0.0.1:1234")
)

func main() {
	flag.Parse()

	if *listen == "" {
		log.Fatal("Must specify --listen")
	}

	fmt.Println("Hello World!")
}
