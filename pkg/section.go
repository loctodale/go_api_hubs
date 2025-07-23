package pkg

type Config struct {
	AccountService  AccountService  `mapstructure:"account_service"`
	SendmailService SendmailService `mapstructure:"sendmail_service"`
	ApisService     ApisService     `mapstructure:"apis_service"`
	DiscordBotLogs  DiscordBotLogs  `mapstructure:"discord-logs-service"`
}

// Service Config
type AccountService struct {
	Ports    PortConfig             `mapstructure:"port"`
	Database AccountServiceDatabase `mapstructure:"database"`
	Kafka    KafkaConfig            `mapstructure:"kafka"`
}

type SendmailService struct {
	Ports    PortConfig       `mapstructure:"port"`
	Database SendmailDatabase `mapstructure:"database"`
	Kafka    KafkaConfig      `mapstructure:"kafka"`
	Email    Email            `mapstructure:"mail"`
}

type AuthService struct {
	Ports           PortConfig `mapstructure:"port"`
	JwksUri         string     `mapstructure:"jwks_uri"`
	AuthServiceHost string     `mapstructure:"auth_service_host"`
}

type ApisService struct {
	Ports    PortConfig          `mapstructure:"port"`
	Database ApisServiceDatabase `mapstructure:"database"`
	Kafka    KafkaConfig         `mapstructure:"kafka"`
}

type DiscordBotLogs struct {
	Token   string            `mapstructure:"token"`
	Channel map[string]string `mapstructure:"channel"`
}

// End Service config

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
	Poll     int    `mapstructure:"poll"`
}
type AccountServiceDatabase struct {
	Postgres string      `mapstructure:"postgres"`
	Redis    RedisConfig `mapstructure:"redis"`
}

type SendmailDatabase struct {
	Redis RedisConfig `mapstructure:"redis"`
}

type ApisServiceDatabase struct {
	Postgres string      `mapstructure:"postgres"`
	Redis    RedisConfig `mapstructure:"redis"`
}

type PortConfig struct {
	Local int `mapstructure:"local"`
	Prod  int `mapstructure:"prod"`
}

type KafkaConfig struct {
	Address string `mapstructure:"address"`
	Topic   string `mapstructure:"topic"`
}

type Email struct {
	Account  string `mapstructure:"account"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}

type ChannelConfig struct {
	Account  string `mapstructure:"account"`
	Apis     string `mapstructure:"apis"`
	Payments string `mapstructure:"payment"`
}
