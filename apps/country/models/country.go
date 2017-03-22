package models

import (
	"time"

	db "github.com/tiket-dev/tiket-microservice-configuration/core/mysql"
)

type Country struct {
	ID          int        `json:"id"`
	CountryCode string     `gorm:"size:2" json:"country_code"`
	CountryName string     `json:"country_name"`
	PhoneCode   string     `gorm:"size:4" json:"phone_code"`
	Icon        string     `json:"icon"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
}

func GetAllCountry() []*Country {
	var countries []*Country

	dbConnection := db.ConnectMySQL()

	dbConnection.Find(&countries).Select("id, country_code, country_name, phone_code, icon")

	return countries
}
