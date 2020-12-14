package controller

import "github.com/kataras/iris/v12/mvc"

type ExampleController struct {
}

func (c *ExampleController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome</h1>",
	}
}

// GetPing 服务
// 请求方法:   GET
// 请求资源路径: http://localhost:8080/ping
func (c *ExampleController) GetPing() string {
	return "pong"
}

func (c *ExampleController) GetHello() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}

// CustomHandlerWithoutFollowingTheNamingGuide 服务
// 请求方法:   GET
// 请求资源路径: http://localhost:8080/custom_path
func (c *ExampleController) CustomHandlerWithoutFollowingTheNamingGuide() string {
	return "hello from the custom handler without following the naming guide"
}

// GetUserBy 服务
// 请求方法:   GET
// 请求资源路径: http://localhost:8080/user/{username:string}
//是一个保留的关键字来告诉框架你要在函数的输入参数中绑定路径参数，
//在同一控制器中使用“Get”和“GetBy”可以实现
//

func (c *ExampleController) GetUserBy(username string) mvc.Result {
	return mvc.View{
		Name: "user/username.html",
		Data: username,
	}
}
