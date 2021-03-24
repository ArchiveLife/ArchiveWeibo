package adapter

import (
	"github.com/ArchiveLife/core/model"
)

type ArticleReader interface {
	// call after options inject
	Init() error
	Next() (*model.Article, bool)
}
