package initializers

import "github.com/ckng0221/url-shortener/models"

func SynDatabase() {
	Db.AutoMigrate(
		&models.Url{},
	)
}
