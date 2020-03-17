package internal

import "github.com/ridwanakf/kyc-data-service/internal/entity"

// KTPUC contains logic for all Book Handlers
//go:generate mockgen -destination=usecase/ktp_mock.go -package=usecase github.com/ridwanakf/kyc-data-service/internal KTPUC
type KTPUC interface {
	GetDataPendudukByNoKTP(noKTP string) (entity.KTP, error)
	GetListDataPenduduk() ([]entity.KTP, error)
	InsertDataPenduduk(ktp entity.KTP) (string, error)
	UpdateDataPenduduk(ktp entity.KTP) (string, error)
	DeleteDataPenduduk(noKTP string) (string, error)
}
