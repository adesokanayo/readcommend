package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	api "github.com/adesokanayo/readcommend/api"
	database "github.com/adesokanayo/readcommend/database"
	"github.com/adesokanayo/readcommend/internal/handlers"
	"github.com/pkg/errors"
)

func main() {
	log.Printf("Server started")
  
	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(":8085", router))
}

func start() error {
	//=============================================================//
	log.Printf("main: Started")

	log := log.New(os.Stderr, "ReadCommendAPI :", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	defer log.Println("main :Completed")

	// =================================

	//Setup dependencies

	var cfg struct {
		Web struct {
			Address         string        `conf:"default:localhost:8000"`
			ReadTimeout     time.Duration `conf:"default:5s"`
			WriteTimeout    time.Duration `conf:"default:5s"`
			ShutdownTimeout time.Duration `conf:"default:5s"`
		}
		DB struct {
			User       string `conf:"default:postgres"`
			Password   string `conf:"default:postgres,noprint"`
			Host       string `conf:"default:localhost"`
			Name       string `conf:"default:postgres"`
			DisableTLS bool   `conf:"default:false"`
		}
	}


	//log.Printf("main : Config :\n%v\n", out)
	db, err := database.Open(database.Config{
		Host:       cfg.DB.Host,
		User:       cfg.DB.User,
		Password:   cfg.DB.Password,
		Name:       cfg.DB.Name,
		DisableTLS: cfg.DB.DisableTLS,
	})
	if err != nil {
		return errors.Wrap(err, "unable to open db connection")
	}

	defer db.Close()

	//ps := handlers.Book{DB: db, Log: log}
	// =================================
	api := http.Server{
		Addr:         cfg.Web.Address,
		//Handler:       ps.List(),
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout.Round(2),
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("main : API Listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "listening and serving")
	case <-shutdown:
		log.Println("main: start shutdown")

		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		err := api.Shutdown(ctx)
		if err != nil {
			err = api.Close()
			return errors.Wrap(err, "main:Graceful shutdown did not completed")

		}

		if err != nil {
			return errors.Wrap(err, "main: could not stop server gracefully: %v")
		}

	}
	return nil
}