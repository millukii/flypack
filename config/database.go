package config

type DBConfig struct {
	Host              string
	Port              string
	User              string
	DBName            string
	Password          string
	DBType            string
	MaxIddle          int
	MaxOpen           int
	MaxLifeTime       string
	WriteTimeout      string
	ReadTimeout       string
	ConnectionTimeout string
}
