package common

import (
	"github.com/dayvillefire/tenders/config"
	"github.com/gobuffalo/pop"
)

var (
	// DB is the shared pop database connection
	DB *pop.Connection
)

// InitDB initializes app database connections
func InitDB() {
	var err error
	DB, err = pop.Connect(config.Config.Database.Configuration)
	if err != nil {
		panic(err)
	}
}
