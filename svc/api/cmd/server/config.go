package main

import (
	"fmt"

	"github.com/peterbourgon/ff/v4"
	"github.com/peterbourgon/ff/v4/ffhelp"
	"github.com/peterbourgon/ff/v4/ffyaml"
)

type Config struct {
	Server Server
	DB     DB
}

type Server struct {
	Host string
	Port int
}

type DB struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
	SSL      string
}

func NewConfig(args []string) (*Config, error) {
	fs := ff.NewFlagSet("API Gateway")
	var (
		config     string
		serverHost string
		serverPort int
		dbHost     string
		dbPort     int
		dbName     string
		dbUser     string
		dbPasswd   string
		dbSSL      string
	)
	fs.StringEnumVar(&config, 0, "config", "environment configuration file", "env.yaml")
	fs.StringVar(&serverHost, 0, "server.host", "", "server host address to listen for incoming requests")
	fs.IntVar(&serverPort, 0, "server.port", 8080, "server port number to listen for incoming requests")
	fs.StringVar(&dbHost, 0, "db.host", "", "database host address")
	fs.IntVar(&dbPort, 0, "db.port", 5432, "database port number")
	fs.StringVar(&dbName, 0, "db.name", "", "database name")
	fs.StringVar(&dbUser, 0, "db.user", "", "database user")
	fs.StringVar(&dbPasswd, 0, "db.passwd", "", "database password")
	fs.StringVar(&dbSSL, 0, "db.ssl", "", "database ssl mode")

	if err := ff.Parse(fs, args[1:],
		ff.WithEnvVarPrefix("EFM_API_GATEWAY"),
		ff.WithConfigFileParser(ffyaml.Parse),
		ff.WithConfigFileFlag("config"),
		ff.WithConfigIgnoreUndefinedFlags(),
	); err != nil {
		fmt.Printf("%s\n", ffhelp.Flags(fs))
		fmt.Printf("ERROR\n%v\n", err)
		return &Config{}, err
	}

	return &Config{
		Server: Server{
			Host: serverHost,
			Port: serverPort,
		},
		DB: DB{
			Host:     dbHost,
			Port:     dbPort,
			Name:     dbName,
			User:     dbUser,
			Password: dbPasswd,
			SSL:      dbSSL,
		},
	}, nil
}
