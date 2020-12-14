package repositories

import (
	"irislearn/services"
)

type TestController struct {
	// Our MovieService, it's an interface which
	// is binded from the main application.
	Service services.MovieService
}
