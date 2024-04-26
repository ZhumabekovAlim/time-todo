package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders, makeResponseJSON)

	dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()
	//router := httprouter.New()

	// Client
	mux.Post("/create-client", dynamicMiddleware.ThenFunc(app.signupClient))
	mux.Post("/login", dynamicMiddleware.ThenFunc(app.loginClient))
	mux.Post("/verification", dynamicMiddleware.ThenFunc(app.verifyClient))

	// Convoy
	mux.Post("/create-convoy", dynamicMiddleware.ThenFunc(app.createConvoy))    //work
	mux.Get("/get-convoy", standardMiddleware.ThenFunc(app.getConvoy))          //work http://localhost:4000/get-convoy?id=5
	mux.Put("/update-convoy", dynamicMiddleware.ThenFunc(app.updateConvoy))     //work
	mux.Del("/delete-convoy/:id", dynamicMiddleware.ThenFunc(app.deleteConvoy)) //work http://localhost:4000/delete-convoy/124

	// Machine
	mux.Post("/create-machine", dynamicMiddleware.ThenFunc(app.createMachine))    //work
	mux.Get("/get-machine", standardMiddleware.ThenFunc(app.getMachine))          //work http://localhost:4000/get-machine?id=4
	mux.Put("/update-machine", dynamicMiddleware.ThenFunc(app.updateMachine))     //work
	mux.Del("/delete-machine/:id", dynamicMiddleware.ThenFunc(app.deleteMachine)) //work http://localhost:4000/delete-machine/34

	// MhKm
	mux.Post("/create-mhkm", dynamicMiddleware.ThenFunc(app.createMHKM))    //work
	mux.Get("/get-mhkm", standardMiddleware.ThenFunc(app.getMHKM))          //work http://localhost:4000/get-mhkm?id=353
	mux.Put("/update-mhkm", dynamicMiddleware.ThenFunc(app.updateMHKM))     //work
	mux.Del("/delete-mhkm/:id", dynamicMiddleware.ThenFunc(app.deleteMHKM)) //work http://localhost:4000/delete-mhkm/355

	// Service
	mux.Post("/create-service", dynamicMiddleware.ThenFunc(app.createService))
	mux.Get("/get-service", standardMiddleware.ThenFunc(app.getService))
	mux.Put("/update-service", dynamicMiddleware.ThenFunc(app.updateService))
	mux.Del("/delete-service/:id", dynamicMiddleware.ThenFunc(app.deleteService))

	// Repair
	mux.Post("/create-repair", dynamicMiddleware.ThenFunc(app.createRepair))
	mux.Get("/get-repair", standardMiddleware.ThenFunc(app.getRepair))
	mux.Put("/update-repair", dynamicMiddleware.ThenFunc(app.updateRepair))
	mux.Del("/delete-repair/:id", dynamicMiddleware.ThenFunc(app.deleteRepair))

	// ServiceDone
	mux.Post("/create-service-done", dynamicMiddleware.ThenFunc(app.createServiceDone))
	mux.Get("/get-service-done/:id", standardMiddleware.ThenFunc(app.getServiceDone))
	mux.Put("/update-service-done", dynamicMiddleware.ThenFunc(app.updateServiceDone))
	mux.Del("/delete-service-done/:id", dynamicMiddleware.ThenFunc(app.deleteServiceDone))

	// RepairDone
	mux.Post("/create-repair-done", dynamicMiddleware.ThenFunc(app.createRepairDone))
	mux.Get("/get-repair-done/:id", standardMiddleware.ThenFunc(app.getRepairDone))
	mux.Put("/update-repair-done", dynamicMiddleware.ThenFunc(app.updateRepairDone))
	mux.Del("/delete-repair-done/:id", dynamicMiddleware.ThenFunc(app.deleteRepairDone))

	// Default values

	mux.Get("/get-machine-info", standardMiddleware.ThenFunc(app.getMachineInfo)) //work http://localhost:4000/get-machine-info?id_convoy=1
	mux.Get("/get-convoy-info", standardMiddleware.ThenFunc(app.getConvoyInfo))   //work http://localhost:4000/get-convoy-info?id_client=14
	mux.Get("/get-number-by-status", standardMiddleware.ThenFunc(app.getNumberMachineByStasus))
	mux.Get("/get-one-machine-info", standardMiddleware.ThenFunc(app.getOneMachineInfo))

	// Info photo
	mux.Post("/create-info-photo", dynamicMiddleware.ThenFunc(app.createInfoPhoto))
	mux.Get("/get-info-photo", standardMiddleware.ThenFunc(app.getInfoPhoto))          // work http://localhost:4000/get-info-photo?id=60
	mux.Del("/delete-info-photo/:id", dynamicMiddleware.ThenFunc(app.deleteInfoPhoto)) //work http://localhost:4000/delete-info-photo/60

	// Marka
	mux.Get("/get-all-markas", standardMiddleware.ThenFunc(app.getAllMarkas))

	// Types
	mux.Get("/get-all-types", standardMiddleware.ThenFunc(app.getAllTypes))

	// Model
	mux.Get("/get-all-models", standardMiddleware.ThenFunc(app.getAllModels))

	// Balance
	mux.Post("/create-balance/:id", dynamicMiddleware.ThenFunc(app.createBalance)) // work http://localhost:4000/create-balance/1
	mux.Del("/delete-balance/:id", dynamicMiddleware.ThenFunc(app.deleteBalance))  // work http://localhost:4000/delete-balance/1
	mux.Put("/update-balance/:id", dynamicMiddleware.ThenFunc(app.updateBalance))  // work http://localhost:4000/delete-balance/1

	return standardMiddleware.Then(mux)
}
