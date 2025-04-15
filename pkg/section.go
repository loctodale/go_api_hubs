package pkg

type Config struct {
	AccountService AccountService `mapstructure:"account_service"`
}

type AccountRedis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
	Poll     int    `mapstructure:"poll"`
}
type AccountServiceDatabase struct {
	Postgres string       `mapstructure:"postgres"`
	Redis    AccountRedis `mapstructure:"redis"`
}
type AccountPort struct {
	Local int `mapstructure:"local"`
	Prod  int `mapstructure:"prod"`
}
type AccountService struct {
	Ports    AccountPort            `mapstructure:"port"`
	Database AccountServiceDatabase `mapstructure:"database"`
}
