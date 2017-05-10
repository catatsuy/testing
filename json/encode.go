package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type codeResponse struct {
	Tree     *codeNode `json:"tree"`
	Username string    `json:"username"`
}

type codeNode struct {
	Name     string      `json:"name"`
	Kids     []*codeNode `json:"kids"`
	CLWeight float64     `json:"cl_weight"`
	Touches  int         `json:"touches"`
	MinT     int64       `json:"min_t"`
	MaxT     int64       `json:"max_t"`
	MeanT    int64       `json:"mean_t"`
}

func main() {
	f, err := os.Open("testdata/code.json.gz")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(gz)
	if err != nil {
		panic(err)
	}

	codeJSON := data
	var codeStruct codeResponse

	err = json.Unmarshal(codeJSON, &codeStruct)
	if err != nil {
		panic("unmarshal code.json: " + err.Error())
	}

	start := time.Now()
	for i := 0; i < 1000; i++ {
		data, err = json.Marshal(&codeStruct)
		if err != nil {
			panic("marshal code.json: " + err.Error())
		}
	}
	nanosec := time.Since(start).Nanoseconds()
	fmt.Printf("%fs\n", float64(nanosec)/1000000000)
}
