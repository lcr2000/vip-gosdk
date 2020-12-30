package vip_gosdk

import (
	"testing"
)

func TestClient_OrderCreate(t *testing.T) {
	type fields struct {
		appKey     string
		secret     string
		userNumber string
	}
	type args struct {
		req *OrderCreateReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			fields: fields{
				appKey:     appKey,
				secret:     secret,
				userNumber: userNumber,
			},
			args: args{req: &OrderCreateReq{
				AreaId:    "XXX",
				SizeInfo:  "XXX",
				Address:   "XXX",
				Consignee: "XXX",
				Mobile:    "XXX",
				ClientIp:  "XXX",
				TraceId:   "XXX",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				appKey:     tt.fields.appKey,
				secret:     tt.fields.secret,
				userNumber: tt.fields.userNumber,
			}
			gotResp, err := client.OrderCreate(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(gotResp)
		})
	}
}

func TestClient_OrderCancel(t *testing.T) {
	type fields struct {
		appKey     string
		secret     string
		userNumber string
	}
	type args struct {
		orderSn string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			fields: fields{
				appKey:     appKey,
				secret:     secret,
				userNumber: userNumber,
			},
			args:    args{orderSn: "XXX"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				appKey:     tt.fields.appKey,
				secret:     tt.fields.secret,
				userNumber: tt.fields.userNumber,
			}
			gotResp, err := client.OrderCancel(tt.args.orderSn)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderCancel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(gotResp)
		})
	}
}

func TestClient_OrderInfoList(t *testing.T) {
	type fields struct {
		appKey     string
		secret     string
		userNumber string
	}
	type args struct {
		orderSnList string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			fields: fields{
				appKey:     appKey,
				secret:     secret,
				userNumber: userNumber,
			},
			args:    args{orderSnList: "XXX"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				appKey:     tt.fields.appKey,
				secret:     tt.fields.secret,
				userNumber: tt.fields.userNumber,
			}
			gotResp, err := client.OrderInfoList(tt.args.orderSnList)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderInfoList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(gotResp)
		})
	}
}
