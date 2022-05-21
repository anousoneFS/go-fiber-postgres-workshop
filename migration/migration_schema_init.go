package migration

import (
	"fmt"

	"github.com/anousoneFS/administrative-divisions/repository"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	var err error
	err = db.AutoMigrate(&repository.ProvinceDB{}, &repository.DistrictDB{}, &repository.VillageDB{})
	if err != nil {
		return fmt.Errorf("Failed to migrate database: %v", err)
	}
	return nil
}
