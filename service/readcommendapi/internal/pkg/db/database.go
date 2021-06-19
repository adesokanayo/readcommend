package db

import (
	"fmt"
	"time"

	"github.com/adesokanayo/readcommend/internal/pkg/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	DB  *sqlx.DB
	err error
)

type Database struct {
	*sqlx.DB
}

// SetupDB opens a database and saves the reference to `Database` struct.
func SetupDB() {
	var db = DB

	configuration := config.GetConfig()

	driver := configuration.Database.Driver
	database := configuration.Database.Dbname
	username := configuration.Database.Username
	password := configuration.Database.Password
	host := configuration.Database.Host
	port := configuration.Database.Port

	db, err = sqlx.Open(driver, "host="+host+" port="+port+" user="+username+" dbname="+database+"  sslmode=disable password="+password)
	if err != nil {
		fmt.Println("db err: ", err)
	}

	db.SetMaxIdleConns(configuration.Database.MaxIdleConns)
	db.SetMaxOpenConns(configuration.Database.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(configuration.Database.MaxLifetime) * time.Second)
	DB = db
}

func GetDB() *sqlx.DB {
	return DB
}
