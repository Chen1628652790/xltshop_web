package config

type ServerConfig struct {
	Host        string        `mapstructure:"host" json:"host"`
	Tags        []string      `mapstructure:"tags" json:"tags"`
	Name        string        `mapstructure:"name" json:"name"`
	Port        int           `mapstructure:"port" json:"port"`
	Mode        string        `mapstructure:"mode" json:"mode"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	JwtInfo     JwtConfig     `mapstructure:"jwt" json:"jwt"`
	ConsulInfo  ConsulConfig  `mapstructure:"consul" json:"consul"`
}

type UserSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type JwtConfig struct {
	SignatureKey string `mapstructure:"signature_key" json:"signature_key"`
	ExpireSecond int    `mapstructure:"expire_second" json:"expire_second"`
	ExpireCount  int    `mapstructure:"expire_count" json:"expire_count"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      int    `mapstructure:"port" json:"port"`
	Namespace string `mapstructure:"namespace" json:"namespace"`
	User      string `mapstructure:"user" json:"user"`
	Password  string `mapstructure:"password" json:"password"`
	DataID    string `mapstructure:"data_id" json:"data_id"`
	Group     string `mapstructure:"group" json:"group"`
}
