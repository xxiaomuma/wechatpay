package unified

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"wechatpay/wechat/constant"
)

type wxNativeResponse struct {
	CodeUrl string `json:"code_url"`
}

func (ctx WxUnifiedContext) NativePayV3(param WxUnifiedParam) (string, error) {
	requestAddress := fmt.Sprintf(WxUnifiedUrl, constant.WX_NATIVE.String())
	client, err := ctx.Unified(&param, constant.WX_NATIVE)
	if err != nil {
		return "", err
	}
	response, err := client.Post(context.Background(), requestAddress, param)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	result := wxNativeResponse{}
	_ = json.Unmarshal(body, &result)
	return result.CodeUrl, nil
}
