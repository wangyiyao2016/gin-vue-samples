package router

import (
	"github.com/gin-gonic/gin"
	"github.com/itachizhu/gin-vue-samples/chapter10-02/controller"
)

func Route(engine *gin.Engine) {
	engine.GET("/json", controller.IndexJsonAction)
	engine.GET("/xml", controller.IndexXmlAction)
	engine.GET("/yaml", controller.IndexYamlAction)
	engine.GET("/image", controller.IndexImageAction)
	engine.GET("/file", controller.IndexFileAction)

	engine.GET("/login", controller.LoginGetAction)
	engine.POST("/login", controller.LoginPostAction)
	engine.PUT("/login", controller.LoginPutAction)

	engine.GET("/products", controller.ProductList)
	engine.GET("/product", controller.GetProduct)
	engine.GET("/delete-product", controller.DeleteProduct)
	engine.POST("/products", controller.SaveProduct)
}
