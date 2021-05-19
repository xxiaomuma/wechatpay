# wechatpay
微信支付sdk for go 当前仅支持V3直连商户, 持续更新 </br>
 - 如果在集成过程中遇到问题, 请联系: 2521901773@qq.com
 安装
    go get -u github.com/xxiaomuma/wechatpay
微信支付
 - 基础配置
   ```
        pay := NewPay(&config.Config{
             Wx: &config.WxConfig{
                   AppId:               "xxx", // 微信APPID
                   MchId:               "xxx", // 商户号
                   ApiKey:              "xxx", // v2api key(暂时没有用到)
                   V3ApiKey:            "xxx", // v3api key
                   CertificateSerialNo: "xxx", // 证书序列号
                   CertPath:            "xxx/apiclient_cert.p12", // 证书地址
                   KeyPemPath:          "xxx/apiclient_key.pem",  
                   CertPemPath:         "xxx/apiclient_cert.pem",
             },
        })
   ```
   - pay使用全局变量
 - 统一下单
   ```
         param := unified.WxUnifiedParam {
		             OutTradeNo:  "xxxx",
		             NotifyUrl:   "xxx",
		             Description: "测试",
	         }
	         param.SetPayerOpenId("xxx")
	         param.SetCNYPayAmount(1)
	         result, err := pay.GetWxPay().GetUnified().JsApiPayV3(param)
   ```
   - 这里使用的是jsapi统一下单，GetUnified下提供APP, JSAPI, H5, NATIVE
    ```
        param := unified.WxUnifiedParam{
		            OutTradeNo:  "xxx",
		            NotifyUrl:   "xxx",
		            Description: "测试",
	       }
	       param.SetCNYPayAmount(1)
	       result, err := pay.GetWxPay().GetUnified().NativePayV3(param)
    ```
- 查询 </br>
    ```
  GetWxPay().GetQuery()
  ```
- 退款
    ```
    GetWxPay.GetRefund()
     ```