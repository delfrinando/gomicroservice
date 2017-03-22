package migration

import (
	"github.com/jinzhu/gorm"
	country "github.com/tiket-dev/tiket-microservice-configuration/apps/country/models"
)

func Up(db *gorm.DB) {
	db.AutoMigrate(&country.Country{})
}

func Down(db *gorm.DB) {
	if db.HasTable(&country.Country{}) {
		db.DropTable(&country.Country{})
	}
}
