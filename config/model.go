package config

type config struct {
	Postgres    Postgres    `mapstructure:"postgres"`
	Server      Server      `mapstructure:"server"`
	ExternalURL ExternalURL `mapstructure:"externalURL"`
}

type Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

type Server struct {
	Port string
}

type ExternalURL struct {
	ProductAPI string
}
