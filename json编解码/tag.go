package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Teacher struct {
	Name    string `json:"-"`                 //这个编码不参与
	Subject string `json:"subject_name"`      //字段编码成后面的部分
	Age     int    `json:"age,string"`        //字段名为age,string类型
	Address string `json:"address,omitempty"` //如果这个字段为空，则不参与编码
	gender  string
}

func main() {
	t1 := Teacher{
		Name:    "Duke",
		Subject: "Golang",
		Age:     18,
		gender:  "Man",
	}
	fmt.Println(t1)
	encodeInfo, err := json.Marshal(&t1)
	if err != nil {
		log.Fatalf("json.marshal err:%v", err)
		return
	}
	fmt.Println(string(encodeInfo))
}
