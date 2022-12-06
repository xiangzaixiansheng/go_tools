package jsonutil

import (
	"fmt"
	"testing"
)

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestFileCache(t *testing.T) {
	_people := People{Name: "xiangzai", Age: 18}
	fmt.Println("StructToJson", string(StructToJson(_people)))
	fmt.Println("ByteToHex", ByteToHex(StructToJson(_people)))
}
