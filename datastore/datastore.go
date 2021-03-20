package datastore

import (
	"github.com/maxstanley/fast_finder/models"
	"gorm.io/gorm"
)

// database contains the applications datastore connection
var database *gorm.DB

// NewGormDatastore creates a connection to the datastore using gorm.
func NewGormDatastore(
	connector gorm.Dialector,
	options ...gorm.Option,
) (*gorm.DB, error) {
	var option gorm.Option
	if options == nil {
		option = &gorm.Config{}
	} else {
		option = options[0]
	}

	var err error
	database, err = gorm.Open(connector, option)

	// Automatically add schemas where not correct.
	database.AutoMigrate(
		&models.Shortcut{},
	)

	return database, err
}

// Connection returns the gorm datastore connection.
func Connection() *gorm.DB {
	return database
}
