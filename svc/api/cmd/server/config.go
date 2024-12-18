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
	Host    string
	Port    string
	Timeout ServerTimeout
}

type ServerTimeout struct {
	Read     int
	Write    int
	Idle     int
	Shutdown int
}

type DB struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	SSL      string
	DSN      string
}

func NewConfig(args []string) (*Config, error) {
	fs := ff.NewFlagSet("API Gateway")
	var (
		config                string
		serverHost            string
		serverPort            string
		serverTimeoutRead     int
		serverTimeoutWrite    int
		serverTimeoutIdle     int
		serverTimeoutShutdown int
		dbHost                string
		dbPort                string
		dbName                string
		dbUser                string
		dbPasswd              string
		dbSSL                 string
	)
	fs.StringEnumVar(&config, 0, "config", "environment configuration file", "env.yaml")
	fs.StringVar(&serverHost, 0, "server.host", "", "server host address to listen for incoming requests")
	fs.StringVar(&serverPort, 0, "server.port", "", "server port number to listen for incoming requests")
	fs.IntVar(&serverTimeoutRead, 0, "server.timeout.read", 15, "number of seconds that the server will wait for reading requests")
	fs.IntVar(&serverTimeoutWrite, 0, "server.timeout.write", 15, "number of seconds that the server will wait for writing requests")
	fs.IntVar(&serverTimeoutIdle, 0, "server.timeout.idle", 60, "number of seconds that the server will wait for the next request")
	fs.IntVar(&serverTimeoutShutdown, 0, "server.timeout.shutdown", 30, "the duration for which the server gracefully wait for existing connections to finish")
	fs.StringVar(&dbHost, 0, "db.host", "", "database host address")
	fs.StringVar(&dbPort, 0, "db.port", "", "database port number")
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
			Timeout: ServerTimeout{
				Read:     serverTimeoutRead,
				Write:    serverTimeoutWrite,
				Idle:     serverTimeoutIdle,
				Shutdown: serverTimeoutShutdown,
			},
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
