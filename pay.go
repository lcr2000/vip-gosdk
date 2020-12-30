package vip_gosdk

import (
	"encoding/json"
	"net/http"
)

type applyPaymentRequest struct{}

func (o applyPaymentRequest) ApiMethod() string {
	return ApplyPaymentMethod
}

func (o applyPaymentRequest) RequestMethod() string {
	return http.MethodPost
}

type ApplyPaymentResponse struct {
	ApplyPaymentResp
}

func (resp *ApplyPaymentResponse) String() string {
	buf, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// 申请代扣
// http://vop.vip.com/home;jsessionid=74C224D59EAC53838899EF34F2ED9BA7#/api/method/detail/com.vip.wpc.ospservice.vop.WpcVopOspService-1.0.0/applyPayment
func (client *Client) ApplyPayment(orderSn, clientIp string) (resp *ApplyPaymentResponse, err error) {
	if orderSn == "" || clientIp == "" {
		err = ErrOrderSnOrIpEmpty
		return
	}
	resp = &ApplyPaymentResponse{}
	params := map[string]interface{}{
		"vopChannelId": client.appKey,
		"userNumber":   client.userNumber,
		"orderSn":      orderSn,
		"clientIp":     clientIp,
	}
	err = client.Do(applyPaymentRequest{}, params, resp)
	return
}
