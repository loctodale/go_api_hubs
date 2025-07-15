package initialize

func Run() {
	LoadConfig()
	InitPostgresServer()
	InitRedisServer()
	InitKafkaServer()
}
