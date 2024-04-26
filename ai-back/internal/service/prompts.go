package service

import (
	"awesomeProject3/api"
	"encoding/json"
	"os"
)

type PromptService interface {
	AddPrompt(prompts *api.Prompt) error
	DeletePrompt(id string) error
	GetPromptList(prompts *[]api.Prompt, offset int) error
	UpdatePrompt(prompt *api.Prompt) error
}

type promptService struct {
}

func init() {
	go loadJsonData()
}

func loadJsonData() {
	if api.Db.Migrator().HasTable(&api.Prompt{}) {
		return
	}

	api.Db.AutoMigrate(&api.Prompt{})
	file, err := os.Open("prompts-zh.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var prompts []api.Prompt
	json.NewDecoder(file).Decode(&prompts)
	for _, prompt := range prompts {
		api.Db.Create(&prompt)
	}
}

func (p promptService) AddPrompt(prompts *api.Prompt) error {
	api.Db.Create(prompts)
	return nil
}

func (p promptService) DeletePrompt(id string) error {
	//先看看这个id是否存在
	err := api.Db.First(&api.Prompt{}, id)
	if err.Error != nil {
		return err.Error
	}

	err = api.Db.Delete(&api.Prompt{}, id)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (p promptService) GetPromptList(prompts *[]api.Prompt, offset int) error {
	api.Db.Limit(10).Offset(offset * 10).Find(&prompts)
	return nil
}

func (p promptService) UpdatePrompt(prompt *api.Prompt) error {
	api.Db.Save(prompt)
	return nil
}

func NewPromptService() PromptService {
	return &promptService{}
}
