package stx

import "bilibiliSpider/etc"

type StcConfig struct {
	UPList etc.UPList
}

func NewSTXContext() *StcConfig {
	return &StcConfig{
		UPList: etc.NewUPList("./etc/config.yaml"),
	}
}
