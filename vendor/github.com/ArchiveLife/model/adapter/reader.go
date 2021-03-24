package adapter

import (
	"github.com/ArchiveLife/model/model"
)

type ArticleReader interface {
	// call after options inject
	Init() error
	Next() (*model.Article, bool)
}
