package persistence

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/adesokanayo/readcommend/internal/pkg/db"
	"github.com/adesokanayo/readcommend/internal/pkg/models/books"
)

// QueryEngine holds the information about the query
type QueryEngine struct {
	query           string
	values          map[string]bool
	URL             url.URL
	Condition       bool
	ConditionNumber int
	Limit           string
}

// NewQueryEngine creates a new query engine
func NewQueryEngine(u url.URL) *QueryEngine {
	var queryengine QueryEngine
	queryengine.values = make(map[string]bool)
	queryengine.URL = u
	return &queryengine
}

func ListBooks(q string) (*[]books.Book, error) {

	data := []struct {
		BookId          int     `db:"id"`
		BookTitle       string  ` db:"book_title"`
		YearPublished   string  ` db:"year_published"`
		BookRating      float32 ` db:"book_rating"`
		BookPages       int32   `db:"book_pages"`
		GenreId         int     `db:"genre_id"`
		GenreTitle      string  ` db:"genre_title"`
		AuthorId        int     `db:"author_id"`
		AuthorFirstname string  `db:"author_first_name"`
		AuthorLastname  string  `db:"author_last_name"`
	}{}

	if err := db.DB.Select(&data, q); err != nil {
		log.Println("error querying database", err)
		return nil,err
	}

	mystructs := make([]books.Book, len(data))

	for i, b := range data {
		mystructs[i] = books.Book{
			Id:     b.BookId,
			Pages:  b.BookPages,
			Rating: b.BookRating,
			Title:  b.BookTitle,
			Author: books.Author{
				Id:        b.AuthorId,
				Firstname: b.AuthorFirstname,
				Lastname:  b.AuthorLastname,
			},
			Genre: books.Genre{
				Id:    b.GenreId,
				Title: b.GenreTitle,
			},
		}
	}

	return &mystructs, nil
}

func ListAuthors() (*[]books.Author, error) {

	var list []books.Author

	q := `SELECT id  ,first_name ,last_name FROM author`
	if err := db.DB.Select(&list, q); err != nil {
		log.Println("error querying database", err)
	}
	return &list, nil
}

// ListGenres fetches all genres 
func ListGenres() (*[]books.Genre, error) {

	var list []books.Genre

	q := `SELECT id ,title  FROM genre`
	if err := db.DB.Select(&list, q); err != nil {
		log.Println("error querying database", err)
	}
	return &list, nil
}

// ListSizes fetches all sizes 
func ListSizes() (*[]books.Size, error) {

	var list []books.Size

	q := `SELECT id  ,title , min_pages, max_pages from  size`
	if err := db.DB.Select(&list, q); err != nil {
		log.Println("error querying database", err)
	}
	return &list, nil
}

// ListEras fetches all eras 
func ListEras() (*[]books.Era, error) {

	var list []books.Era

	q := `SELECT id ,title , min_year, max_year FROM era`
	if err := db.DB.Select(&list, q); err != nil {
		log.Println("error querying database", err)
	}
	return &list, nil
}

// QueryBuilder builds query bases on query strings
func QueryBuilder(engine QueryEngine) string {

	query := ""
	basic := "SELECT b.id, b.title as book_title , b.year_published as year_published, b.rating as book_rating , b.pages as book_pages "
	author_query := ",a.id as author_id , a.first_name AS author_first_name, a.last_name AS author_last_name  "
	genre_query := ",g.id as genre_id ,g.title AS genre_title "
	from_query := "FROM book b "
	join_author_query := "JOIN author  AS  a ON b.author_id = a.id "
	join_genre_query := "JOIN genre AS g ON b.genre_id  = g.id "
	order_query := " order by b.rating DESC "
	query = basic

	for k, _ := range engine.values {

		if k == "authors" {
			query += author_query
			engine.Condition = true
		}
		if k == "genres" {
			query += genre_query
			engine.Condition = true
		}
	}
	query += from_query
	for k, _ := range engine.values {

		if k == "authors" {
			query += join_author_query
		}
		if k == "genres" {

			query += join_genre_query

		}
	}
	if engine.Condition {
		query += fmt.Sprintf("where %s", engine.query)
	}

	query += order_query
	if engine.Limit != "" {
		query += fmt.Sprintf("limit %s", engine.Limit)
	}
	return query
}

// ProcessAuthor builds query  for authors in query string 
func (q QueryEngine) ProcessAuthor() *QueryEngine {
	authors := q.URL.Query().Get("authors")
	if authors != "" {
		q.query += fmt.Sprintf("a.id in (%s)", parseQueryString(authors))
		q.values["authors"] = true
		q.ConditionNumber++
		return &q
	}
	return &q
}

// ProcessGenre builds query  for genre in query string 
func (q QueryEngine) ProcessGenre() *QueryEngine {

	genres := q.URL.Query().Get("genres")
	if genres != "" {
		if q.ConditionNumber >= 1 {
			q.query += fmt.Sprintf(" and g.id in (%s)", parseQueryString(genres))
		} else {
			q.query += fmt.Sprintf(" g.id in (%s)", genres)
		}
		q.ConditionNumber++
		q.values["genres"] = true
		return &q
	}
	return &q
}

// ProcessPages builds query  for pages in query string 
func (q QueryEngine) ProcessPages() *QueryEngine {

	minPage := q.URL.Query().Get("min-pages")
	maxPage := q.URL.Query().Get("max-pages")
	if minPage != "" {
		fmt.Println("MinPages", minPage)
		if q.ConditionNumber >= 1 {
			q.query += fmt.Sprintf(" and b.pages <= %s ", parseQueryString(minPage))
		} else {
			q.query += fmt.Sprintf(" b.pages <= %s", minPage)
		}
		q.ConditionNumber++
		q.values["minPages"] = true
	}
	if maxPage != "" {
		fmt.Println("MaxPages", maxPage)
		if q.ConditionNumber >= 1 {
			q.query += fmt.Sprintf(" and b.pages >= %s ", parseQueryString(maxPage))
		} else {
			q.query += fmt.Sprintf(" b.pages >= %s", maxPage)
		}
		q.ConditionNumber++
		q.values["maxPages"] = true
	}

	return &q
}

// ProcessYear builds query  for years in query string 
func (q QueryEngine) ProcessYear() *QueryEngine {

	minYear := q.URL.Query().Get("min-year")
	maxYear := q.URL.Query().Get("max-year")
	if minYear != "" {
		fmt.Println("MinPages", minYear)
		if q.ConditionNumber >= 1 {
			q.query += fmt.Sprintf(" and b.year_published <= %s ", parseQueryString(minYear))
		} else {
			q.query += fmt.Sprintf(" b.year_published <= %s", minYear)
		}
		q.ConditionNumber++
		q.values["minYear"] = true
	}
	if maxYear != "" {
		fmt.Println("MaxYear", maxYear)
		if q.ConditionNumber >= 1 {
			q.query += fmt.Sprintf(" and b.pages >= %s ", parseQueryString(maxYear))
		} else {
			q.query += fmt.Sprintf(" b.pages >= %s", maxYear)
		}
		q.ConditionNumber++
		q.values["maxYear"] = true
	}

	return &q
}
// ProcessLimit builds query  for limit in query string 
func (q QueryEngine) ProcessLimit() *QueryEngine {

	limit := q.URL.Query().Get("limit")
	if limit != "" {
		fmt.Println("limit", limit)
		q.ConditionNumber++
		q.values["limit"] = true
		q.Limit = limit
		return &q
	}

	return &q
}

// ProcessLimit builds query  for limit in query string 
func parseQueryString(entry string) (output string) {

	if strings.Contains(entry, ",") {
		arr := strings.Split(entry, ",")

		for i, v := range arr {
			arr[i] = fmt.Sprintf("'%s'", v)
		}
		output = strings.Join(arr, ",")
		return output
	}
	return entry
}
