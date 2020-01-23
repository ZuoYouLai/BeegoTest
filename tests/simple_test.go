package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Employmee struct {
	Name   string
	Age    int
	Number int
}

// 生成 json 对象
func Test_json(t *testing.T) {
	user := Employmee{Name: "samLai", Age: 28, Number: 1000}
	data, _ := json.Marshal(user)
	fmt.Println(string(data))
}

// 解析 json 字符串
func Test_parseJson(t *testing.T) {
	str := `{"Name":"samLai","Age":28,"Number":1000}`
	em := &Employmee{}
	err := json.Unmarshal([]byte(str), em)
	fmt.Println(err)
	fmt.Println(*em)
}
