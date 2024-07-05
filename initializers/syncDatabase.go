package initializers

import "url-shortener/models"

func SynDatabase() {
	Db.AutoMigrate(
		&models.Url{},
	)
}
