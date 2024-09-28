package initializers

import "github.com/orest-kostiuk/fiber-test/models"

func SyncDatabase() {
	err := DB.AutoMigrate(&models.Post{})
	if err != nil {
		panic(err)
		return
	}
}
