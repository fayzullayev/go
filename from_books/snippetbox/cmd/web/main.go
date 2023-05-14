package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/fayzullayev/snippetbox/internal/models"
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type config struct {
	addr string
	dsn  string
}

type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	snippets       *models.SnippetModel
	templateCache  TemplateCacheType
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
	user           *models.UserModel
}

type TemplateCacheType = map[string]*template.Template

func main() {
	var cfg config

	flag.StringVar(&cfg.addr, "addr", ":4001", "HTTP network address")
	flag.StringVar(&cfg.dsn, "dsn", "web:123456789@/snippetbox?parseTime=true", "HTTP network address")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(cfg.dsn)

	if err != nil {
		errorLog.Fatal(err)
		return
	}

	defer db.Close()

	templateCache, err := newTemplateCache()

	if err != nil {
		errorLog.Fatal(err)
		return
	}

	formDecoder := form.NewDecoder()

	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	sessionManager.Cookie.Secure = true

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{
			DB: db,
		},
		user: &models.UserModel{
			DB: db,
		},
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
	}

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         cfg.addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", cfg.addr)

	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")

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
