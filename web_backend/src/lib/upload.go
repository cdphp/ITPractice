package lib

import (
	"context"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

// Uploader
type Uploader struct {
	AccessKey string
	SecretKey string
	Bucket    string
}

// Upload 上传
func (u *Uploader) Upload(localFile, key string) (string, error) {

	putPolicy := storage.PutPolicy{
		Scope: u.Bucket,
	}
	mac := qbox.NewMac(u.AccessKey, u.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}

	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {

		return "", err

	}
	return ret.Key, nil
}
