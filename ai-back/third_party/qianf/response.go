package qianf

import (
	"awesomeProject3/api"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"strings"
)

func GenerateImage(resq ImageRequestBody) (string, error) {

	token, err := Token(api.QFAPIKey, api.QFAPISecretKey)
	accessToken := token.AccessToken
	if err != nil {
		log.Error().Str("step", "GenerateToken").Msg(err.Error())
		return "", err
	}

	request := resq

	err = request.Validate()
	if err != nil {
		return "", err
	}
	js, _ := json.Marshal(request)

	resp, err := http.Post(api.DIFFUSSION_XL_BASEURL+"?access_token="+accessToken, "application/json", strings.NewReader(string(js)))

	resp.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Error().Str("step", "GenerateImage").Msg(err.Error())
		return "", err
	}
	defer resp.Body.Close()

	//读取返回的数据
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
