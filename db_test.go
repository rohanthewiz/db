package db

import (
	"testing"
)

func TestDb(t *testing.T) {
	err := InitDB(DBOpts{
		DBType: DBTypes.Postgres,
		Host: "localhost",
		Port: "5432",
		User: "pg_user", // put db creds in
		Word: "pg_word",
		Database: "test_db",
	})
	if err != nil {
		t.Error(err, "Could not setup database")
	}
	defer CloseDB()

	dbH, err := Db()
	if err != nil {
		t.Error(err, "Could not obtain a valid DB handle")
	}
	err = dbH.Ping()
	if err != nil {
		t.FailNow()
	}
}
