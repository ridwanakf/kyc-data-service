package db

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"
	"github.com/ridwanakf/kyc-data-service/internal/entity"
)

type KTPDB struct {
	db *sql.DB
}

func NewKTPDB(db *sql.DB) *KTPDB {
	return &KTPDB{
		db: db,
	}
}

func (k *KTPDB) GetDataPendudukByNoKTP(noKTP string) (entity.KTP, error) {
	stmt, err := k.db.Prepare(SQLGetDataPendudukByNoKTP)
	if err != nil {
		return entity.KTP{}, errors.Errorf("[repo][GetDataPendudukByNoKTP] invalid prepare statement :%+v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(noKTP)
	var ktps []entity.KTP
	for rows.Next() {
		ktp := entity.KTP{}
		err := rows.Scan(
			&ktp.NoKTP,
			&ktp.Name,
			&ktp.BirthPlace,
			&ktp.BirthDate,
			&ktp.Gender,
			&ktp.Address,
			&ktp.Religion,
			&ktp.MaritalStatus,
			&ktp.Job,
			&ktp.Nationality,
		)
		if err != nil {
			log.Printf("[repo][GetDataPendudukByNoKTP] fail to scan :%+v\n", err)
			continue
		}
		ktps = append(ktps, ktp)
	}
	if len(ktps) < 1 {
		return entity.KTP{}, errors.New("[repo][GetDataPendudukByNoKTP] data not found!")
	}
	return ktps[0], nil
}

func (k *KTPDB) GetListDataPenduduk() ([]entity.KTP, error) {
	stmt, err := k.db.Prepare(SQLGetListDataPenduduk)
	if err != nil {
		return []entity.KTP{}, errors.Errorf("[repo][GetListDataPenduduk] invalid prepare statement :%+v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	var ktps []entity.KTP
	for rows.Next() {
		ktp := entity.KTP{}
		err := rows.Scan(
			&ktp.NoKTP,
			&ktp.Name,
			&ktp.BirthPlace,
			&ktp.BirthDate,
			&ktp.Gender,
			&ktp.Address,
			&ktp.Religion,
			&ktp.MaritalStatus,
			&ktp.Job,
			&ktp.Nationality,
		)
		if err != nil {
			log.Printf("[repo][GetListDataPenduduk] fail to scan :%+v\n", err)
			continue
		}
		ktps = append(ktps, ktp)
	}

	return ktps, nil
}

func (k *KTPDB) InsertDataPenduduk(ktp entity.KTP) (string, error) {
	stmt, err := k.db.Prepare(SQLInsertDataPenduduk)
	if err != nil {
		return "failed", errors.Errorf("[repo][InsertDataPenduduk] invalid prepare statement :%+v\n", err)
	}
	defer stmt.Close()

	_, err = stmt.Query(
		ktp.NoKTP,
		ktp.Name,
		ktp.BirthPlace,
		ktp.BirthDate,
		ktp.Gender,
		ktp.Address,
		ktp.Religion,
		ktp.MaritalStatus,
		ktp.Job,
		ktp.Nationality,
	)
	if err != nil {
		return "failed", errors.Errorf("[repo][InsertDataPenduduk] error inserting new data :%+v\n", err)
	}

	return "success", nil
}

func (k *KTPDB) UpdateDataPenduduk(ktp entity.KTP) (string, error) {
	stmt, err := k.db.Prepare(SQLUpdateDataPenduduk)
	if err != nil {
		return "failed", errors.Errorf("[repo][UpdateDataPenduduk] invalid prepare statement :%+v\n", err)
	}
	defer stmt.Close()

	_, err = stmt.Query(
		ktp.Name,
		ktp.BirthPlace,
		ktp.BirthDate,
		ktp.Gender,
		ktp.Address,
		ktp.Religion,
		ktp.MaritalStatus,
		ktp.Job,
		ktp.Nationality,
		ktp.NoKTP, )
	if err != nil {
		return "failed", errors.Errorf("[repo][UpdateDataPenduduk] error updating data :%+v\n", err)
	}

	return "success", nil
}

func (k *KTPDB) DeleteDataPenduduk(noKTP string) (string, error) {
	stmt, err := k.db.Prepare(SQLDeleteDataPenduduk)
	if err != nil {
		return "failed", errors.Errorf("[repo][DeleteDataPenduduk] invalid prepare statement :%+v\n", err)
	}
	defer stmt.Close()

	_, err = stmt.Query(noKTP)
	if err != nil {
		return "failed", errors.Errorf("[repo][DeleteDataPenduduk] error deleting data :%+v\n", err)
	}

	return "success", nil
}
