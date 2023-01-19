package server

import (
	"GinBlog/utils"
	"GinBlog/utils/errmsg"
	"context"
	"fmt"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiniuServer // 图片上传路径

func UploadFile(fileData multipart.File, fileSize int64) (int, string) {
	fmt.Println(AccessKey, SecretKey, Bucket, ImgUrl)
	// 上传策略
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac) // 业务服务器颁发的上传凭证

	region, _ := storage.GetRegion(AccessKey, Bucket) // 获取所在的服务器区域

	cfg := storage.Config{
		UseHTTPS:      false, // 是否使用https域名
		UseCdnDomains: false, // 上传是否使用CDN上传加速
		Region:        region,
	}

	// 构建表单上传对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{} // 上传成功后返回的数据

	putExtra := storage.PutExtra{} // 可选配置
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, fileData, fileSize, &putExtra)
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR, ""
	}
	return errmsg.SUCCESS, ImgUrl + ret.Key

}
