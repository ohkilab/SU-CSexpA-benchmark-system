package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("request comed")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hi:)"))
	})
	lsnr, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	defer lsnr.Close()

	fmt.Printf("Listen on http://localhost:%d\n", lsnr.Addr().(*net.TCPAddr).Port)
	if err := http.Serve(lsnr, http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}
}
