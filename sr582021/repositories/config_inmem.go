package repositories

import (
	"sr582021/model"
	"strconv"
	"errors"
)

type ConfigInMemRepository struct {
	configs map[string]model.Config
}

func NewConfigInMemRepository() model.ConfigRepository{
	return ConfigInMemRepository {
		configs: make(map[string]model.Config),
	}
}

func (c ConfigInMemRepository) Add(config model.Config){
	c.configs[config.Name+"/"+strconv.Itoa(config.Version)] = config
}

func (c ConfigInMemRepository) Get(name string, version int) (model.Config, error){
	config, ok := c.configs[name+"/"+strconv.Itoa(version)]
	if !ok{
		return model.Config{}, errors.New("config not found")
	}
	return config, nil
}