package db

const (
	SQLGetDataPendudukByNoKTP = `SELECT no_ktp, "name", birth_place, birth_date, gender, address, religion, marital_status, job, nationality 
								FROM kyc_mst_ktp 
								WHERE no_ktp = $1`
	SQLGetListDataPenduduk = `SELECT no_ktp, "name", birth_place, birth_date, gender, address, religion, marital_status, job, nationality
								FROM kyc_mst_ktp`
	SQLInsertDataPenduduk = `INSERT INTO kyc_mst_ktp (no_ktp, "name", birth_place, birth_date, gender, address, religion, marital_status, job, nationality ) 
							VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	SQLUpdateDataPenduduk = `	UPDATE kyc_mst_ktp SET "name" = $1, birth_place = $2, birth_date = $3, gender = $4, address = $5, religion = $6, marital_status = $7, job = $8, nationality = $9
							WHERE no_ktp = $10`
	SQLDeleteDataPenduduk = `DELETE FROM kyc_mst_ktp
							WHERE no_ktp = $1`
)
