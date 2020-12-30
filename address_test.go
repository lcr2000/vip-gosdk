package vip_gosdk

import (
	"testing"
)

func TestClient_SelectAddress(t *testing.T) {
	type fields struct {
		appKey     string
		secret     string
		userNumber string
	}
	type args struct {
		areaCode string
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
			args:    args{},
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
			gotResp, err := client.SelectAddress(tt.args.areaCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(gotResp)
		})
	}
}
