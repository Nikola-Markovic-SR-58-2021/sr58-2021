package model

type Config struct {
	Name    string            `json:"name"`
	Version int               `json:"version"`
	Params  map[string]string `json:"params"`
	Id      string            `json:"id"`
	Labels  map[string]string `json:"labels"`
}

type ConfigGroup struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Version int      `json:"version"`
	Configs []Config `json:"configs"`
}

type ConfigRepository interface {
	Add(config Config)
	Get(name string, version int) (Config, error)
	GetAll() (map[string]Config, error)
	DeleteByVersion(name string, version int) (Config, error)
	GetGroup(name string, version int) (ConfigGroup, error)
	GetAllGroups() (map[string]ConfigGroup, error)
	AddGroup(group ConfigGroup)
	DeleteGroupByVersion(name string, version int) (ConfigGroup, error)
	UpdateGroup(group ConfigGroup) error
	PutGroup(group ConfigGroup, oldName string, oldVersion int) error
}
