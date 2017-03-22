package models

import (
db "../../../core/mysql"
	// "database/sql"
"log"
"fmt"
)

type Country struct {
	id 				int `json:"Country.id"`
	country_code	string `json:"Country.country_code"`
	country_name	string `json:"Country.country_name"`
	phone_code 		string `json:"Country.phone_code"`
	icon			string `json:"Country.icon"`
}

func GetAllCountry() []*Country{

	dbConnection := db.ConnectMySQL()

	rows, err := dbConnection.Query("SELECT id, country_code, country_name, phone_code, icon FROM country;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	country := new(Country)
	var countries []*Country
	for rows.Next() {
		err := rows.Scan(&country.id, &country.country_code, &country.country_name, &country.phone_code, &country.icon)
		if err != nil {
			log.Fatal(err)
		}
		countries = append(countries, country)
		fmt.Println(country)
	}
	return countries
}