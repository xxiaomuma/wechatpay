package pay

import (
	"wechatpay/config"
	"wechatpay/wechat"
)

type Pay struct {
	cfg   *config.Config
	wxPay wechat.WxPay
}

//初始化支付
func NewPay(cfg *config.Config) *Pay {
	pay := Pay{cfg: cfg}
	if cfg.Wx != nil {
		pay.wxPay = wechat.NewWxPay(cfg.Wx)
	}
	return &pay
}

func (pay *Pay) GetWxPay() wechat.WxPay {
	return pay.wxPay
}
