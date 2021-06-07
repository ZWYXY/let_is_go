package main

import (
	"encoding/json"
	"fmt"
)

/**
演示结构体标签在JSON中的应用
*/

type Movie struct {
	Name  string   `json:"movieName"`
	Price int      `json:"price"`
	Actor []string `json:"actorList"`
}

func main() {

	oneMovie := Movie{"千与千寻", 1000, []string{"新海诚", "宫崎骏"}}

	marshalJSON, err := json.Marshal(oneMovie)
	if err != nil {
		fmt.Println("JSON格式化失败")
		return
	}
	fmt.Printf("%s\n", marshalJSON)

	// 反序列化
	myMovie := Movie{}
	err = json.Unmarshal(marshalJSON, &myMovie)
	if err != nil {
		fmt.Println("反序列化失败")
	}
	fmt.Println(myMovie)
}
