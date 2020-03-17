package internal

import "github.com/ridwanakf/kyc-data-service/internal/entity"

// KTPRepo contains logic for getting KTP from repo
//go:generate mockgen -destination=repo/ktp_mock.go -package=repo github.com/ridwanakf/kyc-data-service/internal KTPRepo
type KTPRepo interface {
	GetDataPendudukByNoKTP(noKTP string) (entity.KTP, error)
	GetListDataPenduduk() ([]entity.KTP, error)
	InsertDataPenduduk(ktp entity.KTP) (string, error)
	UpdateDataPenduduk(ktp entity.KTP) (string, error)
	DeleteDataPenduduk(noKTP string) (string, error)
}
