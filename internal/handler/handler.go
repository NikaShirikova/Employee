package handler

import (
	"Employee/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/add", h.AddEmpl)
	router.DELETE("/delete/:id", h.DeleteEmpl)
	router.PUT("/update", h.UpdateEmpl)
	router.GET("/company/:nameCompany", h.GetCompany)
	router.GET("/department/:nameDepartment", h.GetDepartment)

	return router
}
