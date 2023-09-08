package usecase

import (
	"github.com/wataruhigasi/Katurao-Hackathon-Back/article/infra"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/domain/model"
)

type ArticleUsecase interface {
	GetArticles() ([]*model.Article, error)
}

type articleUsecaseImpl struct {
	ar infra.ArticleRepository
}

func NewArticleUsecase(ar infra.ArticleRepository) *articleUsecaseImpl {
	return &articleUsecaseImpl{
		ar: ar,
	}
}

func (au *articleUsecaseImpl) GetArticles() ([]*model.Article, error) {
	as, err := au.ar.FindArticles()
	if err != nil {
		return nil, err
	}
	return as, nil
}

func CreateArticle() error {
	return nil
}
