package config

type Config struct {
	Redis    RedisConfig    `yaml:"redis"`
	Mysql    MysqlConfig    `yaml:"mysql"`
	Nacos    NacosConfig    `yaml:"nacos"`
	RabbitMQ RabbitMQConfig `yaml:"rabbitmq"`
	Cas      CasConfig      `yaml:"cas"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       int    `yaml:"db"`
	Password string `yaml:"password"`
}

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type NacosConfig struct {
	Host string `yaml:"host"`
	Port uint64 `yaml:"port"`
}

type RabbitMQConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Hostname string `yaml:"hostname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`
}

type CasConfig struct {
	Enable          bool   `yaml:"enable"`
	LoginUrl        string `yaml:"loginUrl"`
	IgnorePattern   string `yaml:"ignorePattern"`
	ServerUrlPrefix string `yaml:"serverUrlPrefix"`
	ServerLoginUrl  string `yaml:"serverLoginUrl"`
	ServerLogoutUrl string `yaml:"serverLogoutUrl"`
	ClientHostUrl   string `yaml:"clientHostUrl"`
	CallBackUrl     string `yaml:"callBackUrl"`
}
