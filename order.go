package vip_gosdk

import (
	"encoding/json"
	"net/http"
)

type orderCreateRequest struct{}

func (o orderCreateRequest) ApiMethod() string {
	return OrderCreateMethod
}

func (o orderCreateRequest) RequestMethod() string {
	return http.MethodPost
}

type OrderCreateResponse struct {
	OrderCreateResp
}

func (resp *OrderCreateResponse) String() string {
	buf, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// 创建订单
// http://vop.vip.com/home;jsessionid=74C224D59EAC53838899EF34F2ED9BA7#/api/method/detail/com.vip.wpc.ospservice.vop.WpcVopOspService-1.0.0/orderCreate
func (client *Client) OrderCreate(req *OrderCreateReq) (resp *OrderCreateResponse, err error) {
	if req == nil {
		err = ErrInvalidOrderCreate
		return
	}
	resp = &OrderCreateResponse{}
	params := map[string]interface{}{
		"areaId":       req.AreaId,
		"vopChannelId": client.appKey,
		"userNumber":   client.userNumber,
		"sizeInfo":     req.SizeInfo,
		"address":      req.Address,
		"consignee":    req.Consignee,
		"mobile":       req.Mobile,
		"clientIp":     req.ClientIp,
		"traceId":      req.TraceId,
	}
	err = client.Do(orderCreateRequest{}, params, resp)
	return
}

//------------------------------------------------------------------------------

type orderCancelRequest struct{}

func (o orderCancelRequest) ApiMethod() string {
	return OrderCancelMethod
}

func (o orderCancelRequest) RequestMethod() string {
	return http.MethodPost
}

type OrderCancelResponse struct {
	OrderCancelResp
}

func (resp *OrderCancelResponse) String() string {
	buf, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// 取消订单
// http://vop.vip.com/home;jsessionid=74C224D59EAC53838899EF34F2ED9BA7#/api/method/detail/com.vip.wpc.ospservice.vop.WpcVopOspService-1.0.0/orderCancel
func (client *Client) OrderCancel(orderSn string) (resp *OrderCancelResponse, err error) {
	if orderSn == "" {
		err = ErrOrderSnRequired
		return
	}
	resp = &OrderCancelResponse{}
	params := map[string]interface{}{
		"vopChannelId": client.appKey,
		"userNumber":   client.userNumber,
		"orderSn":      orderSn,
	}
	err = client.Do(orderCancelRequest{}, params, resp)
	return
}

//------------------------------------------------------------------------------

type orderInfoListRequest struct{}

func (o orderInfoListRequest) ApiMethod() string {
	return GetOrderInfoListMethod
}

func (o orderInfoListRequest) RequestMethod() string {
	return http.MethodPost
}

type OrderInfoListResponse struct {
	OrderInfoListResp
}

func (resp *OrderInfoListResponse) String() string {
	buf, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// 订单状态信息
// http://vop.vip.com/home;jsessionid=74C224D59EAC53838899EF34F2ED9BA7#/api/method/detail/com.vip.wpc.ospservice.vop.WpcVopOspService-1.0.0/getOrderInfoList
func (client *Client) OrderInfoList(orderSnList string) (resp *OrderInfoListResponse, err error) {
	if orderSnList == "" {
		err = ErrOrderSnRequired
		return
	}
	resp = &OrderInfoListResponse{}
	params := map[string]interface{}{
		"vopChannelId": client.appKey,
		"userNumber":   client.userNumber,
		"page":         DefaultPage,
		"pageSize":     DefaultPageSize,
		"orderSnList":  orderSnList,
	}
	err = client.Do(orderInfoListRequest{}, params, resp)
	return
}
