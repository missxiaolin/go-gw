package qiniu

import (
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"go-gw/config"
)

// 获取七牛token
func GetToken() string {
	putPolicy := storage.PutPolicy{
		Scope: config.Cfg.QiNiu.Bucket,
	}
	mac := qbox.NewMac(config.Cfg.QiNiu.Ak, config.Cfg.QiNiu.Sk)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}