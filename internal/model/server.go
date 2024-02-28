// Package model provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package model

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/gin-gonic/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// list of foodposts
	// (GET /foodpost)
	ListFoodPosts(c *gin.Context, params ListFoodPostsParams)
	// create food post
	// (POST /foodpost)
	CreateFoodPost(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// ListFoodPosts operation middleware
func (siw *ServerInterfaceWrapper) ListFoodPosts(c *gin.Context) {

	var err error

	c.Set(BasicAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListFoodPostsParams

	// ------------- Optional query parameter "restaurantName" -------------

	err = runtime.BindQueryParameter("form", true, false, "restaurantName", c.Request.URL.Query(), &params.RestaurantName)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter restaurantName: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ListFoodPosts(c, params)
}

// CreateFoodPost operation middleware
func (siw *ServerInterfaceWrapper) CreateFoodPost(c *gin.Context) {

	c.Set(BasicAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.CreateFoodPost(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {

	errorHandler := options.ErrorHandler

	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/foodpost", wrapper.ListFoodPosts)

	router.POST(options.BaseURL+"/foodpost", wrapper.CreateFoodPost)

	return router
}
