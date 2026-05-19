package repositories

import (
	"sr582021/model"
)

type ConfigConsulRepository struct {
}

// AddGroup implements [model.ConfigRepository].
func (c ConfigConsulRepository) AddGroup(group model.ConfigGroup) {
	panic("unimplemented")
}

// DeleteByVersion implements [model.ConfigRepository].
func (c ConfigConsulRepository) DeleteByVersion(name string, version int) (model.Config, error) {
	panic("unimplemented")
}

// DeleteGroupByVersion implements [model.ConfigRepository].
func (c ConfigConsulRepository) DeleteGroupByVersion(name string, version int) (model.ConfigGroup, error) {
	panic("unimplemented")
}

// GetAll implements [model.ConfigRepository].
func (c ConfigConsulRepository) GetAll() (map[string]model.Config, error) {
	panic("unimplemented")
}

// GetAllGroups implements [model.ConfigRepository].
func (c ConfigConsulRepository) GetAllGroups() (map[string]model.ConfigGroup, error) {
	panic("unimplemented")
}

// GetGroup implements [model.ConfigRepository].
func (c ConfigConsulRepository) GetGroup(name string, version int) (model.ConfigGroup, error) {
	panic("unimplemented")
}

// PutGroup implements [model.ConfigRepository].
func (c ConfigConsulRepository) PutGroup(group model.ConfigGroup, oldName string, oldVersion int) error {
	panic("unimplemented")
}

// UpdateGroup implements [model.ConfigRepository].
func (c ConfigConsulRepository) UpdateGroup(group model.ConfigGroup) error {
	panic("unimplemented")
}

func (ConfigConsulRepository) Add(config model.Config) {
	panic("unimplemented")
}

func (ConfigConsulRepository) Get(name string, version int) (model.Config, error) {
	panic("unimplemented")
}

func NewConfigConsulRepository() model.ConfigRepository {
	return ConfigConsulRepository{}
}
