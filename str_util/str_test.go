package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFileCache(t *testing.T) {
	// 判断类型
	fmt.Println(reflect.TypeOf(ToString(map[string]string{"xiangzai": "test"})))
	fmt.Println(ToString(map[string]string{"xiangzai": "test"}))
}
