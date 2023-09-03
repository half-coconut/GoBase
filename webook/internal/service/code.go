package service

import (
	"GoBase/webook/internal/repository"
	"GoBase/webook/internal/service/sms"
	"context"
	"fmt"
	"math/rand"
)

/**
mockgen -source=webook/internal/service/code.go -package=svcmocks -destination=webook/internal/service/mocks/code.mock.go
*/

var ErrCodeSendTooMany = repository.ErrCodeSendTooMany

const codeTplId = "1877556"

type CodeService interface {
	Send(ctx context.Context, biz string, phone string) error
	Verify(ctx context.Context, biz string, phone string, inputCode string) (bool, error)
}

// codeService 短信验证码的实现
type codeService struct {
	repo   repository.CodeRepository
	smsSvc sms.Service
}

func NewCodeService(repo repository.CodeRepository, smsSvc sms.Service) CodeService {
	return &codeService{
		repo:   repo,
		smsSvc: smsSvc,
	}
}

// Send 生成一个随机验证码，并发送
func (c *codeService) Send(ctx context.Context, biz string, phone string) error {
	code := c.generate()
	err := c.repo.Store(ctx, biz, phone, code)
	if err != nil {
		return err
	}
	err = c.smsSvc.Send(ctx, codeTplId, []string{code}, phone)
	return err
}

// Verify 验证验证码
func (c *codeService) Verify(ctx context.Context,
	biz string,
	phone string,
	inputCode string) (bool, error) {
	ok, err := c.repo.Verify(ctx, biz, phone, inputCode)
	// 这里我们在 service 层面上对 Handler 屏蔽了最为特殊的错误
	if err == repository.ErrCodeVerifyTooManyTimes {
		// 在接入了告警之后，这边要告警
		// 因为这意味着有人在搞你
		return false, nil
	}
	return ok, err
}

func (c *codeService) generate() string {
	// 用随机数生成一个
	num := rand.Intn(999999)
	return fmt.Sprintf("%6d", num)
}
