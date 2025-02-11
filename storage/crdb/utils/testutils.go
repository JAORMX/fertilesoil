//go:build testtools
// +build testtools

package utils

import (
	"database/sql"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"testing"

	"github.com/cockroachdb/cockroach-go/v2/testserver"
	"github.com/pressly/goose/v3"

	"github.com/infratographer/fertilesoil/storage/crdb/migrations"
)

type StopServerFunc func()

var once sync.Once

func NewTestDBServer() (*url.URL, StopServerFunc, error) {
	srv, err := testserver.NewTestServer()
	if err != nil {
		return nil, nil, err
	}

	if err := srv.WaitForInit(); err != nil {
		return nil, nil, err
	}

	dbURL := srv.PGURL()

	// Reset Path so we can use the database in general
	dbURL.Path = "/"

	return dbURL, func() {
		srv.Stop()
	}, nil
}

func NewTestDBServerOrDie() (*url.URL, StopServerFunc) {
	dbURL, cleanup, err := NewTestDBServer()
	if err != nil {
		panic(fmt.Sprintf("error creating test database server: %v", err))
	}
	return dbURL, cleanup
}

func GetNewTestDB(t *testing.T, baseDBURL *url.URL) *sql.DB {
	t.Helper()

	dbName := strings.ToLower(strings.ReplaceAll(t.Name(), "/", "_"))

	baseDB, err := sql.Open("postgres", baseDBURL.String())
	if err != nil {
		t.Fatalf("error opening database: %v", err)
	}

	if _, err := baseDB.Exec("CREATE DATABASE " + dbName); err != nil {
		t.Fatalf("error creating database: %v", err)
	}

	dbURL := baseDBURL.JoinPath(dbName)
	dbConn, err := sql.Open("postgres", dbURL.String())
	if err != nil {
		t.Fatalf("error opening database: %v", err)
	}

	once.Do(func() {
		goose.SetBaseFS(migrations.Migrations)

		if err := goose.SetDialect("postgres"); err != nil {
			t.Fatalf("error setting dialect: %v", err)
		}
	})

	if err := goose.Up(dbConn, "."); err != nil {
		t.Fatalf("error running migrations: %v", err)
	}
	return dbConn
}
