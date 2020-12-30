package vip_gosdk

import (
	"testing"
)

func TestClient_ApplyPayment(t *testing.T) {
	type fields struct {
		appKey     string
		secret     string
		userNumber string
	}
	type args struct {
		orderSn  string
		clientIp string
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
			args:    args{orderSn: "XXX", clientIp: "XXX"},
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
			gotResp, err := client.ApplyPayment(tt.args.orderSn, tt.args.clientIp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyPayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(gotResp)
		})
	}
}
