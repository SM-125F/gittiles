package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var ip = flag.String("ip", "localhost", "ip to run on")
	var port = flag.String("port", ":8001", "port to run on")
	flag.Parse()

	address := *ip + *port

	http.Handle("/gitvis/", http.StripPrefix("/gitvis/",
		http.FileServer(http.Dir("static"))))

	fmt.Println("Webserver started at", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Panic("Could not start webapp! ", err.Error())
	}
}
