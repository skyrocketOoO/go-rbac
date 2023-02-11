package http

import (
	RepoDomain "go-authorization/domain/repository"
	UsecaseDomain "go-authorization/domain/usecase"

	"github.com/gin-gonic/gin"
)

type ResourceHandler struct {
	ResourceUsecase UsecaseDomain.ResourceUsecase
}

func NewResourceHandler(server *gin.Engine, usecase UsecaseDomain.ResourceUsecase) {
	handler := &ResourceHandler{
		ResourceUsecase: usecase,
	}

	user := server.Group("/resources")
	user.GET("/", handler.GetAll)
	user.GET("/:name", handler.Get)
	user.POST("/", handler.Create)
	user.PUT("/:name", handler.Update)
	user.DELETE("/:name", handler.Delete)
}

func (rsh *ResourceHandler) GetAll(c *gin.Context) {
	res, err := rsh.ResourceUsecase.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, res)
}

func (rsh *ResourceHandler) Get(c *gin.Context) {
	resource, err := rsh.ResourceUsecase.Get(c.Request.Context(), c.Param("name"))
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, resource)
}

func (rsh *ResourceHandler) Create(c *gin.Context) {
	resource := RepoDomain.Resource{}
	if err := c.BindJSON(&resource); err != nil {
		c.JSON(400, err.Error())
	}

	if err := rsh.ResourceUsecase.Create(c.Request.Context(), &resource); err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "success")
}

func (rsh *ResourceHandler) Update(c *gin.Context) {
	resource := RepoDomain.Resource{}
	if err := c.BindJSON(&resource); err != nil {
		c.JSON(400, err.Error())
	}

	if err := rsh.ResourceUsecase.Update(c.Request.Context(), c.Param("name"), &resource); err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "success")
}

func (rsh *ResourceHandler) Delete(c *gin.Context) {
	if err := rsh.ResourceUsecase.Delete(c.Request.Context(), c.Param("name")); err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "success")
}
