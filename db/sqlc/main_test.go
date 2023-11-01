package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"

	"github.com/Ali-Adel-Nour/Bank/util"

)



var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
config,err := util.LoadConfig("../..")
if err != nil {
	log.Fatal("cannot load config: ", err)
}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	if err = testDB.Ping(); err != nil {
		log.Fatal("cannot ping database", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
