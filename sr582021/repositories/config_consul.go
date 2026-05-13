package repositories

import (
	
	"sr582021/model"
)

type ConfigConsulRepository struct {

}

func (ConfigConsulRepository) Add(config model.Config){
	panic("unimplemented")
}

func (ConfigConsulRepository) Get(name string, version int) (model.Config, error){
	panic("unimplemented")
}

func NewConfigConsulRepository() model.ConfigRepository{
	return ConfigConsulRepository{}
}