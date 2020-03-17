package usecase

import (
	"github.com/ridwanakf/kyc-data-service/internal"
	"github.com/ridwanakf/kyc-data-service/internal/entity"
)

type KTPUsecase struct {
	repo internal.KTPRepo
}

func NewKTPUsecase(repo internal.KTPRepo) *KTPUsecase {
	return &KTPUsecase{
		repo: repo,
	}
}

func (k *KTPUsecase) GetDataPendudukByNoKTP(noKTP string) (entity.KTP, error) {
	return k.repo.GetDataPendudukByNoKTP(noKTP)
}

func (k *KTPUsecase) GetListDataPenduduk() ([]entity.KTP, error) {
	return k.repo.GetListDataPenduduk()
}

func (k *KTPUsecase) InsertDataPenduduk(ktp entity.KTP) (string, error) {
	return k.repo.InsertDataPenduduk(ktp)
}

func (k *KTPUsecase) UpdateDataPenduduk(ktp entity.KTP) (string, error) {
	return k.repo.UpdateDataPenduduk(ktp)
}

func (k *KTPUsecase) DeleteDataPenduduk(noKTP string) (string, error) {
	return k.repo.DeleteDataPenduduk(noKTP)
}
