package api

import (
	"awesomeProject3/config"
	"github.com/sashabaranov/go-openai"
)

var Db, _ = config.GetConfig().MySQL.GetDb()

const (
	MODEL     = openai.GPT3Dot5Turbo
	MAXTOKENS = 2000
)

/*文心一格*/
const (
	DIFFUSSION_XL_BASEURL = ""
	QFBASEURL             = ""
	QFTOKENURL            = ""
	QFAPIKey              = ""
	QFAPISecretKey        = ""
)
