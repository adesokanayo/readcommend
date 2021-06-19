package books

type Book struct {
	Id            int     `json:"id" db:"id"`
	Title         string  `json:"title" db:"title"`
	YearPublished string  `json:"yearPublished" db:"year_published"`
	Rating        float32 `json:"rating" db:"rating"`
	Pages         int32   `json:"pages" db:"pages"`
	Author        Author  `json:"author" db:"author"`
	Genre         Genre   `json:"genre" db:"genre"`
}

type Genre struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
}

type Author struct {
	Id        int    `json:"id" db:"id"`
	Firstname string `json:"firstName" db:"first_name"`
	Lastname  string `json:"lastName" db:"last_name"`
}

type Size struct {
	Id       int `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	MinPages *int `json:"minpages" db:"min_pages"`
	MaxPages *int `json:"maxpages" db:"max_pages"`
}

type Era struct {
	Id      int    `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	MinYear *int   `json:"minyear" db:"min_year"`
	MaxYear *int   `json:"maxyear" db:"max_year" `
}
