package migration

import (
	"errors"
	"os"

	"fmt"

	"github.com/jinzhu/gorm"
	db "github.com/tiket-dev/tiket-microservice-configuration/core/mysql"
	country "github.com/tiket-dev/tiket-microservice-configuration/databases/mysql/migration/country"
	logger "github.com/tiket-dev/tiket-microservice-configuration/helpers/logger"
)

func Init() {
	if len(os.Args) > 2 {
		dbConnection := db.ConnectMySQL()

		switch app := os.Args[2]; app {
		case "country":
			migrateCountry(dbConnection)
		default:
			msg := fmt.Sprintf("No app named %s", app)
			err := errors.New(msg)
			logger.Error(err, " Migration error")
		}
	}
}

func migrateCountry(dbConnection *gorm.DB) {
	logger.Info("Migrating country...")

	country.Down(dbConnection)
	country.Up(dbConnection)

	logger.Info("Migrating done.")
}
