package unified

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"wechatpay/wechat/constant"
)

type wxH5V3Response struct {
	H5Url string `json:"h5_url"`
}

func (ctx WxUnifiedContext) H5PayV3(param WxUnifiedParam) (string, error) {
	requestAddress := fmt.Sprintf(WxUnifiedUrl, constant.WX_H5.String())
	client, err := ctx.Unified(&param, constant.WX_H5)
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
	result := wxH5V3Response{}
	_ = json.Unmarshal(body, &result)
	return result.H5Url, nil
}
