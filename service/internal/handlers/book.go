package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adesokanayo/readcommend/service/book"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Book struct {
	DB  *sqlx.DB
	Log *log.Logger
}

func (b *Book) List(w http.ResponseWriter, r *http.Request) error {

	list, err := book.List(b.DB)
	if err != nil {
		return err
	}

	data, err := json.Marshal(list)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return errors.Wrap(err, "Marshalling value of json")
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		return errors.Wrap(err, "Writing Response")
	}

	return nil

}

