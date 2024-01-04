package postgre

import "database/sql"

func New(url string) (
	s, err := sql.Open("postgres", url)
)