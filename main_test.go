package vip_gosdk

import (
	"flag"
	"fmt"
	"testing"
)

var (
	appKey     string
	secret     string
	userNumber string
)

// 鉴于安全原因, 不便暴露 appKey/secret 等信息
// 所以在测试命令后提供了如下参数.
func TestMain(m *testing.M) {
	flag.StringVar(&appKey, "appKey", "", "appKey")
	flag.StringVar(&secret, "secret", "", "secret")
	flag.StringVar(&userNumber, "userNumber", "", "userNumber")
	if !flag.Parsed() {
		flag.Parse()
	}
	fmt.Printf("-args -appId %s -secret %s -userNumber %s\n", appKey, secret, userNumber)
	m.Run()
}
