package cos

import (
	"awesomeProject3/config"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"strings"
)

var u *url.URL
var b = &cos.BaseURL{}
var c *cos.Client

func init() {
	u, _ = url.Parse("")
	b = &cos.BaseURL{BucketURL: u}

	fmt.Println(config.GetConfig().Cos.SecretID)

	c = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.GetConfig().Cos.SecretID,
			SecretKey: config.GetConfig().Cos.SecretKey,
		},
	})
}

func SaveImageByBase64(imageCode string, name string) (string, error) {
	decodeString, err := base64.StdEncoding.DecodeString(imageCode)

	if err != nil {
		log.Error().Str("step", "at decodeString").Msg(err.Error())
		return "", err
	}

	fd := strings.NewReader(string(decodeString))
	_, err = c.Object.Put(context.Background(), name, fd, nil)

	if err != nil {
		log.Error().Str("step", "at c.Object.Put").Msg(err.Error())
		panic(err)
	}
	log.Info().Str("step", "at c.Object.Put").Str("url", config.GetConfig().Cos.BaseURL+"/"+name).Msg("success")
	//返回图片地址
	return config.GetConfig().Cos.BaseURL + "/" + name, nil
}

func SaveFileByBase64(fileCode string, name string) (string, error) {
	decodeString, err := base64.StdEncoding.DecodeString(fileCode)

	if err != nil {
		log.Error().Str("step", "at decodeString").Msg(err.Error())
		return "", err
	}

	fd := strings.NewReader(string(decodeString))
	_, err = c.Object.Put(context.Background(), name, fd, nil)

	if err != nil {
		log.Error().Str("step", "at c.Object.Put").Msg(err.Error())
		panic(err)
	}
	log.Info().Str("step", "at c.Object.Put").Str("url", config.GetConfig().Cos.BaseURL+"/"+name).Msg("success")
	//返回文件地址
	return config.GetConfig().Cos.BaseURL + "/" + name, nil
}
