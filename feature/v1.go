package feature

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/storage"
	authStorage "github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/storage"
	authTransport "github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/transport"
	authUsecase "github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/usecase"
	userStorage "github.com/cesc1802/onboarding-and-volunteer-service/feature/user/storage"
	userTransport "github.com/cesc1802/onboarding-and-volunteer-service/feature/user/transport"
	userUsecase "github.com/cesc1802/onboarding-and-volunteer-service/feature/user/usecase"
	"net/http"

	"github.com/cesc1802/share-module/system"
	"github.com/gin-gonic/gin"
)

func RegisterHandlerV1(mono system.Service) {
	router := mono.Router()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{
			"data": "success",
		})
	})
	v1 := router.Group("/api/v1")
	// Initialize database connection
	storage.DBContext()

	// Initialize repository
	authRepo := authStorage.NewAuthenticationRepository(storage.DB)
	userRepo := userStorage.NewAdminRepository(storage.DB)

	// Initialize usecase
	authUseCase := authUsecase.NewUserUsecase(authRepo)
	userUseCase := userUsecase.NewAdminUsecase(userRepo)

	// Initialize handler
	authHandler := authTransport.NewAuthenticationHandler(authUseCase)
	userHandler := userTransport.NewAuthenticationHandler(userUseCase)

	auth := v1.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)

		auth.POST("/register", func(c *gin.Context) {

		})
	}

	admin := v1.Group("/admin")
	{
		admin.GET("/list-request", userHandler.GetListRequest)
		admin.GET("/request/:id", userHandler.GetRequestById)
		admin.POST("/approve-request/:id", userHandler.ApproveRequest)
		admin.POST("/reject-request/:id", userHandler.RejectRequest)
		admin.POST("/add-reject-notes/:id", userHandler.AddRejectNotes)
		admin.DELETE("/delete-request/:id", userHandler.DeleteRequest)
	}

}
