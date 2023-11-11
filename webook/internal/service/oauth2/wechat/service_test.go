//go:build manual

package wechat

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	type args struct {
		appId     string
		appSecret string
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.appId, tt.args.appSecret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_AuthURL(t *testing.T) {
	type fields struct {
		appId     string
		appSecret string
		client    *http.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				appId:     tt.fields.appId,
				appSecret: tt.fields.appSecret,
				client:    tt.fields.client,
			}
			got, err := s.AuthURL(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AuthURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_manual_VerifyCode(t *testing.T) {
	os.Setenv("WECHAT_APP_ID", "wx1b4c7610fc671845")
	os.Setenv("WECHAT_APP_SECRET", "")
	//show("WECHAT_APP_ID")
	//show("WECHAT_APP_SECRET")

	appId, ok := os.LookupEnv("WECHAT_APP_ID")
	if !ok {
		panic("没有找到环境变量 WECHAT_APP_ID")
	}
	appKey, ok := os.LookupEnv("WECHAT_APP_SECRET")
	if !ok {
		panic("没有找到环境变量 WECHAT_APP_SECRET")
	}
	svc := NewService(appId, appKey)
	res, err := svc.VerifyCode(context.Background(), "", "")
	require.NoError(t, err)
	t.Log(res)
}
