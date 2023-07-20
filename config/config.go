package config

type Server struct {
	Debug      bool       `mapstructure:"debug" json:"debug" yaml:"debug"`
	Mysql      Mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis      Redis      `mapstructure:"redis" json:"redis" yaml:"redis"`
	System     System     `mapstructure:"system" json:"system" yaml:"system"`
	LogLevel   string     `mapstructure:"loglevel" json:"loglevel" yaml:"loglevel"`
	Aws        Aws        `mapstructure:"aws" json:"aws" yaml:"aws"`
	BlockChain BlockChain `mapstructure:"blockchain" json:"blockchain" yaml:"blockchain"`
	Match      Match      `mapstructure:"match" json:"match" yaml:"match"`
	Horse      Horse      `mapstructure:"horse" json:"horse" yaml:"horse"`
}

type System struct {
	Addr   int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	Domain string `mapstructure:"domain" json:"domain" yaml:"domain"`
}

type Mysql struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
}

type Aws struct {
	BucketName string `mapstructure:"bucket-name" json:"bucket-name" yaml:"bucket-name"`
	Endpoint   string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	Region     string `mapstructure:"region" json:"region" yaml:"region"`
	Key        string `mapstructure:"key" json:"key" yaml:"key"`
	Secret     string `mapstructure:"secret" json:"secret" yaml:"secret"`
}

type BlockChain struct {
	PrivateKey    string `mapstructure:"private-key" json:"private-key" yaml:"private-key"`
	MintAddress   string `mapstructure:"mint-address" json:"mint-address" yaml:"mint-address"`
	HorseNft      string `mapstructure:"horse-nft" json:"horse-nft" yaml:"horse-nft"`
	BlindBoxNft   string `mapstructure:"blind-box-nft" json:"blind-box-nft" yaml:"blind-box-nft"`
	ChainId       int64  `mapstructure:"chain-id" json:"chain-id" yaml:"chain-id"`
	Gateway       string `mapstructure:"gateway" json:"gateway" yaml:"gateway"`
	Arena         string `mapstructure:"arena" json:"arena" yaml:"arena"`
	PledgeAddress string `mapstructure:"pledge-address" json:"pledge-address" yaml:"pledge-address"`
}
type Match struct {
	UserSignMaxCount int `mapstructure:"user-sign-max-count" json:"user-sign-max-count" yaml:"user-sign-max-count"`
}
type Horse struct {
	Grade []int `mapstructure:"grade" json:"grade" yaml:"grade"`
}
