package app

import (
	"github.com/ridwanakf/kyc-data-service/internal"
	"github.com/ridwanakf/kyc-data-service/internal/usecase"
)

type Usecases struct {
	KTP internal.KTPUC //Book Usecase
}

func newUsecases(repos *Repos) *Usecases {
	return &Usecases{
		KTP: usecase.NewKTPUsecase(repos.ktpRepo),
	}
}

func (*Usecases) Close() []error {
	var errs []error
	return errs
}
