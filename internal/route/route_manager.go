package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "empty/docs" // replace with actual module path
	"empty/internal/delivery/http"
	"empty/internal/repository"
	"empty/internal/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(r *gin.Engine, db *pgxpool.Pool) {
	userRoute := r.Group("/api/v1/users")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	SetupUserRoutes(userRoute, db)
}

func SetupUserRoutes(r *gin.RouterGroup, db *pgxpool.Pool) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := http.NewUserHandler(userService)
	r.POST("/", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUserByID)
}
