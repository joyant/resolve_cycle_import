package controller

import "github.com/joyant/resolve_cycle_import/interfaces/article"

type Controller interface {
    Like(id int64, status byte)
    Detail(id int64) *article.Detail
}
