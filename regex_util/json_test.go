package regex_util

import (
	"fmt"
	"testing"
)

func TestJsonPool(t *testing.T) {
	inputString := `以下数据：{"title": "标题一", "text": "内容一", "tag": "tag1"}{"title": "标题二", "text": "内容二", "tag": "tag二"}`

	jsonData, err := ExtractDataFromString(inputString)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(jsonData)

}
