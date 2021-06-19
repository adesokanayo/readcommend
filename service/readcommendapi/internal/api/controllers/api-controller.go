package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/adesokanayo/readcommend/internal/pkg/persistence"
	error "github.com/adesokanayo/readcommend/pkg/error"
	"github.com/gin-gonic/gin"
)

// GetBooks retrieves all books, can also filter based on query strings
func GetBooks(c *gin.Context) {

	queryeng := persistence.NewQueryEngine(*c.Request.URL)

	queryeng =  queryeng.ProcessAuthor().ProcessGenre().
	ProcessPages().ProcessYear().ProcessLimit()
	                
	query := persistence.QueryBuilder(*queryeng)

	q, err := persistence.ListBooks(query)
	_ = c.Bind(&q)
	if err != nil {
		error.NewError(c, http.StatusNotFound, errors.New("books not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, q)
	}
}

func GetAuthors(c *gin.Context) {

	q, err := persistence.ListAuthors()
	_ = c.Bind(&q)
	if err != nil {
		error.NewError(c, http.StatusNotFound, errors.New("genres not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, q)
	}
}
func GetGenres(c *gin.Context) {

	q, err := persistence.ListGenres()
	_ = c.Bind(&q)
	if err != nil {
		error.NewError(c, http.StatusNotFound, errors.New("genres not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, q)
	}
}

func GetSizes(c *gin.Context) {

	q, err := persistence.ListSizes()
	_ = c.Bind(&q)
	if err != nil {
		error.NewError(c, http.StatusNotFound, errors.New("sizes not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, q)
	}
}

func GetEras(c *gin.Context) {

	q, err := persistence.ListEras()
	_ = c.Bind(&q)
	if err != nil {
		error.NewError(c, http.StatusNotFound, errors.New("eras not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, q)
	}
}
