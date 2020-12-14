package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"irislearn/web/controller"

	//"github.com/kataras/iris/v12/_examples/mvc/overview/web/controllers"
	"github.com/kataras/iris/v12/mvc"
	_ "irislearn/web/controller"
)

func main() {
	fmt.Println("hello iris")

	//
	app := iris.New()

	// 模板
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./web/views", ".html"))

	app.Get("/hello2", func(c iris.Context) {
		fmt.Println("hello")
	})

	// 注册控制器
	//mvc.New(app.Party("/mvc")).Handle(new (controllers.MovieController))
	//mvc.New(app.Party("/mvc")).Handle(new (ExampleController))
	mvc.New(app).Handle(new(controller.ExampleController))

	_ = app.Run(iris.Addr("localhost:801"))
}

func helloCtl() {
	fmt.Println("helloCtl")
}
