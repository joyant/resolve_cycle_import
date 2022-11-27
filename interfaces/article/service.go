package article

func NewService() Service {
    return &service{}
}

func (s *service) Exists(id int64) bool {
    return true
}

func (s *service) GetDetail(id int64) *Detail {
    return &Detail{
        ID:         id,
        Title:      "my title",
        Content:    "my content",
        LikeAmount: s.like.GetLikeAmount(id),
    }
}

func (s *service) SetLikeService(like like) {
    s.like = like
}
