package unified

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"time"
	"wechatpay/config"
	"wechatpay/wechat/constant"
)

const (
	// V3支付
	// https://pay.weixin.qq.com/wiki/doc/apiv3/index.shtml
	WxUnifiedUrl = "https://api.mch.weixin.qq.com/v3/pay/transactions/%s"
)

type WxUnified interface {
	JsApiPayV3(param WxUnifiedParam) (map[string]string, error)
	NativePayV3(param WxUnifiedParam) (string, error)
	H5PayV3(param WxUnifiedParam)
}

type WxUnifiedContext struct {
	cfg         *config.WxConfig
	certificate *x509.Certificate
	privateKey  *rsa.PrivateKey
}

func NewWxUnified(cfg *config.WxConfig, certificate *x509.Certificate, privateKey *rsa.PrivateKey) WxUnified {
	return &WxUnifiedContext{cfg: cfg, certificate: certificate, privateKey: privateKey}
}

// 统一下单
type WxUnifiedParam struct {
	AppId       string     `json:"appid"`
	MchId       string     `json:"mchid"`
	Description string     `json:"description"`
	OutTradeNo  string     `json:"out_trade_no"`
	Attach      string     `json:"attach,omitempty"`
	NotifyUrl   string     `json:"notify_url"`
	GoodsTag    string     `json:"goods_tag,omitempty"`
	Amount      *amount    `json:"amount"`
	Payer       *payer     `json:"payer,omitempty"`
	SceneInfo   *sceneInfo `json:"scene_info,omitempty"`
}

type amount struct {
	Total    int    `json:"total"`
	Currency string `json:"currency,omitempty"`
}

type payer struct {
	Openid string `json:"openid"`
}

type sceneInfo struct {
	PayerClientIp string  `json:"payer_client_ip"`
	H5Info        *h5Info `json:"h5_info"`
}

type h5Info struct {
	Type string `json:"type"`
}

// 设置金额
func (param *WxUnifiedParam) SetCNYPayAmount(price float64) {
	priceInt := int(price * 100)
	param.Amount = &amount{Total: priceInt, Currency: "CNY"}
}

// 设置支付者
// 当支付JSAPI支付时必须设置openid
func (param *WxUnifiedParam) SetPayerOpenId(openId string) {
	param.Payer = &payer{Openid: openId}
}

// 支付场景描述
// 当支付H5支付时必须设置支付者ip和h5类型
func (param *WxUnifiedParam) SetSceneInfo(payerClientIp string, h5Type string) {
	param.SceneInfo = &sceneInfo{PayerClientIp: payerClientIp, H5Info: &h5Info{Type: h5Type}}
}

func (ctx WxUnifiedContext) Unified(param *WxUnifiedParam, unifiedType constant.WxUnifiedType) (*core.Client, error) {
	param.AppId = ctx.cfg.AppId
	param.MchId = ctx.cfg.MchId
	if len(param.NotifyUrl) == 0 {
		param.NotifyUrl = ctx.cfg.NotifyUrl
	}
	err := param.checkParam(unifiedType)
	if err != nil {
		return nil, err
	}
	opts := []option.ClientOption{
		option.WithMerchant(ctx.cfg.MchId, ctx.cfg.CertificateSerialNo, ctx.privateKey),
		option.WithWechatPay([]*x509.Certificate{ctx.certificate}),
		option.WithTimeout(5 * time.Second),
		option.WithoutValidator(),
	}
	client, err := core.NewClient(context.Background(), opts...)
	if err != nil {
		return nil, fmt.Errorf("new wechat pay client err:%s", err.Error())
	}
	return client, nil
}

func (param WxUnifiedParam) parse2RequestBody() (string, error) {
	jsonParam, err := json.Marshal(param)
	if err != nil {
		return "", err
	}
	return string(jsonParam), nil
}

func (param WxUnifiedParam) checkParam(unifiedType constant.WxUnifiedType) error {
	if len(param.AppId) == 0 {
		return errors.New("appId must not be empty")
	} else if len(param.MchId) == 0 {
		return errors.New("mchId must not be empty")
	} else if len(param.Description) == 0 {
		return errors.New("description must not be empty")
	} else if len(param.OutTradeNo) == 0 {
		return errors.New("outTradeNo must not be empty")
	} else if len(param.NotifyUrl) == 0 {
		return errors.New("notifyUrl must not be empty")
	} else if param.Amount == nil {
		return errors.New("amount must not be null")
	} else if param.Amount.Total == 0 {
		return errors.New("amount.total must not be null")
	}

	switch unifiedType {
	case constant.WX_JSAPI:
		if param.Payer == nil {
			return errors.New("payer must not be null")
		} else if len(param.Payer.Openid) == 0 {
			return errors.New("payer.openId must not be null")
		}
	case constant.WX_APP:
		break
	case constant.WX_H5:
		if param.SceneInfo == nil {
			return errors.New("scene_info must not be null")
		} else if len(param.SceneInfo.PayerClientIp) == 0 {
			return errors.New("payer_client_ip must not be null")
		} else if param.SceneInfo.H5Info == nil {
			return errors.New("h5_info must not be null")
		} else if len(param.SceneInfo.H5Info.Type) == 0 {
			return errors.New("type must not be null")
		}
	case constant.WX_NATIVE:
		break
	}
	return nil
}
