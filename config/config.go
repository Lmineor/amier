package config

type Server struct {
	System System `json:"system" yaml:"system"`
	Mysql  Mysql  `json:"mysql" yaml:"mysql"`
	ZAP    Zap    `json:"zap" yaml:"zap"`
}
