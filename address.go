package vip_gosdk

import (
	"encoding/json"
	"net/http"
)

type SelectAddressRequest struct{}

func (a SelectAddressRequest) ApiMethod() string {
	return SelectAddressMethod
}

func (a SelectAddressRequest) RequestMethod() string {
	return http.MethodPost
}

type SelectAddressResponse struct {
	SelectAddressResp
}

func (resp *SelectAddressResponse) String() string {
	buf, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// 查询地址库
// http://vop.vip.com/home;jsessionid=74C224D59EAC53838899EF34F2ED9BA7#/api/method/detail/com.vip.wpc.ospservice.vop.WpcVopOspService-1.0.0/selectAddress
func (client *Client) SelectAddress(areaCode string) (resp *SelectAddressResponse, err error) {
	resp = &SelectAddressResponse{}
	params := map[string]interface{}{
		"vopChannelId": client.appKey,
		"userNumber":   client.userNumber,
		"areaCode":     areaCode,
	}
	err = client.Do(SelectAddressRequest{}, params, resp)
	return
}