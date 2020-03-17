package app

import (
	"database/sql"

	"github.com/ridwanakf/kyc-data-service/internal"
	db2 "github.com/ridwanakf/kyc-data-service/internal/repo/db"
)

type Repos struct {
	ktpRepo internal.KTPRepo
}

func newRepos(db *sql.DB) (*Repos, error) {
	r := &Repos{
		ktpRepo: db2.NewKTPDB(db),
	}

	return r, nil
}

func (r *Repos) Close() []error {
	var errs []error
	return errs
}
