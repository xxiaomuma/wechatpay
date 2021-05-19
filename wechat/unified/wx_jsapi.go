package unified

import (
	"context"
	"crypto"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"io/ioutil"
	"strconv"
	"time"
	"wechatpay/wechat/constant"
)

type wxJsApiV3Response struct {
	PrepayId string `json:"prepay_id"`
}

func (ctx *WxUnifiedContext) JsApiPayV3(param WxUnifiedParam) (map[string]string, error) {
	requestAddress := fmt.Sprintf(WxUnifiedUrl, constant.WX_JSAPI.String())
	client, err := ctx.Unified(&param, constant.WX_JSAPI)
	if err != nil {
		return nil, err
	}
	response, err := client.Post(context.Background(), requestAddress, param)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	result := wxJsApiV3Response{}
	_ = json.Unmarshal(body, &result)
	return ctx.jsApiStartPay(result.PrepayId)
}

func (ctx WxUnifiedContext) jsApiStartPay(prepayId string) (map[string]string, error) {
	var result = map[string]string{}
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce, err := generateNonceStr()
	if err != nil {
		return nil, err
	}
	wrap := "prepay_id=" + prepayId
	message := fmt.Sprintf("%s\n%s\n%s\n%s\n", ctx.cfg.AppId, timestamp, nonce, wrap)
	h := crypto.Hash.New(crypto.SHA256)
	h.Write([]byte(message))
	hashed := h.Sum(nil)
	signature, err := ctx.privateKey.Sign(rand.Reader, hashed, crypto.SHA256)
	if err != nil {
		return nil, err
	}
	result["appid"] = ctx.cfg.AppId
	result["timeStamp"] = timestamp
	result["nonceStr"] = nonce
	result["package"] = wrap
	result["signType"] = "RSA"
	result["paySign"] = base64.StdEncoding.EncodeToString(signature)
	return result, nil
}

func generateNonceStr() (string, error) {
	bytes := make([]byte, consts.NonceLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	symbolsByteLength := byte(len(consts.Symbols))
	for i, b := range bytes {
		bytes[i] = consts.Symbols[b%symbolsByteLength]
	}
	return string(bytes), nil
}
