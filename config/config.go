package config

type Config interface {
	Get(key string) string
}
