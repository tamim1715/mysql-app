package router

import (
	v1 "github.com/tamim1715/mysql-app/api/v1"

	"github.com/labstack/echo/v4"
)

func V1Routes(e *echo.Group) {
	e.GET("/:id", v1.CacheController().Get)
	e.POST("", v1.CacheController().Create)
	e.PUT("", v1.CacheController().Update)
	e.DELETE("/:id", v1.CacheController().Delete)
}
