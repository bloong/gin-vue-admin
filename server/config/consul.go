package config

type Consul struct {
	Addr  string `mapstructure:"addr" json:"addr" yaml:"addr"`    // 服务器地址:端口
	Token string `mapstructure:"token" json:"token" yaml:"token"` // token
}
