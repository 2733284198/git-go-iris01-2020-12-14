package controller

import (
	"github.com/kataras/iris/v12/mvc"
)

type MovieController struct {
}

func (c *MovieController) Get() mvc.View {

	return mvc.View{
		Name: "movie/index.html",
		Data: "movie_controller.go",
	}
}
