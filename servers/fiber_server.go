package servers

import (
	"fmt"

	"Github.com/Yobubble/go-clean-boilerplate/configs"
	"Github.com/Yobubble/go-clean-boilerplate/pkg/auth/handlers"
	"Github.com/Yobubble/go-clean-boilerplate/pkg/auth/repositories"
	"Github.com/Yobubble/go-clean-boilerplate/pkg/auth/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type fiberServer struct {
	cfg configs.Config
	app *fiber.App
	db  *gorm.DB
}

// InitAuthHttpHandler implements Server.
func (f *fiberServer) InitAuthHttpHandler() {
	authRepository := repositories.NewAuthPostgresqlRepository(f.db)
	authUseCase := usecases.NewAuthUseCase(authRepository)
	authHttpHandler := handlers.NewAuthHttpHandler(authUseCase)

	v1 := f.app.Group("/v1/auth")
	v1.Post("/sign-in", authHttpHandler.SignIn)
}

// Start implements Server.
func (f *fiberServer) Start() {
	port := fmt.Sprintf(":%d", f.cfg.App.Port)
	f.InitAuthHttpHandler()
	// middlewares
	f.app.Listen(port)
}

func NewFiberServer(db *gorm.DB, cfg configs.Config) Server {
	return &fiberServer{
		cfg: cfg,
		app: fiber.New(),
		db:  db,
	}
}
