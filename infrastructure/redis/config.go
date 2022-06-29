package redis

var Config RedisConfig

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}
