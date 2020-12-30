package vip_gosdk

// base
const (
	Format  = "JSON"
	Service = "com.vip.wpc.ospservice.vop.WpcVopOspService"
	Domain  = "https://gw.vipapis.com"
	Version = "1.0.0"
)

// method
const (
	ApplyPaymentMethod     = "applyPayment"
	OrderCreateMethod      = "orderCreate"
	OrderCancelMethod      = "orderCancel"
	SelectAddressMethod    = "selectAddress"
	GetOrderInfoListMethod = "getOrderInfoList"
)

// Default pagination parameters
const (
	DefaultPage     = 1
	DefaultPageSize = 20
)
