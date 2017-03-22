package main

import (
	"errors"
	"fmt"
	"os"

	migrate "github.com/tiket-dev/tiket-microservice-configuration/databases/mysql/migration"
	seed "github.com/tiket-dev/tiket-microservice-configuration/databases/mysql/seeder"
	logger "github.com/tiket-dev/tiket-microservice-configuration/helpers/logger"
	"github.com/tiket-dev/tiket-microservice-configuration/routes"
)

func main() {
	if len(os.Args) > 1 {
		switch arg := os.Args[1]; arg {
		case "migrate":
			migrate.Init()
		case "seed":
			seed.Init()
		default:
			msg := fmt.Sprintf("No command named %s", arg)
			err := errors.New(msg)
			logger.Error(err, "Command error")
		}
	} else {
		routes.Routing()
	}
}
