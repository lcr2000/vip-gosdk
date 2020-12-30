 # vip-sdk
 唯品会开放平台Golang版本SDK
 
 ## 单元测试
 鉴于安全原因, 不便暴露 appKey/secret 等信息.
 所以在测试命令后提供了如下参数.
 ```go
 go test -v -args -appKey xxx -secret xxx -userNumber xxx
 ```

## 快速开始
 ```go
 import (
 	"fmt"
 	"github.com/lcr2000/vip-gosdk"
 )
 
 var Client *vip_gosdk.Client
 
 func init()  {
 	Client = vip_gosdk.NewClient("appKey", "secret", "userNumber")
 }
 
 func GetAddress() {
 	res, err := Client.SelectAddress("")
 	if err != nil {
 		fmt.Println(err)
 	}
 	fmt.Println(res)
 }
 ```

