package mysql

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"graphql-go-template/config"
	"graphql-go-template/util"
	"log"

	"github.com/jinzhu/gorm"

	// Register using gorm mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error

	cfg := config.GetConfig()

	connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQL.User,
		cfg.MySQL.Pass,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.DBName)
	db, err = gorm.Open(
		"mysql",
		connectionString,
	)

	if err != nil {
		log.Fatal(errors.Wrap(err, "Failed to connect"))
	}

	log.Println("Connected mysql db")

	db.DB().SetMaxIdleConns(100)

	if cfg.Debug {
		db = db.Debug()
	}
}

// nolint
func GetClient(ctx context.Context) *gorm.DB {
	cloneDB := &gorm.DB{}
	*cloneDB = *db

	// use transaction per request
	if util.IsEnableTx(ctx) {
		tx := util.GetTx(ctx)
		return tx
	}

	return cloneDB
}

func Disconnect() {
	if db != nil {
		db.Close()
	}
}
