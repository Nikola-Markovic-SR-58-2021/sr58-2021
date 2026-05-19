package services

import (
	"errors"
	"fmt"
	"sr582021/model"

	"github.com/google/uuid"
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

func (c ConfigService) Add(config model.Config) {
	c.repo.Add(config)
}

func (c ConfigService) GetAll() (map[string]model.Config, error) {
	return c.repo.GetAll()
}

func (c ConfigService) Post(name string, version int, params map[string]string) error {
	// Proveri da li konfiguracija već postoji zbog idempotentnosti
	_, err := c.repo.Get(name, version)
	if err == nil {
		return errors.New("config already exists, delete and recreate to replace")
	}
	config := model.Config{
		Id:      uuid.New().String(),
		Name:    name,
		Version: version,
		Params:  params,
	}
	c.repo.Add(config)
	return nil
}

func (c ConfigService) DeleteByVersion(name string, version int) error {
	c.repo.DeleteByVersion(name, version)
	return nil
}

func (c ConfigService) Replace(name string, version int, params map[string]string) error {
	_, err := c.repo.Get(name, version)
	if err != nil {
		return errors.New("config not found")
	}
	c.repo.DeleteByVersion(name, version)
	config := model.Config{
		Id:      uuid.New().String(),
		Name:    name,
		Version: version,
		Params:  params,
	}
	c.repo.Add(config)
	return nil
}
