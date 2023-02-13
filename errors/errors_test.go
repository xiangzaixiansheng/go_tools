package openwechat

import (
	"errors"
	"fmt"
	"testing"
)

func TestIsNetworkError(t *testing.T) {
	var err = errors.New("test error")
	//err.Error() 打印error的字符串
	err = fmt.Errorf("%w: %s", NetworkErr, err.Error())
	if !IsNetworkError(err) {
		t.Error("err is not network error")
	}

	err = errors.New("这是一个error")
	fmt.Printf("type:%T val:%v\n", err, err) // 输出结果：type:*errors.errorString val:这是一个error

	err2 := fmt.Errorf("这是第%d个error", 2)
	fmt.Printf("err2Type:%T err2Val:%v\n", err2, err2) // 输出结果：err2Type:*errors.errorString err2Val:这是第2个error

}
