package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Redis    Redis    `yaml:"redis"`
	Log      Log      `yaml:"log"`
	LogLevel string   `yaml:"loglevel"`
	Debug    bool     `yaml:"debug"`
	Env      string   `yaml:"env"`
	Http     Http     `yaml:"http"`
	Security Security `yaml:"security"`
	Jwt      Jwt      `yaml:"jwt"`
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

type Log struct {
	LogLevel   string `yaml:"log_level"`
	Encoding   string `yaml:"encoding"`
	LogFileDir string `yaml:"log_file_dir"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
	MaxSize    int    `yaml:"max_size"`
	Compress   bool   `yaml:"compress"`
}

type Http struct {
	Port int `yaml:"port"`
}

type Security struct {
	AppKey      int `yaml:"app_key"`
	AppSecurity int `yaml:"app_security"`
}

type Jwt struct {
	Key string `yaml:"key"`
}
