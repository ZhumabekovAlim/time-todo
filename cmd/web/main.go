package main

import (
	"database/sql"
	"flag"
	"github.com/golangcollege/sessions"
	"github.com/rs/cors"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	"time-todo/pkg/models/dbs"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	templateCache map[string]*template.Template
	client        *dbs.ClientModel
	convoy        *dbs.ConvoyModel
	machine       *dbs.MachineModel
	mhkm          *dbs.MhKmModel
	service       *dbs.ServiceModel
	repair        *dbs.RepairModel
	serviceDone   *dbs.ServiceDoneModel
	repairDone    *dbs.RepairDoneModel
	machineInfo   *dbs.MachineInfoModel
	convoyInfo    *dbs.ConvoyInfoModel
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/v-1831_technic"
	addr := flag.String("addr", ":4000", "HTTP network address")

	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")

	flag.Parse()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowCredentials: true,
	})

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	app := &application{
		errorLog:    errorLog,
		infoLog:     infoLog,
		session:     session,
		client:      &dbs.ClientModel{DB: db},
		convoy:      &dbs.ConvoyModel{DB: db},
		machine:     &dbs.MachineModel{DB: db},
		mhkm:        &dbs.MhKmModel{DB: db},
		service:     &dbs.ServiceModel{DB: db},
		repair:      &dbs.RepairModel{DB: db},
		serviceDone: &dbs.ServiceDoneModel{DB: db},
		repairDone:  &dbs.RepairDoneModel{DB: db},
		machineInfo: &dbs.MachineInfoModel{DB: db},
		convoyInfo:  &dbs.ConvoyInfoModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  c.Handler(app.routes()),
		// TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", *addr)

	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
