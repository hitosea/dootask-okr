package config

type System struct {
	Mode  string `mapstructure:"mode"`
	Host  string `mapstructure:"host"`
	Port  string `mapstructure:"port"`
	Cache string `mapstructure:"cache"`
	Dsn   string `mapstructure:"dsn"`
	Start string `mapstructure:"start"`
}

type Jwt struct {
	SecretKey string `mapstructure:"secret_key"`
}

type Redis struct {
	RedisUrl      string `mapstructure:"redis_url"`  // Redis源配置
	RedisHost     string `mapstructure:"redis_host"` // Redis主机地址
	RedisPort     string `mapstructure:"redis_port"` // Redis密码
	RedisPassword string `mapstructure:"redis_pass"` // Redis密码
}

type ServerConfig struct {
	Redis  Redis  `mapstructure:"redis"`
	Jwt    Jwt    `mapstructure:"system"`
	System System `mapstructure:"system"`
}
