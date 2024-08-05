package models

type Configuration struct {
	Server   Server   `yaml:"server" ENV:"SERVER"`
	Database Database `yaml:"database" ENV:"DATABASE"`
}

type Server struct {
	Port int    `yaml:"port" ENV:"SERVER_PORT"`
	Cors string `yaml:"cors" ENV:"SERVER_CORS"`
}

type Database struct {
	Server   string `yaml:"server" ENV:"DATABASE_SERVER"`
	Port     string `yaml:"port" ENV:"DATABASE_PORT"`
	User     string `yaml:"user" ENV:"DATABASE_USER"`
	Password string `yaml:"password" ENV:"DATABASE_PASSWORD"`
	Name     string `yaml:"name" ENV:"DATABASE_NAME"`
}
