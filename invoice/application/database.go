package application

import (
	"fmt"

	"github.com/syauqeesy/accounting-service/invoice/configuration"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type databaseApplication struct {
	configuration         *configuration.Configuration
	database              *gorm.DB
	databaseConfiguration *gorm.Config
	databaseConnectionUrl string
}

func (a *databaseApplication) Init() error {
	a.databaseConnectionUrl = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", a.configuration.Database.User, a.configuration.Database.Password, a.configuration.Database.Host, a.configuration.Database.Port, a.configuration.Database.Name)

	a.databaseConfiguration = &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	return nil
}

func (a *databaseApplication) Run() (err error) {
	a.database, err = gorm.Open(mysql.Open(a.databaseConnectionUrl), a.databaseConfiguration)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	fmt.Println("connected to database")

	return nil
}

func (a *databaseApplication) Close() error {
	sql, err := a.database.DB()
	if err != nil {
		return fmt.Errorf("failed to get database object: %w", err)
	}

	err = sql.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}

	fmt.Println("database connection closed.")

	return nil
}
