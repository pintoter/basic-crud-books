package model

type Book struct {
	ID              int     `json:"id"`
	Author          string  `json:"author"`
	Title           string  `json:"name"`
	Rating          float64 `json:"rating"`
	PublicationYear int64   `json:"pubyear"`
}

// непонятно почему тут айти, в репозитори мы используем модель, хотя нам неизвестно ID