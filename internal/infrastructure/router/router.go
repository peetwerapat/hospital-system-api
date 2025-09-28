package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/peetwerapat/hospital-system-api/docs"
	"github.com/peetwerapat/hospital-system-api/internal/infrastructure/di"
	"github.com/peetwerapat/hospital-system-api/internal/interface/controller"
	"github.com/peetwerapat/hospital-system-api/pkg/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(appUc *di.AppUseCase) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// CORS config
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PACTH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// Staff
	staffController := controller.NewStaffController(appUc.StaffUC)
	r.POST("/staff/create", staffController.CreateStaff)
	r.POST("/staff/login", staffController.StaffLogin)

	// Patient
	patientController := controller.NewPatientController(appUc.PatientUC)
	r.GET("/patient/search", middleware.AuthMiddleware(), patientController.GetPatientsByHospitalID)

	return r
}
