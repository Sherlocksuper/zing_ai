package service

import (
	"awesomeProject3/api"
	"awesomeProject3/third_party/cos"
	qianf2 "awesomeProject3/third_party/qianf"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
	"strconv"
	"time"
)

var location = "service/chat.go"

type Chat = api.Chat

type ChatService interface {
	// StartAChat / 增删改查
	StartAChat(chat *api.Chat) error

	DeleteChat(id string) error

	DeleteAllChat(userId string) error

	GetChatDetail(id string) (Chat, error)

	GetChatList(chats *[]Chat, userId string) error

	SendMessage(context *gin.Context, chatId string, message string) error

	SendForImage(chatId string, message string) (string, error)
}

type chatService struct{}

func (c chatService) StartAChat(chat *api.Chat) error {
	title := chat.Title
	userId := chat.UserID

	log.Info().Msg("用户" + strconv.Itoa(int(userId)) + "开始一个" + title + "的聊天" + "location is :service/chat.go  StartAChat")
	chat.Messages = []api.Message{}

	api.Db.Create(&chat)
	return nil
}

// DeleteChat delete chat
func (c chatService) DeleteChat(id string) error {
	chat := Chat{}
	api.Db.Find(&chat, id)
	if chat.ID == 0 {
		return errors.New("chat not found")
	}
	api.Db.Unscoped().Delete(&Chat{}, id)
	return nil
}

// DeleteAllChat delete all chat
func (c chatService) DeleteAllChat(userId string) error {
	//打印
	log.Info().Msg("用户" + userId + "删除所有聊天" + "location is :service/chat.go  DeleteAllChat")
	Idint, _ := strconv.Atoi(userId)
	err := api.Db.Where("user_id = ?", Idint).Delete(&Chat{})
	if err.Error != nil {
		return errors.New("delete chat failed")
	}
	return nil
}

// GetChatDetail 获取聊天详情
func (c chatService) GetChatDetail(id string) (Chat, error) {
	chat := Chat{}
	err := api.Db.Model(&Chat{}).Preload("Messages").Find(&chat, id)
	if chat.ID == 0 || err.RowsAffected == 0 {
		return chat, errors.New("chat not found")
	}
	log.Info().Msg("chat is :" + chat.Title + "location is :service/chat.go  GetChatDetail")
	return chat, nil
}

// GetChatList 获取聊天列表
func (c chatService) GetChatList(chats *[]Chat, userId string) error {
	Idint, _ := strconv.Atoi(userId)
	err := api.Db.Model(&Chat{}).Preload("Messages", func(db *gorm.DB) *gorm.DB {
		return api.Db.Order("id desc")
	}).Where("user_id = ?", Idint).Find(&chats)
	if err.RowsAffected == 0 {
		return errors.New("chat not found")
	}
	return nil
}

func (c chatService) SendMessage(context *gin.Context, chatId string, message string) error {

	log.Info().Msg("chatId is :" + chatId + "location is :service/chat.go  SendMessage")
	chat := Chat{}
	err := api.Db.Model(&Chat{}).Preload("Messages").Find(&chat, chatId)
	if err.Error != nil || err.RowsAffected == 0 {
		return errors.New("chat not found")
	}

	//数据库保存信息
	chat.Messages = append(chat.Messages, api.Message{Role: openai.ChatMessageRoleUser, Content: message})
	response := streamMessages(buildOpenAIMessages(&chat.Messages), context)
	chat.Messages = append(chat.Messages, api.Message{Role: openai.ChatMessageRoleAssistant, Content: response})
	err = api.Db.Save(&chat)
	log.Info().Msg(message + "   is saved to database" + "location is :service/chat.go  SendMessage")
	log.Info().Msg(response + "   is saved to database" + "location is :service/chat.go  SendMessage")

	if err.Error != nil {
		return errors.New("send message failed")
	}
	return nil
}

func (c chatService) SendForImage(chatId string, message string) (string, error) {

	log.Info().Str("chatId", chatId).Str("message", message).Msg("SendForImage")

	resp, err := qianf2.GenerateImage(qianf2.ImageRequestBody{Prompt: message})

	if err != nil {
		log.Error().Str("step", "generate image").Msg(err.Error())
		return "", err
	}

	//把response转换成json
	var res map[string]interface{}
	err = json.Unmarshal([]byte(resp), &res)

	if err != nil {
		log.Error().Str("step", "unmarshal response").Msg(err.Error())
		return "", err
	}

	image := res["data"].([]interface{})[0].(map[string]interface{})["b64_image"].(string)

	//上传到云
	//获取现在的时间
	timeNow := time.Now()
	path, err := cos.SaveImageByBase64(image, chatId+"/"+timeNow.Format("2006-01-02-15-04-05")+".png")

	if err != nil {
		log.Error().Str("step", "save image to cos").Msg(err.Error())
		return "", err
	}

	chat := Chat{}
	api.Db.Model(&Chat{}).Preload("Messages").Find(&chat, chatId)
	if err != nil {
		return "", errors.New("chat not found")
	}
	chat.Messages = append(chat.Messages, api.Message{Role: openai.ChatMessageRoleUser, Content: message})
	chat.Messages = append(chat.Messages, api.Message{Role: openai.ChatMessageRoleAssistant, Content: path})
	api.Db.Save(&chat)
	if err != nil {
		log.Error().Str("location", location).Msg("save message failed")
		return "", errors.New("save message failed")
	}

	return path, nil
}

// SaveAIResponse 保存ai回复 到数据库
func (c chatService) SaveAIResponse(chatId string, content string) {
	chat := Chat{}
	api.Db.Find(&chat, chatId)
	if chat.ID == 0 {
		return
	}
	chat.Messages = append(chat.Messages, api.Message{Role: openai.ChatMessageRoleAssistant, Content: content})
	err := api.Db.Save(&chat)
	if err.Error != nil {
		return
	}
	return
}

// NewChatService 创建一个chatService
func NewChatService() ChatService {
	return &chatService{}
}
