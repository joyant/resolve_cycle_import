package controller

import (
    "github.com/joyant/resolve_cycle_import/interfaces/article"
    "github.com/joyant/resolve_cycle_import/interfaces/like"
)

type controller struct {
    articleService article.Service
    likeService    like.Service
}

func NewController() Controller {
    likeService := like.NewService()
    articleService := article.NewService()
    likeService.SetArticleService(articleService)
    articleService.SetLikeService(likeService)

    return &controller{
        articleService: articleService,
        likeService:    likeService,
    }
}

func (c *controller) Like(id int64, status byte) {
    c.likeService.Like(id, status)
}

func (c *controller) Detail(id int64) *article.Detail {
    return c.articleService.GetDetail(id)
}
