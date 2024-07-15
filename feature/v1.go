package feature

import (
	"net/http"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/storage"
	authStorage "github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/storage"
	authTransport "github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/transport"
	authUsecase "github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/usecase"
	userStorage "github.com/cesc1802/onboarding-and-volunteer-service/feature/user/storage"
	userTransport "github.com/cesc1802/onboarding-and-volunteer-service/feature/user/transport"
	userUsecase "github.com/cesc1802/onboarding-and-volunteer-service/feature/user/usecase"

	appliIdentityStorage "github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/storage"
	appliIdentityTransport "github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/transport"
	appliIdentityUsecase "github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/usecase"

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
	applicantRepo := userStorage.NewApplicantRepository(storage.DB)
	applicantRequestRepo := userStorage.NewApplicantRequestRepository(storage.DB)
	applicantIdentityRepo := appliIdentityStorage.NewUserIdentityRepository(storage.DB)

	// Initialize usecase
	authUseCase := authUsecase.NewUserUsecase(authRepo)
	userUseCase := userUsecase.NewAdminUsecase(userRepo)
	applicantUseCase := userUsecase.NewApplicantUsecase(applicantRepo)
	applicantRequestUseCase := userUsecase.NewApplicantRequestUsecase(applicantRequestRepo)
	applicantIdenityUseCase := appliIdentityUsecase.NewUserIdentityUsecase(applicantIdentityRepo)

	// Initialize handler
	authHandler := authTransport.NewAuthenticationHandler(authUseCase)
	userHandler := userTransport.NewAuthenticationHandler(userUseCase)
	applicantHandler := userTransport.NewApplicantHandler(applicantUseCase)
	applicantRequestHandler := userTransport.NewRequestHandler(applicantRequestUseCase)
	applicantIdentityHandler := appliIdentityTransport.NewUserIdentityHandler(applicantIdenityUseCase)

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

	applicant := v1.Group("/applicant")
	{
		applicant.POST("/", applicantHandler.CreateApplicant)
		applicant.PUT("/update/:id", applicantHandler.UpdateApplicant)
		applicant.DELETE("/:id", applicantHandler.CreateApplicant)
		applicant.GET("/:id")
	}

	appliRequest := v1.Group("/applicant-request")
	{
		appliRequest.POST("/", applicantRequestHandler.CreateRequest)
	}

	appliIdentity := v1.Group("applicant-identity")
	{
		appliIdentity.POST("/", applicantIdentityHandler.CreateUserIdentity)
		appliIdentity.GET("/:id", applicantIdentityHandler.FindUserIdentity)
		appliIdentity.POST("/:id", applicantIdentityHandler.UpdateUserIdentity)
	}

}
