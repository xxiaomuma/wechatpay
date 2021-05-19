package config

type Config struct {
	Wx *WxConfig `yaml:"wx"`
}

type WxConfig struct {
	AppId               string `yaml:"appId"`
	MchId               string `yaml:"mchId"`
	ApiKey              string `yaml:"apiKey"`
	V3ApiKey            string `yaml:"v3ApiKey"`
	CertificateSerialNo string `yaml:"certificateSerialNo"`
	CertPath            string `yaml:"certPath"`
	KeyPemPath          string `yaml:"keyPemPath"`
	CertPemPath         string `yaml:"certPemPath"`
	NotifyUrl           string `yaml:"notifyUrl"`
}
