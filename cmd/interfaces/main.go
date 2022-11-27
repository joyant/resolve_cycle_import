package main

import (
    "fmt"
    "github.com/joyant/resolve_cycle_import/interfaces/controller"
)

func main() {
    c := controller.NewController()
    c.Like(1, 1)
    detail := c.Detail(1)
    fmt.Printf("%+v", detail)
}
