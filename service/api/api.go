
 package api

 import (
	 "net/http"
 )
 
 func GetAuthors(w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	 w.WriteHeader(http.StatusOK)
 }
 
 func GetBooks(w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	 w.WriteHeader(http.StatusOK)
 }
 
 func GetEras(w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	 w.WriteHeader(http.StatusOK)
 }
 
 func GetGenres(w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	 w.WriteHeader(http.StatusOK)
 }
 
 func GetSizes(w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	 w.WriteHeader(http.StatusOK)
 }
 