package handler

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-blog/errno"
	"io"
	"net/http"
	"strconv"
	"time"
)

type OssConfig struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"accessKey_id"`
	AccessKeySecret string `json:"accessKey_secret"`
	Region          string `json:"region"`
	Bucket          string `json:"bucket"`
	Secure          bool   `json:"secure"`
	Cname           bool   `json:"cname"`
}

func Upload(c *gin.Context) {
	//通过form-data上传文件，文件名：file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}

	fileHandle, err := file.Open() //打开上传文件
	defer fileHandle.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}

	url, err := LocalUrl(fileHandle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errno.ConstructErrResp(string(rune(errno.ERROR)), err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": gin.H{
			"url": url,
		},
	})
}

func OssCennect() (Bucket *oss.Bucket, err error) {
	ossConfig := OssConfig{}
	//此处需要进入阿里云oss控制台配置域名
	ossConfig.Endpoint = "image-gem.oss-accelerate.aliyuncs.com"
	ossConfig.AccessKeyId = "LTAIaQk2r0874ARk"
	ossConfig.AccessKeySecret = "5hTbNyMyVT2E7uDNdzlm6ZVRVB238z"
	ossConfig.Bucket = "image-gem"
	ossConfig.Secure = true
	ossConfig.Cname = true
	client, err := oss.New(ossConfig.Endpoint, ossConfig.AccessKeyId, ossConfig.AccessKeySecret, oss.UseCname(true))
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	Bucket, err = client.Bucket("image-gem")
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return Bucket, nil
}

func LocalUrl(file io.Reader) (url string, err error) {
	Bucket, err := OssCennect()
	if err != nil {
		return "", err
	}

	t := time.Now()
	//拼接文件名称
	fileName := fmt.Sprintf("%s%s%s%s%s%s%s", strconv.Itoa(t.Year()), strconv.Itoa(int(t.Month())), strconv.Itoa(t.Day()), strconv.Itoa(t.Hour()), strconv.Itoa(t.Minute()), strconv.Itoa(t.Second()), strconv.Itoa(int(t.Unix())))
	str := "blog/" + fileName + ".jpg"

	err = Bucket.PutObject(str, file)
	if err != nil {
		url = ""
	} else {
		url = fmt.Sprintf("%s%s", "https://image-gem.oss-cn-shanghai.aliyuncs.com/", str)
	}
	return url, err
}
