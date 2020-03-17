package rest

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/gops/agent"
	"github.com/julienschmidt/httprouter"
	"github.com/ridwanakf/kyc-data-service/internal/app"
	"github.com/ridwanakf/kyc-data-service/internal/delivery/rest/service"
	"github.com/ridwanakf/kyc-data-service/internal/entity"
)

func initFlags(args *entity.Args) {
	port := os.Getenv("PORT")
	if port == "" {
		port = *(flag.String("PORT", "5000", "port number for your apps"))
	}
	args.Port = port
}

func initRouter(router *httprouter.Router, svc *service.Services) {

	//Default path
	router.GET("/", svc.DefaultService.Index)

	//KYC API paths
	router.GET("/data/:noKTP", svc.KTPService.GetDataPendudukByNoKTP)
	router.GET("/list", svc.KTPService.GetListDataPenduduk)
	router.POST("/data", svc.KTPService.InsertDataPenduduk)
	router.PUT("/data", svc.KTPService.UpdateDataPenduduk)
	router.DELETE("/data/:noKTP", svc.KTPService.DeleteDataPenduduk)

	// `httprouter` library uses `ServeHTTP` method for it's 404 pages
	router.NotFound = svc.DefaultService
}

func Start(app *app.KycApp) {
	args := new(entity.Args)
	initFlags(args)

	svc := service.GetServices(app)
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
	initRouter(router, svc)

	fmt.Println("Apps served on :" + args.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":"+args.Port), router))
}
