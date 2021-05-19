package refund

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
	// 微信申请退款 V3
	// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_5_9.shtml
	WxRefundUrl = "https://api.mch.weixin.qq.com/v3/refund/domestic/refunds"
)

type WxRefund interface {
	Refund(param WxRefundParam) (*WxRefundResponse, error)
}

type WxRefundContent struct {
	cfg         *config.WxConfig
	certificate *x509.Certificate
	privateKey  *rsa.PrivateKey
}

func NewWxRefund(cfg *config.WxConfig, certificate *x509.Certificate, privateKey *rsa.PrivateKey) WxRefund {
	return &WxRefundContent{cfg: cfg, certificate: certificate, privateKey: privateKey}
}

type WxRefundParam struct {
	TransactionId string           `json:"transaction_id,omitempty"`
	OutTradeNo    string           `json:"out_trade_no,omitempty"`
	OutRefundNo   string           `json:"out_refund_no"`
	Reason        string           `json:"reason,omitempty"`
	NotifyUrl     string           `json:"notify_url,omitempty"`
	FundsAccount  string           `json:"funds_account,omitempty"`
	Amount        *wxAmount        `json:"amount"`
	GoodsDetail   []*wxGoodsDetail `json:"goods_detail,omitempty"`
}

type wxAmount struct {
	Refund   int    `json:"refund"`
	Total    int    `json:"total"`
	Currency string `json:"currency"`
}

type wxGoodsDetail struct {
	MerchantGoodsId  string `json:"merchant_goods_id"`
	WechatpayGoodsId string `json:"wechatpay_goods_id"`
	GoodsName        string `json:"goods_name"`
	UnitPrice        int    `json:"unit_price"`
	RefundAmount     int    `json:"refund_amount"`
	RefundQuantity   int    `json:"refund_quantity"`
}

type WxRefundResponse struct {
	RefundId            string `json:"refund_id"`
	OutRefundNo         string `json:"out_refund_no"`
	TransactionId       string `json:"transaction_id"`
	OutTradeNo          string `json:"out_trade_no"`
	Channel             string `json:"channel"`
	UserReceivedAccount string `json:"user_received_account"`
	SuccessTime         string `json:"success_time"`
	CreateTime          string `json:"create_time"`
	Status              string `json:"status"`
	FundsAccount        string `json:"funds_account"`
}

func (param *WxRefundParam) SetRefundAmount(refundAmount float64, totalAmount float64) {
	refundInt := int(refundAmount * 100)
	totalInt := int(totalAmount * 100)
	param.Amount = &wxAmount{Refund: refundInt, Total: totalInt, Currency: "CNY"}
}

func (ctx WxRefundContent) Refund(param WxRefundParam) (*WxRefundResponse, error) {
	requestAddress := WxRefundUrl
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
	response, err := client.Post(context.Background(), requestAddress, param)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	result := WxRefundResponse{}
	_ = json.Unmarshal(body, &result)
	return &result, nil
}
