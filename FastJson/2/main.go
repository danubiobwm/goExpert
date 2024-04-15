package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var p fastjson.Parser

	jsonData := `{"user":{"name": "John", "age":30}}`

	value, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	// user := value.GetObject("user")
	// fmt.Printf("User name: %s\n", user.Get("name"))
	// fmt.Printf("User idade: %s\n", user.Get("age"))

	userJson := value.GetObject("user").String()
	var user User
	if err := json.Unmarshal([]byte(userJson), &user); err != nil {
		panic(err)
	}

	fmt.Println(user.Name, user.Age)

}
