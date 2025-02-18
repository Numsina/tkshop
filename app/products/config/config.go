package config

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	UserName string `mapstructure:"username" json:"username"`
	PassWord string `mapstructure:"password" json:"password"`
	DBName   string `mapstructure:"dbname" json:"dbname"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	PassWord string `mapstructure:"password" json:"password"`
}

type Config struct {
	MysqlInfo MysqlConfig `mapstructure:"mysql" json:"mysql"`
	RedisInfo RedisConfig `mapstructure:"redis" json:"redis"`
}
