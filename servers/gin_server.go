package servers

import (
	"fmt"

	"Github.com/Yobubble/go-clean-boilerplate/configs"
	"Github.com/Yobubble/go-clean-boilerplate/pkg/features/auth/handlers"
	"Github.com/Yobubble/go-clean-boilerplate/pkg/features/auth/repositories"
	"Github.com/Yobubble/go-clean-boilerplate/pkg/features/auth/usecases"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ginServer struct {
	cfg configs.Config
	app *gin.Engine
	db  *gorm.DB
}

// InitAuthHttpHandler implements Server.
func (g *ginServer) InitAuthHttpHandler() {
	authRepository := repositories.NewAuthPostgresqlRepository(g.db)
	authUseCase := usecases.NewAuthUseCase(authRepository)
	authHttpHandler := handlers.NewAuthHttpHandler(authUseCase)

	v1 := g.app.Group("/v1/auth")
	v1.POST("/sign-in", authHttpHandler.SignIn)
	v1.POST("/sign-up", authHttpHandler.SignUp)
}

// Start implements Server.
func (g *ginServer) Start() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = g.cfg.Client.AllowedClient
	g.app.Use(cors.New(corsConfig))

	g.InitAuthHttpHandler()
	port := fmt.Sprintf(":%d", g.cfg.App.Port)
	g.app.Run(port)
}

func NewGinServer(cfg configs.Config, db *gorm.DB) Server {
	return &ginServer{
		cfg: cfg,
		app: gin.Default(),
		db:  db,
	}
}
