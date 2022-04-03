package model

import (
	"context"
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiniuSever

func UpLoadFile(file multipart.File, fileSize int64) (string, int) { //入参文件流和文件大小 出参url 状态码
	PutPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey) //详见七牛文档
	upToken := PutPolicy.UploadToken(mac)    //获取上传凭证

	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	PutExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{} //回调 默认回调格式为{"hash":"Ftgm-CkWePC9fzMBTRNmPMhGBcSV","key":"qiniu.jpg"} 一个Hash一个key

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &PutExtra)
	if err != nil {
		return "", errmsg.ERROR
	}
	url := ImgUrl + ret.Key
	return url, errmsg.SUCCESS
}
