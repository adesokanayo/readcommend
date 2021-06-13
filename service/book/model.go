package model

type Book struct {
	Id            int    `json:"Id" db:"Id"`
	Title         string `json:"title" db:"title"`
	YearPublished string `json:"yearPublished" db:"year_published"`
	Rating        int8   `json:"rating" db:"rating"`
	Pages         int32  `json:"pages" db:"pages"`
	GENREID       string `json:"genreId" db:"genre_id"`
	Author        Author `json:"author" db:"author"`
	Genre         Genre  `json:"genre" db:"genre"`
}

type Genre struct {
	Id    int    `json:"Id"`
	Title string `json:"title" db:"title"`
}

type Author struct {
	Id        int    `json:"Id"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname  string `json:"lastname" db:"lastname"`
}
