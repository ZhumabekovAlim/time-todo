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

	// Client
	mux.Post("/create-client", dynamicMiddleware.ThenFunc(app.signupClient))
	mux.Post("/login", dynamicMiddleware.ThenFunc(app.loginClient))

	// Convoy
	mux.Post("/create-convoy", dynamicMiddleware.ThenFunc(app.createConvoy))
	mux.Get("/get-convoy", standardMiddleware.ThenFunc(app.getConvoy))
	mux.Put("/update-convoy", dynamicMiddleware.ThenFunc(app.updateConvoy))
	mux.Del("/delete-convoy/:id", dynamicMiddleware.ThenFunc(app.deleteConvoy))

	// Machine
	mux.Post("/create-machine", dynamicMiddleware.ThenFunc(app.createMachine))
	mux.Get("/get-machine", standardMiddleware.ThenFunc(app.getMachine))
	mux.Put("/update-machine", dynamicMiddleware.ThenFunc(app.updateMachine))
	mux.Del("/delete-machine/:id", dynamicMiddleware.ThenFunc(app.deleteMachine))

	// MhKm
	mux.Post("/create-mhkm", dynamicMiddleware.ThenFunc(app.createMHKM))
	mux.Get("/get-mhkm", standardMiddleware.ThenFunc(app.getMHKM))
	mux.Put("/update-mhkm", dynamicMiddleware.ThenFunc(app.updateMHKM))
	mux.Del("/delete-mhkm/:id", dynamicMiddleware.ThenFunc(app.deleteMHKM))

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

	return standardMiddleware.Then(mux)
}
