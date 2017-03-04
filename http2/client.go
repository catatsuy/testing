package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

func main() {
	c := &http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	for i := 0; i < 100; i++ {
		req, err := http.NewRequest("GET", "https://test.isucon.net:3000", nil)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		go func() {
			res, err := c.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}

			defer res.Body.Close()
			fmt.Println(res.Proto)
		}()
	}

	time.Sleep(100 * time.Second)
}
