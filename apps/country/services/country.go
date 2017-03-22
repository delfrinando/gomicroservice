package apps

import (
	"encoding/json"
	"log"

	countryModel "github.com/tiket-dev/tiket-microservice-configuration/apps/country/models"
	logger "github.com/tiket-dev/tiket-microservice-configuration/helpers/logger"
)

func GetCountry(msg string) []byte {

	type countryResponse struct {
		Countries []*countryModel.Country `json:"countries"`
	}

	countries := countryModel.GetAllCountry()
	log.Println("rowInService: ", countries)

	jsResult, err := json.Marshal(countryResponse{Countries: countries})

	if err != nil {
		logger.Error(err, "Failed on GetCountry")
	}

	return jsResult
}
