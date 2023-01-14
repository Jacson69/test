package wx_util

import (
	"fmt"
	"testing"
)

func TestGetWxSign(t *testing.T) {
	url := "http://baidu.com"
	sign, err := GetWxSign(url)
	fmt.Println(sign, err)
}
