package wechat

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"wechatpay/config"
	"wechatpay/wechat/query"
	"wechatpay/wechat/refund"
	"wechatpay/wechat/unified"
)

type WxPay interface {
	GetUnified() unified.WxUnified
	GetQuery() query.WxQuery
	GetRefund() refund.WxRefund
}

type WxPayContext struct {
	Cfg         *config.WxConfig
	Certificate *x509.Certificate
	PrivateKey  *rsa.PrivateKey
}

// 微信初始化
func NewWxPay(cfg *config.WxConfig) WxPay {
	privateKey, err := utils.LoadPrivateKeyWithPath(cfg.KeyPemPath)
	if err != nil {
		panic(fmt.Errorf("load private err:%s", err.Error()))
	}
	certificate, err := utils.LoadCertificateWithPath(cfg.CertPemPath)
	if err != nil {
		panic(fmt.Errorf("load certificate err:%s", err.Error()))
	}
	return &WxPayContext{Cfg: cfg, Certificate: certificate, PrivateKey: privateKey}
}

func (wx WxPayContext) GetUnified() unified.WxUnified {
	return unified.NewWxUnified(wx.Cfg, wx.Certificate, wx.PrivateKey)
}

func (wx WxPayContext) GetQuery() query.WxQuery {
	return query.NewWxQuery(wx.Cfg, wx.Certificate, wx.PrivateKey)
}

func (wx WxPayContext) GetRefund() refund.WxRefund {
	return refund.NewWxRefund(wx.Cfg, wx.Certificate, wx.PrivateKey)
}
