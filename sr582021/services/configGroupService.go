package services

import (
	"errors"
	"sr582021/model"

	"github.com/google/uuid"
)

type ConfigGroupService struct {
	repo model.ConfigRepository
}

func NewConfigGroupService(repo model.ConfigRepository) ConfigGroupService {
	return ConfigGroupService{
		repo: repo,
	}
}

func (c ConfigGroupService) GetGroup(name string, version int) (model.ConfigGroup, error) {
	return c.repo.GetGroup(name, version)
}

func (c ConfigGroupService) GetAllGroups() (map[string]model.ConfigGroup, error) {
	return c.repo.GetAllGroups()
}

func (c ConfigGroupService) DeleteGroupByVersion(name string, version int) error {
	_, err := c.repo.DeleteGroupByVersion(name, version)
	return err
}

func (c ConfigGroupService) DeleteConfigByVersion(groupName string, groupVersion int, configName string, configVersion int) error {
	group, err := c.repo.GetGroup(groupName, groupVersion)
	if err != nil {
		return err
	}
	newConfigs := []model.Config{}
	for _, config := range group.Configs {
		if config.Name != configName || config.Version != configVersion {
			newConfigs = append(newConfigs, config)
		}
	}
	group.Configs = newConfigs
	return c.repo.UpdateGroup(group)
}

func (c ConfigGroupService) PutGroup(name string, version int, updatedGroup model.ConfigGroup) error {
	group, err := c.repo.GetGroup(name, version)
	if err != nil {
		return err
	}
	group.Name = updatedGroup.Name
	group.Version = updatedGroup.Version
	group.Configs = updatedGroup.Configs
	return c.repo.UpdateGroup(group)
}

func matchesLabels(config model.Config, query map[string]string) bool {
	for k, v := range query {
		if config.Labels[k] != v {
			return false
		}
	}
	return true
}

func (c ConfigGroupService) GetConfigsByLabels(groupName string, groupVersion int, query map[string]string) ([]model.Config, error) {
	group, err := c.repo.GetGroup(groupName, groupVersion)
	if err != nil {
		return nil, err
	}

	var result []model.Config
	for _, config := range group.Configs {
		if matchesLabels(config, query) {
			result = append(result, config)
		}
	}
	return result, nil
}

func (c ConfigGroupService) DeleteConfigsByLabels(groupName string, groupVersion int, query map[string]string) error {
	group, err := c.repo.GetGroup(groupName, groupVersion)
	if err != nil {
		return err
	}

	var remaining []model.Config
	for _, config := range group.Configs {
		if !matchesLabels(config, query) {
			remaining = append(remaining, config)
		}
	}
	group.Configs = remaining
	return c.repo.UpdateGroup(group)
}

func (c ConfigGroupService) PostGroup(name string, version int, configs []model.Config) error {
	_, err := c.repo.GetGroup(name, version)
	if err == nil {
		return errors.New("group already exists")
	}
	group := model.ConfigGroup{
		Id:      uuid.New().String(),
		Name:    name,
		Version: version,
		Configs: configs,
	}
	c.repo.AddGroup(group)
	return nil
}
