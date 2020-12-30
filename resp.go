package vip_gosdk

type BaseResp struct {
	ReturnCode    string `json:"returnCode"`
	ReturnMessage string `json:"returnMessage"`
}

// 查询地址库
type SelectAddressResp struct {
	BaseResp
	Result struct {
		ChildList []struct {
			AreaCode string `json:"areaCode"`
			AreaName string `json:"areaName"`
		} `json:"childList"`
		Info struct {
			AreaCode string `json:"areaCode"`
			AreaName string `json:"areaName"`
		} `json:"info"`
	} `json:"result"`
}

// 创建订单
type OrderCreateResp struct {
	BaseResp
	Result struct {
		Orders []struct {
			OrderSn string `json:"orderSn"`
		} `json:"orders"`
	} `json:"result"`
}

// 取消订单
type OrderCancelResp struct {
	BaseResp
	Result struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	} `json:"result"`
}

// 订单状态信息
type OrderInfoListResp struct {
	BaseResp
	Result []struct {
		ChildOrderSnList []struct {
			RealPayTotal string `json:"RealPayTotal"`
			ShippingFee  string `json:"ShippingFee"`
			Goods        []struct {
				GoodFullID string `json:"goodFullId"`
				Price      string `json:"price"`
				SizeID     string `json:"sizeId"`
				SizeNum    int64  `json:"sizeNum"`
			} `json:"goods"`
			OrderSn       string `json:"orderSn"`
			RefundStatus  int64  `json:"refundStatus"`
			StatusCode    int64  `json:"statusCode"`
			StatusName    string `json:"statusName"`
			TransportNo   string `json:"transportNo"`
			TransportName string `json:"transportName"`
		} `json:"childOrderSnList"`
		ParentOrderSn string `json:"parentOrderSn"`
	} `json:"result"`
}

// 申请代扣
type ApplyPaymentResp struct {
	BaseResp
	Result struct {
		ApplySuccess bool   `json:"applySuccess"`
		FailReason   string `json:"failReason"`
	} `json:"result"`
}
