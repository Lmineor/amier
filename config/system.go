package config

type System struct {
	Mode   string `json:"mode" yaml:"mode"`
	Port   string `json:"port" yaml:"port"`
	DbType string `json:"dbtype" yaml:"dbtype"`
}
