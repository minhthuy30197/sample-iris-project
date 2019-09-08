package router

import (
	"github.com/kataras/iris"
	"sample-project/controller"

)

func BlogRoutes(c *controller.Controller, api *iris.Application) {
	api.Get("/blog", c.GetBlog)
	api.Get("/blog/{id}", c.GetPostById)
}