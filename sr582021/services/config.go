package services

import (
	"fmt"
	"sr582021/model"
)

type ConfigService struct {
	repo model.ConfigRepository
}

func NewCOnfigservice(repo model.ConfigRepository) ConfigService {
	return ConfigService{
		repo: repo,
	}

}

func (s ConfigService) Hello() {
	fmt.Println("Hello config service")
}

func (c ConfigService) Get(name string, version int) (model.Config, error) {
	return c.repo.Get(name, version)
}
