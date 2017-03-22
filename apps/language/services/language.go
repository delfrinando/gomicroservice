package apps

import (
	response "github.com/tiket-dev/tiket-microservice-configuration/helpers/response"
)

func GetLanguage(msg string) []byte {
	languages := []response.Language{
		response.Language{
			Label: "IDR",
		},
	}

	result := response.SendSuccess(response.Http200Ok(), languages)

	return result
}
