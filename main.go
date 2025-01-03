package main

import (
	"github.com/gin-gonic/gin"
	"go_mico_demo/internal/handler"
	"go_mico_demo/internal/repository"
	"go_mico_demo/internal/service"
	"go_mico_demo/pkg/config"
	"go_mico_demo/pkg/database"
	"log"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化数据库连接
	db, err := database.NewMySQLConnection(
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.Username,
		cfg.MySQL.Password,
		cfg.MySQL.Database,
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 初始化依赖
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// 设置路由
	router := gin.Default()

	// 用户相关路由
	userGroup := router.Group("/api/users")
	{
		userGroup.POST("", userHandler.CreateUser)
		userGroup.GET("/:id", userHandler.GetUser)
	}

	// 启动服务器
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
