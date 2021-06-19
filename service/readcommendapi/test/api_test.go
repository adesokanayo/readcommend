package test

import (
	_ "fmt"
	"io/ioutil"
	"log"
	"net/url"
	"testing"

	"github.com/adesokanayo/readcommend/internal/pkg/config"
	"github.com/adesokanayo/readcommend/internal/pkg/db"
	"github.com/adesokanayo/readcommend/internal/pkg/persistence"
	_ "github.com/adesokanayo/readcommend/internal/pkg/persistence"
	"github.com/stretchr/testify/assert"
)

func Setup() {
	config.Setup("./")
	db.SetupDB()

	//Run startup script
	c, ioErr := ioutil.ReadFile("migrate.sql")
	if ioErr != nil {
		
		log.Println("Unable to locate migration file",ioErr)
	}
	sql := string(c)
	_, err := db.GetDB().Exec(sql)
	if err != nil {
		//Enable this if teardown is Implemented
		//log.Println("Unable to execute migration file",err)
	}

}

func TestListBooks(t *testing.T) {
	Setup()
	u, err := url.Parse("http://localhost:5000/api/v1/books")
	
	engine:= persistence.NewQueryEngine(*u)
    query := persistence.QueryBuilder(*engine)
	assert.Equal(t,query,`SELECT b.id, b.title as book_title , b.year_published as year_published, b.rating as book_rating , b.pages as book_pages FROM book b  order by b.rating DESC `)
	s, err := persistence.ListBooks(query)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(*s), 58)
}


func TestListAuthors(t *testing.T) {
	Setup()
	s, err := persistence.ListAuthors()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(*s),41)
}

func TestListEras(t *testing.T) {
	Setup()
	s, err := persistence.ListEras()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(*s), 3)

}

func TestListGenres(t *testing.T) {
	Setup()

	s, err := persistence.ListGenres()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(*s), 8)
}

