package core

type Config struct {
	Message string
}

func NewConfig() *Config {
	return &Config{
		Message: "Hello from config!",
	}
}
