package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

type FlagsConfig struct {
	Port    int
	Migrate bool
}

type Configs struct {
	DBConfig DBConfig
	Flags    FlagsConfig
}

// Singleton essentials
var (
	instance *Configs
	once     sync.Once
)

func GetConfig() *Configs {
	once.Do(func() {
		instance = &Configs{
			DBConfig: NewDBConfig(),
		}
	})
	return instance
}

func (c *Configs) Load() {
	godotenv.Load()

	c.ParseFlags()
	c.LoadDBConfig()
}

func (c *Configs) ParseFlags() {
	flag.IntVar(&c.Flags.Port, "port", 8080, "Port to run the server on")
	flag.BoolVar(&c.Flags.Migrate, "migrate", false, "Run with migration")

	flag.Parse()
}

func (c *Configs) LoadDBConfig() {
	cfg := &c.DBConfig

	// Add new db envs here
	requiredEnvsForDB := map[string]*string{
		"DB_HOST":     &cfg.Host,
		"DB_USER":     &cfg.User,
		"DB_PASSWORD": &cfg.Password,
		"DB_NAME":     &cfg.Name,
		"DB_PORT":     &cfg.Port,
	}
	var missing []string

	for field, valuePtr := range requiredEnvsForDB {
		value := os.Getenv(field)
		if value == "" {
			missing = append(missing, field)
		} else {
			*valuePtr = value
		}
	}

	if len(missing) > 0 {
		log.Fatalf("\nMissing required env vars:\n----\n%s\n----\n", strings.Join(missing, ",\n"))
	}
}

func (c *Configs) GetPortString() string {
	return fmt.Sprintf(":%s", strconv.Itoa(c.Flags.Port))
}
