package provision

import (
	"errors"
	"log"
	"reflect"
	"time"

	"github.com/ArchiveLife/ArchiveWeibo/api"
	"github.com/ArchiveLife/model/adapter"
	"github.com/ArchiveLife/model/model"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

const KEY_WEIBO_ARTICLE_TYPE = "Weibo"
const KEY_WEIBO_USER_TYPE = "WeiboUser"
const KEY_WEIBO_RESOURCE_TYPE = "WeiboResource"

func createSingleUserWeiboService() adapter.ArchiveService {
	uidDesc := "the 'uid' of weibo user"
	uidLabel := "Weibo User ID"
	return adapter.NewServiceWrapper(
		"weibo user",
		"get all weibo of single user",
		&SingleUserWeiboReader{},
		&adapter.Option{
			Order:       0,
			Name:        "Uid",
			Label:       &uidLabel,
			Description: &uidDesc,
			Optional:    false, // mandatory
			ValueType:   reflect.String,
		},
	)
}

type SingleUserWeiboReader struct {
	Uid         string
	currentPage int
	tmp         []*model.Article
	api         *api.WeiboAPI
	convertor   *md.Converter
}

func (r *SingleUserWeiboReader) Init() error {
	r.api = api.NewWeiboAPI()
	r.convertor = md.NewConverter("", true, nil)
	r.currentPage = 0
	if len(r.Uid) == 0 {
		return errors.New("must provide uid")
	}
	return nil
}

func (r *SingleUserWeiboReader) Next() (*model.Article, bool) {
	if len(r.tmp) > 0 {
		rt := r.tmp[0]
		r.tmp = r.tmp[1:]
		return rt, true
	}
	r.currentPage++
	page, err := r.api.GetUserPagesIndex(r.Uid, r.currentPage)
	if err != nil {
		log.Print(err)
	}
	if page.Ok == 1 {
		if len(page.Data.Cards) > 0 {
			cards := r.convertPageToArticles(page.Data.Cards)
			if len(cards) > 0 {
				rt := cards[0]
				r.tmp = cards[1:]
				return rt, true
			}
		}
	}

	return nil, false
}

func (r *SingleUserWeiboReader) convertPageToArticles(cards []api.Card) (rt []*model.Article) {

	for _, card := range cards {
		if card.Mblog != nil {
			mblog := card.Mblog
			article := &model.Article{
				ID:     model.CreateID(KEY_WEIBO_ARTICLE_TYPE, mblog.ID),
				Medias: []*model.Media{},
			}
			if mblog.CreatedAt != nil {
				if createAt, err := time.Parse(time.RubyDate, *mblog.CreatedAt); err != nil {
					article.PublishDate = &createAt
				}
			}
			if mblog.Text != nil {
				md, err := r.convertor.ConvertString(*mblog.Text)
				if err != nil {
					log.Println("convert md failed", err)
					article.Content = mblog.Text
				} else {
					article.Content = &md
				}
			}
			if mblog.User != nil {
				user := mblog.User
				article.Author = &model.Author{
					ID:       model.CreateID(KEY_WEIBO_USER_TYPE, user.ID),
					FullName: user.ScreenName,
				}
			}
			if len(mblog.Pics) > 0 {
				imageType := "image/jpg"
				for _, pic := range mblog.Pics {
					article.Medias = append(article.Medias, &model.Media{
						ID:           model.CreateID(KEY_WEIBO_RESOURCE_TYPE, pic.URL),
						MimeType:     &imageType,
						ExternalLink: &pic.URL,
					})
				}
			}
			rt = append(rt, article)
		}
	}
	return rt
}
