package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

func main() {
	fmt.Println("hello iris")

	app := iris.New()

	app.Get("/hello2", func(c iris.Context) {
		fmt.Println("hello")
	})

	app.Run(iris.Addr("801"))
}

func helloCtl() {
	fmt.Println("helloCtl")
}
