package model

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

//List returns all known products
func List(db *sqlx.DB) ([]Book, error) {

	var list []Book

	const q = `SELECT product_id, name, cost, quantity, date_updated, date_created FROM products`
	if err := db.Select(&list, q); err != nil {
		log.Println("error querying database", err)
	}
	return list, nil
}

//Retrieves a single product
func Retrieve(db *sqlx.DB, id int) (*Book, error) {

	var b Book

	const q = `SELECT Id, title, yearPublished, rating , pages 
	FROM books WHERE Id = $1`
	if err := db.Get(&p, q); err != nil {
		log.Println("error querying database", err)
	}
	return &b, nil
}

func Create(db *sqlx.DB, np Book, now time.Time) (*Book, error) {

	b := Book{
		Id:            1,
		Title:         np.Title,
		YearPublished: "2014",
		Rating:        5,
		Pages:         100,
	}

	const q = `INSERT INTO Books
	(Id, title, yearpublished, rating, pages)
	VALUES($1,$2,$3,$4,$5)`

	if _, err := db.Exec(q, b.Id, b.Title, b.YearPublished, b.Rating, b.Pages); err != nil {
		return nil, errors.Wrap(err, "selecting single book")
	}
	return &b, nil
}
