package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	// key 变小写
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
	Age      int
	Birthday string
	Sex      string
	Email    string
	Phone    string
}

func testMap() {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["userName"] = "user2"
	m["age"] = 99
	m["sex"] = "男"

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("json.Marshal failed,err", err)
		return
	}

	fmt.Printf("%s\n", string(data))

}

func main() {

	user1 := &User{
		UserName: "user1",
		NickName: "上课看似",
		Age:      100,
		Birthday: "2019/04/07",
		Sex:      "男",
		Email:    "test@qq.com",
		Phone:    "112",
	}

	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Printf("json.Marshal failed,err", err)
		return
	}

	fmt.Printf("%s\n", string(data))

	testMap()

}
