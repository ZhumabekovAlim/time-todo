package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
	"github.com/rs/cors"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	logger        *log.Logger
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
	infoPhoto     *dbs.InfoPhotoModel
	balance       *dbs.BalanceModel
	marka         *dbs.MarkaModel
	types         *dbs.TypeModel
	models        *dbs.ModelModel
	gormDB        *gorm.DB
}

func main() {
	// todo: go get -u gorm.io/gorm
	dsn := "root:rootpassword@tcp(localhost:3307)/technic"
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
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("gormDB failed to connect")
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	app := &application{
		errorLog:    errorLog,
		infoLog:     infoLog,
		logger:      logger,
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
		infoPhoto:   &dbs.InfoPhotoModel{DB: db},
		balance:     &dbs.BalanceModel{DB: db, GormDB: gormDB},
		marka:       &dbs.MarkaModel{DB: db},
		types:       &dbs.TypeModel{DB: db},
		models:      &dbs.ModelModel{DB: db},
		gormDB:      gormDB,
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
