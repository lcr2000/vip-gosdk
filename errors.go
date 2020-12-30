package vip_gosdk

import "errors"

var (
	// ErrInvalidOrderCreate invalid order create params
	ErrInvalidOrderCreate = errors.New("invalid order create params")
	// ErrOrderSnRequired order number required
	ErrOrderSnRequired = errors.New("order number required")
	// ErrOrderSnAndIpRequired order number or ip is empty
	ErrOrderSnOrIpEmpty = errors.New("order number or ip is empty")
)
