/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package migrations

import (
	"github.com/codoworks/go-boilerplate/pkg/db/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	m := &gormigrate.Migration{}

	m.ID = "2022081802_create_users_table"

	m.Migrate = func(db *gorm.DB) error {
		type User struct {
			models.ModelBase
			Username  string `gorm:"size:255"`
			Password  string `gorm:"size:255"`
			Firstname string `gorm:"size:255"`
			Lastname  string `gorm:"size:255"`
			Email     string `gorm:"size:255"`
		}

		return AutoMigrateAndLog(db, &User{}, m.ID)
	}

	m.Rollback = func(db *gorm.DB) error {
		if err := db.Migrator().DropTable("users"); err != nil {
			logFail(m.ID, err, true)
		}
		logSuccess(m.ID, true)
		return nil
	}

	AddMigration(m)
}
