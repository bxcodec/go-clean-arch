package repository

import "github.com/bxcodec/go-clean-arch/author"

type AuthorRepository interface {
	GetByID(id int64) (*author.Author, error)
}
