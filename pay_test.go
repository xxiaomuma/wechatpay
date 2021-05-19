package pay

import (
	"fmt"
	"testing"
	"wechatpay/config"
	"wechatpay/wechat/refund"
	"wechatpay/wechat/unified"
)

func TestPay_JsApi(t *testing.T) {
	pay := NewPay(&config.Config{
		Wx: &config.WxConfig{
			AppId:               "xxx",
			MchId:               "xxx",
			ApiKey:              "xxx",
			V3ApiKey:            "xxx",
			CertificateSerialNo: "xxx",
			CertPath:            "xxx/apiclient_cert.p12",
			KeyPemPath:          "xxx/apiclient_key.pem",
			CertPemPath:         "xxx/apiclient_cert.pem",
		},
	})
	param := unified.WxUnifiedParam{
		OutTradeNo:  "xxxx",
		NotifyUrl:   "xxx",
		Description: "测试",
	}
	param.SetPayerOpenId("xxx")
	param.SetCNYPayAmount(1)
	result, err := pay.GetWxPay().GetUnified().JsApiPayV3(param)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(result)
}

func TestPay_Native(t *testing.T) {
	pay := NewPay(&config.Config{
		Wx: &config.WxConfig{
			AppId:               "xxx",
			MchId:               "xxx",
			ApiKey:              "xxx",
			V3ApiKey:            "xxx",
			CertificateSerialNo: "xxx",
			CertPath:            "xxx/apiclient_cert.p12",
			KeyPemPath:          "xxx/apiclient_key.pem",
			CertPemPath:         "xxx/apiclient_cert.pem",
		},
	})
	param := unified.WxUnifiedParam{
		OutTradeNo:  "xxx",
		NotifyUrl:   "xxx",
		Description: "测试",
	}
	param.SetCNYPayAmount(1)
	result, err := pay.GetWxPay().GetUnified().NativePayV3(param)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(result)
}

func TestPay_Query(t *testing.T) {
	pay := NewPay(&config.Config{
		Wx: &config.WxConfig{
			AppId:               "xxx",
			MchId:               "xxx",
			ApiKey:              "xxx",
			V3ApiKey:            "xxx",
			CertificateSerialNo: "xxx",
			CertPath:            "xxx/apiclient_cert.p12",
			KeyPemPath:          "xxx/apiclient_key.pem",
			CertPemPath:         "xxx/apiclient_cert.pem",
		},
	})
	result, err := pay.GetWxPay().GetQuery().QueryByTradeNo("xxx")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(result)
}

func TestNewPay_Refund(t *testing.T) {
	pay := NewPay(&config.Config{
		Wx: &config.WxConfig{
			AppId:               "xxx",
			MchId:               "xxx",
			ApiKey:              "xxx",
			V3ApiKey:            "xxx",
			CertificateSerialNo: "xxx",
			CertPath:            "xxx/apiclient_cert.p12",
			KeyPemPath:          "xxx/apiclient_key.pem",
			CertPemPath:         "xxx/apiclient_cert.pem",
		},
	})
	param := refund.WxRefundParam{
		OutTradeNo:  "xxx",
		OutRefundNo: "xxx",
	}
	param.SetRefundAmount(2, 1)
	result, err := pay.GetWxPay().GetRefund().Refund(param)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(result)
}
