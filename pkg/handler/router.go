package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetRouter() *echo.Echo {
	router := echo.New()

	router.POST("/api/data/create", h.InsertCars)

	router.GET("/api/data", h.GetCars)

	router.PUT("/api/data/:id", h.UpdateInfo)

	router.PUT("/api/owner/:id", h.UpdateOwner)

	router.DELETE("/api/data/:id", h.DeleteCarByID)

	return router
}
