package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ridwanakf/kyc-data-service/internal"
	"github.com/ridwanakf/kyc-data-service/internal/app"
	"github.com/ridwanakf/kyc-data-service/internal/entity"
	"github.com/ridwanakf/kyc-data-service/internal/helper"
)

type KTPService struct {
	uc internal.KTPUC
}

func NewKTPService(app *app.KycApp) *KTPService {
	return &KTPService{
		uc: app.UseCases.KTP,
	}
}

func (k *KTPService) GetDataPendudukByNoKTP(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	noKTP := param.ByName("noKTP")

	res, err := k.uc.GetDataPendudukByNoKTP(noKTP)
	if err != nil {
		log.Printf("[service][GetDataPendudukByNoKTP] fail to get data :%+v\n", err)
		helper.RenderJSON(w, []byte(`
		{
			"status": "failed",
			"message": "ktp not found!"
		}
		`), http.StatusOK)
		return
	}

	bytes, err := json.Marshal(res)
	if err != nil {
		log.Printf("[service][GetDataPendudukByNoKTP] fail to marshal data :%+v\n", err)
		return
	}

	helper.RenderJSON(w, bytes, http.StatusOK)
}

func (k *KTPService) GetListDataPenduduk(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	res, err := k.uc.GetListDataPenduduk()
	if err != nil {
		log.Printf("[service][GetListDataPenduduk] fail to get data :%+v\n", err)
		return
	}

	bytes, err := json.Marshal(res)
	if err != nil {
		log.Printf("[service][GetListDataPenduduk] fail to marshal data :%+v\n", err)
		return
	}

	helper.RenderJSON(w, bytes, http.StatusOK)
}

func (k *KTPService) InsertDataPenduduk(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		helper.RenderJSON(w, []byte(`
			"message": "Failed to read body"
		`), http.StatusBadRequest)
		return
	}

	var ktp entity.KTP
	err = json.Unmarshal(body, &ktp)
	if err != nil {
		log.Printf("[service][InsertDataPenduduk] error unmarshal :%+v\n", err)
		return
	}

	res, _ := k.uc.InsertDataPenduduk(ktp)

	helper.RenderJSON(w, []byte("{\"message\": "+"\""+res+"\""+"}"), http.StatusOK)
}

func (k *KTPService) UpdateDataPenduduk(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		helper.RenderJSON(w, []byte(`
			"message": "Failed to read body"
		`), http.StatusBadRequest)
		return
	}

	var ktp entity.KTP
	err = json.Unmarshal(body, &ktp)
	if err != nil {
		log.Printf("[service][UpdateDataPenduduk] error unmarshal :%+v\n", err)
		return
	}

	res, _ := k.uc.UpdateDataPenduduk(ktp)

	helper.RenderJSON(w, []byte("{\"message\": "+"\""+res+"\""+"}"), http.StatusOK)
}

func (k *KTPService) DeleteDataPenduduk(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	noKTP := param.ByName("noKTP")

	res, _ := k.uc.DeleteDataPenduduk(noKTP)

	helper.RenderJSON(w, []byte("{\"message\": "+"\""+res+"\""+"}"), http.StatusOK)
}
