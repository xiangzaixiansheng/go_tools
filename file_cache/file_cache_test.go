package cache

import (
	"fmt"
	"testing"
)

func TestFileCache(t *testing.T) {
	fc := NewFileCache("hello.txt")
	fc.SetItem("test1", "{123:456}")
	fc.SetItem("test2", "{123:456}")

	fmt.Println(fc.GetStrFromCache("test2"))

	//time.Sleep(5 * time.Second)
	//fc.Clear()

}
