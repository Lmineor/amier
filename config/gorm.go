package config

type Mysql struct {
	Path         string `json:"path" yaml:"path"`
	Config       string `json:"config" yaml:"config"`
	Dbname       string `json:"dbname" yaml:"dbname"`
	Username     string `json:"username" yaml:"username"`
	Password     string `json:"password" yaml:"password"`
	MaxIdleConns int    `json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `json:"logMode" yaml:"log-mode"`
	LogZap       string `json:"logZap" yaml:"log-zap"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
}
