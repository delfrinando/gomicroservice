package seeder

import (
	"io/ioutil"

	"encoding/json"

	"github.com/jinzhu/gorm"
	country "github.com/tiket-dev/tiket-microservice-configuration/apps/country/models"
	logger "github.com/tiket-dev/tiket-microservice-configuration/helpers/logger"
)

func Seed(db *gorm.DB) {
	path := "./databases/mysql/seeder/country/country.json"

	data, err := ioutil.ReadFile(path)
	logger.Error(err, "Failed to open file")

	str := string(data)
	countries := []country.Country{}

	json.Unmarshal([]byte(str), &countries)

	for _, country := range countries {
		db.Create(&country)
	}
}
