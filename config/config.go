package config

type Config struct {
	Mysql    Mysql  `yaml:"mysql"`
	Redis    Redis  `yaml:"redis"`
	LogLevel string `yaml:"loglevel"`
	Debug    bool   `yaml:"debug"`
	Env      string `yaml:"env"`
	Http     Http   `yaml:"http"`
}

type Mysql struct {
	User         string `yaml:"user"`
	MaxIdleConns int    `yaml:"max-idle-conns"`
	MaxOpenConns int    `yaml:"max-open-conns"`
}

type Redis struct {
	Addr         string `yaml:"addr"`
	Password     string `yaml:"password"`
	Db           int    `yaml:"db"`
	ReadTimeout  string `yaml:"read_timeout"`
	WriteTimeout string `yaml:"write_timeout"`
}

type Http struct {
	Port int `yaml:"port"`
}
