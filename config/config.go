package config

type Config struct {
	Port string
	AdminKey string
	JWTSecret string
}

func NewConfig() *Config {
	return &Config{
		Port: ":3000",
		AdminKey: "mdcyuldf834978fgw][drjjnejh76r3jevwneqwe@WEQREFQ@373t58dggt",
		JWTSecret: "mdcyu76r3jevwneqwe@WEQREFQ@373t58dggt",
	}
}