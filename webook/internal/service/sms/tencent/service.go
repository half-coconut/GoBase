package tencent

import (
	"GoBase/webook/pkg/ratelimit"
	"context"
	"fmt"
	"github.com/ecodeclub/ekit"
	"github.com/ecodeclub/ekit/slice"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type Service struct {
	// 因为腾讯云 api 要求这样传
	appId    *string
	signName *string
	client   *sms.Client
	limiter  ratelimit.Limiter
}

func NewService(client *sms.Client, appId string, signName string, limiter ratelimit.Limiter) *Service {
	return &Service{
		client:   client,
		appId:    ekit.ToPtr[string](appId),
		signName: ekit.ToPtr[string](signName),
		limiter:  limiter,
	}
}

func (s *Service) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	// 内部调用可以不校验，外部 微服务可以校验一下，参数是否为空
	req := sms.NewSendSmsRequest()
	req.SmsSdkAppId = s.appId
	req.SignName = s.signName
	req.TemplateId = ekit.ToPtr[string](tpl)
	req.PhoneNumberSet = s.toStringPtrSlice(numbers)
	req.TemplateParamSet = s.toStringPtrSlice(args)
	resp, err := s.client.SendSms(req)
	if err != nil {
		return err
	}
	for _, status := range resp.Response.SendStatusSet {
		if status.Code == nil || *(status.Code) != "Ok" {
			return fmt.Errorf("发送短信失败 %s,%s", *status.Code, *status.Message)
		}
	}
	return nil
}

func (s *Service) toStringPtrSlice(src []string) []*string {
	return slice.Map[string, *string](src, func(idx int, src string) *string {
		return &src
	})
}
