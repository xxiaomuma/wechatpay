package query

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"io/ioutil"
	"time"
	"wechatpay/config"
)

const (
	// 微信查询  V3
	// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_4_2.shtml
	WxQueryUrl = "https://api.mch.weixin.qq.com/v3/pay/transactions/out-trade-no/%s?mchid=%s"
)

type WxQuery interface {
	// 根据商户号查询
	QueryByTradeNo(tradeNo string) (*WxOrderResponse, error)
}

type WxQueryContext struct {
	cfg         *config.WxConfig
	certificate *x509.Certificate
	privateKey  *rsa.PrivateKey
}

func NewWxQuery(cfg *config.WxConfig, certificate *x509.Certificate, privateKey *rsa.PrivateKey) WxQuery {
	return &WxQueryContext{cfg: cfg, certificate: certificate, privateKey: privateKey}
}

type WxOrderResponse struct {
	AppId           string             `json:"appid"`
	MchId           string             `json:"mchid"`
	OutTradeNo      string             `json:"out_trade_no"`
	TransactionId   string             `json:"transaction_id"`
	TradeType       string             `json:"trade_type"`
	TradeState      string             `json:"trade_state"`
	TradeStateDesc  string             `json:"trade_state_desc"`
	BankType        string             `json:"bank_type"`
	Attach          string             `json:"attach"`
	SuccessTime     time.Time          `json:"success_time"`
	Payer           WxPayer            `json:"payer"`
	Amount          WxAmount           `json:"amount"`
	SceneInfo       *WxSceneInfo       `json:"scene_info"`
	PromotionDetail *WxPromotionDetail `json:"promotion_detail"`
}

type WxPayer struct {
	Openid string `json:"openid"`
}

type WxAmount struct {
	Total         int    `json:"total"`
	PayerTotal    int    `json:"payer_total"`
	Currency      string `json:"currency"`
	PayerCurrency string `json:"payer_currency"`
}

type WxSceneInfo struct {
	DeviceId string `json:"device_id"`
}

type WxPromotionDetail struct {
	CouponId            string         `json:"coupon_id"`
	Name                string         `json:"name"`
	Scope               string         `json:"scope"`
	Type                string         `json:"type"`
	Amount              int            `json:"amount"`
	StockId             string         `json:"stock_id"`
	WechatpayContribute int            `json:"wechatpay_contribute"`
	MerchantContribute  int            `json:"merchant_contribute"`
	OtherContribute     int            `json:"other_contribute"`
	Currency            string         `json:"currency"`
	GoodsDetail         *WxGoodsDetail `json:"goods_detail"`
}

type WxGoodsDetail struct {
	GoodsId        string `json:"goods_id"`
	Quantity       int    `json:"quantity"`
	UnitPrice      int    `json:"unit_price"`
	DiscountAmount int    `json:"discount_amount"`
	GoodsRemark    string `json:"goods_remark"`
}

func (ctx WxQueryContext) QueryByTradeNo(tradeNo string) (*WxOrderResponse, error) {
	requestAddress := fmt.Sprintf(WxQueryUrl, tradeNo, ctx.cfg.MchId)
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
	response, err := client.Get(context.Background(), requestAddress)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	result := WxOrderResponse{}
	_ = json.Unmarshal(body, &result)
	return &result, nil
}
