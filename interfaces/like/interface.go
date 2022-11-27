package like

type article interface {
    Exists(id int64) bool
}

type Service interface {
    GetLikeAmount(id int64) int64
    Like(id int64, status byte)
    SetArticleService(article article)
}
