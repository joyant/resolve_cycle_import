package article

type service struct {
    like like
}

type Detail struct {
    ID         int64
    Title      string
    Content    string
    LikeAmount int64
}

type like interface {
    GetLikeAmount(id int64) int64
}

type Service interface {
    Exists(id int64) bool
    GetDetail(id int64) *Detail
    SetLikeService(like like)
}
