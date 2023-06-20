package config

type CommonConfig struct {
	Env    string
	Log    LogConfig
	Server ServerConfig
}

type LogConfig struct {
	Level string
}

type ServerConfig struct {
	Host string
	Port string
}
