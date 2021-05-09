package main

import (
	"testing"
	//"github.com/DATA-DOG/go-sqlmock"
)

func Test_createDatabaseString(t *testing.T) {
	params := databaseParams{"james", "password", "1234", "test"}

	dbString := createDatabaseString(params, params.dbName)

	if dbString != "james:password@tcp(1234)/test" {
		t.Errorf("Expected the output string to be james:password@tcp(1234)/test, instead it was %s", dbString)
	}
}
