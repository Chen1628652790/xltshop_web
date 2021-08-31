package config

type ServerConfig struct {
	Name        string        `mapstructure:"name" json:"name"`
	Port        int           `mapstructure:"port" json:"port"`
	Mode        string        `mapstructure:"mode" json:"mode"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	JwtInfo     JwtConfig     `mapstructure:"jwt" json:"jwt"`
}

type UserSrvConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type JwtConfig struct {
	SignatureKey string `mapstructure:"signature_key" json:"signature_key"`
	ExpireSecond int    `mapstructure:"expire_second" json:"expire_second"`
	ExpireCount  int    `mapstructure:"expire_count" json:"expire_count"`
}
