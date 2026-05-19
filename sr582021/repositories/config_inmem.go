package repositories

import (
	"errors"
	"sr582021/model"
	"strconv"
)

type ConfigInMemRepository struct {
	configs map[string]model.Config
	groups  map[string]model.ConfigGroup
}

func NewConfigInMemRepository() model.ConfigRepository {
	return ConfigInMemRepository{
		configs: make(map[string]model.Config),
		groups:  make(map[string]model.ConfigGroup),
	}
}

func (c ConfigInMemRepository) Add(config model.Config) {
	c.configs[config.Name+"/"+strconv.Itoa(config.Version)] = config
}

func (c ConfigInMemRepository) Get(name string, version int) (model.Config, error) {
	config, ok := c.configs[name+"/"+strconv.Itoa(version)]
	if !ok {
		return model.Config{}, errors.New("config not found")
	}
	return config, nil
}

func (c ConfigInMemRepository) GetAll() (map[string]model.Config, error) {
	return c.configs, nil
}

func (c ConfigInMemRepository) DeleteByVersion(name string, version int) (model.Config, error) {
	config, ok := c.configs[name+"/"+strconv.Itoa(version)]
	if !ok {
		return model.Config{}, errors.New("config not found")
	}
	delete(c.configs, name+"/"+strconv.Itoa(version))
	return config, nil
}

func (c ConfigInMemRepository) GetGroup(name string, version int) (model.ConfigGroup, error) {
	group, ok := c.groups[name+"/"+strconv.Itoa(version)]
	if !ok {
		return model.ConfigGroup{}, errors.New("config group not found")
	}
	return group, nil
}

func (c ConfigInMemRepository) GetAllGroups() (map[string]model.ConfigGroup, error) {
	return c.groups, nil
}

func (c ConfigInMemRepository) AddGroup(group model.ConfigGroup) {
	c.groups[group.Name+"/"+strconv.Itoa(group.Version)] = group
}

func (c ConfigInMemRepository) DeleteGroupByVersion(name string, version int) (model.ConfigGroup, error) {
	group, ok := c.groups[name+"/"+strconv.Itoa(version)]
	if !ok {
		return model.ConfigGroup{}, errors.New("config group not found")
	}
	delete(c.groups, name+"/"+strconv.Itoa(version))
	return group, nil
}

func (c ConfigInMemRepository) UpdateGroup(group model.ConfigGroup) error {
	_, ok := c.groups[group.Name+"/"+strconv.Itoa(group.Version)]
	if !ok {
		return errors.New("config group not found")
	}
	c.groups[group.Name+"/"+strconv.Itoa(group.Version)] = group
	return nil
}

func (c ConfigInMemRepository) PutGroup(group model.ConfigGroup, oldName string, oldVersion int) error {
	_, ok := c.groups[oldName+"/"+strconv.Itoa(oldVersion)]
	if !ok {
		return errors.New("config group not found")
	}
	delete(c.groups, oldName+"/"+strconv.Itoa(oldVersion))
	c.groups[group.Name+"/"+strconv.Itoa(group.Version)] = group
	return nil
}
