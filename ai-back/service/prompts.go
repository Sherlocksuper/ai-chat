package service

import "awesomeProject3/api"

type PromptService interface {
	AddPrompt(prompts *api.Prompt) error
	DeletePrompt(id string) error
	GetPromptList(prompts *[]api.Prompt) error
	UpdatePrompt(prompt *api.Prompt) error
}

type promptService struct {
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

func (p promptService) GetPromptList(prompts *[]api.Prompt) error {
	api.Db.Find(&prompts)
	return nil
}

func (p promptService) UpdatePrompt(prompt *api.Prompt) error {
	api.Db.Save(prompt)
	return nil
}

func NewPromptService() PromptService {
	return &promptService{}
}
