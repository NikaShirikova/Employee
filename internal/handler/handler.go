package handler

import (
	"employee/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	log      *zap.Logger
	services *service.Service
}

func NewHandler(services *service.Service, log *zap.Logger) *Handler {
	return &Handler{
		services: services,
		log:      log.Named("handler")}
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
