package config

type GoodsSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type UserOPSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ServerConfig struct {
	Name            string          `mapstructure:"name" json:"name"`
	Port            int             `mapstructure:"port" json:"port"`
	Host            string          `mapstructure:"host" json:"host"`
	Tags            []string        `mapstructure:"tags" json:"tags"`
	Mode            string          `mapstructure:"mode" json:"mode"`
	GoodsSrvInfo    GoodsSrvConfig  `mapstructure:"goods_srv" json:"goods_srv"`
	UserOPSrvConfig UserOPSrvConfig `mapstructure:"userop_srv" json:"userop_srv"`
	JwtInfo         JWTConfig       `mapstructure:"jwt" json:"jwt"`
	ConsulInfo      ConsulConfig    `mapstructure:"consul" json:"consul"`
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
