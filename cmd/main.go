package main

import (
	"log"

	"github.com/ridwanakf/kyc-data-service/internal/app"
	"github.com/ridwanakf/kyc-data-service/internal/delivery/rest"
)

func main() {
	// init app
	kycApp, err := app.NewKycApp()
	if err != nil {
		log.Fatalf("marshal error %+v", err)
	}
	defer func() {
		errs := kycApp.Close()
		for e := range errs {
			log.Println(e)
		}
	}()

	rest.Start(kycApp)
}
