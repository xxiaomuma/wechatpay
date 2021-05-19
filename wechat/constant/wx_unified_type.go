package constant

type WxUnifiedType int

const (
	WX_JSAPI  WxUnifiedType = 1
	WX_APP    WxUnifiedType = 2
	WX_H5     WxUnifiedType = 3
	WX_NATIVE WxUnifiedType = 4
)

func (wx WxUnifiedType) String() string {
	switch wx {
	case WX_JSAPI:
		return "jsapi"
	case WX_APP:
		return "app"
	case WX_H5:
		return "h5"
	case WX_NATIVE:
		return "native"
	}
	return "N/A"
}
