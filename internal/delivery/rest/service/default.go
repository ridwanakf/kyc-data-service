package service

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ridwanakf/kyc-data-service/internal/helper"
)

type DefaultService struct {
}

func NewDefaultService() *DefaultService {
	return &DefaultService{
	}
}

// Index Page
func (h *DefaultService) Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	helper.RenderJSON(w, []byte(`
		{
			"module": "kyc-data-service",
			"version": "1.0.0-clean-arch", 
			"date":"17 March 2020",
			"desc": "This is index page"
		}
	`), http.StatusOK)

}

// ServeHTTP is used for 404 page
func (h *DefaultService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	helper.RenderJSON(w, []byte(`
		{
			"message": "There's nothing here"
		}
	`), http.StatusNotFound)
}
