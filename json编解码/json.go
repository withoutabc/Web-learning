package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Student 结构体小写不能编码
type Student struct {
	Id     int
	Name   string
	Age    int
	Gender string
}

// 结构体 ===> json 编码
// json ===> 结构体 解码
func main() {
	lily := Student{
		Id:     1,
		Name:   "Lily",
		Age:    20,
		Gender: "女士",
	}
	//编码
	//func Marshal(v any) ([]byte, error)
	encodeInfo, err := json.Marshal(&lily)
	if err != nil {
		log.Fatalf("json.Marshal err:%v", err)
		return
	}
	fmt.Println("encodeInfo:", string(encodeInfo))
	//解码
	//func Unmarshal(data []byte, v any) error
	var lily2 Student
	err = json.Unmarshal([]byte(encodeInfo), &lily2)
	if err != nil {
		log.Fatalf("json.Unmarshal err:%v", err)
		return
	}
	fmt.Println(lily2.Gender)
}
