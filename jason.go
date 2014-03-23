package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson" // for json get
)

type MyData struct {
	Name  string  `json:"item"`
	Other float32 `json:"amount"`
}

func main() {
	var detail MyData
	//啊名为   v发发
	detail.Name = "1"
	detail.Other = 2

	body, err := json.Marshal(detail)
	if err != nil {
		panic(err.Error())
		///
	}

	js, err := simplejson.NewJson(body)
	if err != nil {
		panic(err.Error())
	}
	////////ccccc

	fmt.Println(js)

}
