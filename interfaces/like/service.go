package like

import "fmt"

type service struct {
    article article
}

func NewService() Service {
    return &service{}
}

func (s *service) GetLikeAmount(id int64) int64 {
    return 99
}

func (s *service) Like(id int64, status byte) {
    if status == 1 {
        fmt.Println("like article", id)
    } else {
        fmt.Println("cancel like article", id)
    }
}

func (s *service) SetArticleService(article article) {
    s.article = article
}
