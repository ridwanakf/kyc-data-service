package service

import "github.com/ridwanakf/kyc-data-service/internal/app"

type Services struct {
	*KTPService
	*DefaultService
}

func GetServices(app *app.KycApp) *Services {
	return &Services{
		KTPService:     NewKTPService(app),
		DefaultService: NewDefaultService(),
	}
}
