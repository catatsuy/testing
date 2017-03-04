package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

func main() {
	certFile, _ := filepath.Abs("ssl/oreore.crt")
	keyFile, _ := filepath.Abs("ssl/oreore.key")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	})

	err := http.ListenAndServeTLS(":3000", certFile, keyFile, nil)

	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
}
