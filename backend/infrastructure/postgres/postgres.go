package postgres

import (
	"context"
	"fmt"
	"github.com/umerm-work/hatch-assignment/config"
	"github.com/umerm-work/hatch-assignment/data"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	dbInstance *gorm.DB
}

func New(ctx context.Context, config config.Config) *DB {
	// dsn := "unix://user:pass@dbname/var/run/postgresql/.s.PGSQL.5432"
	arg := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DB.Host,
		config.DB.Username,
		config.DB.Password,
		config.DB.Database,
		config.DB.Port,
	)
	db, err := gorm.Open(postgres.Open(arg), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	dbConn, err := db.DB()
	if err != nil {
		panic(err)
		return nil
	}
	dbConn.SetMaxOpenConns(50)
	//Migrate the schema
	err = db.AutoMigrate(&data.Todo{})
	if err != nil {
		panic(err)
		return nil
	}
	return &DB{
		dbInstance: db,
	}
}
