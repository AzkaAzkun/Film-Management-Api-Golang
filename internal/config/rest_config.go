package config

import (
	"film-management-api-golang/db"
	"film-management-api-golang/internal/api/controller"
	"film-management-api-golang/internal/api/repository"
	"film-management-api-golang/internal/api/routes"
	"film-management-api-golang/internal/api/service"
	"film-management-api-golang/internal/middleware"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type RestConfig struct {
	server *gin.Engine
}

func NewRest() RestConfig {
	db := db.New()
	app := gin.Default()
	server := NewRouter(app)
	middleware := middleware.New(db)

	var (
		//=========== (PACKAGE) ===========//
		// mailerService mailer.Mailer         = mailer.New()

		//=========== (REPOSITORY) ===========//
		userRepository  repository.UserRepository  = repository.NewUser(db)
		genreRepository repository.GenreRepository = repository.NewGenre(db)

		//=========== (SERVICE) ===========//
		authService  service.AuthService  = service.NewAuth(userRepository, db)
		genreService service.GenreService = service.NewGenre(genreRepository, db)

		//=========== (CONTROLLER) ===========//
		authController  controller.AuthController  = controller.NewAuth(authService)
		genreController controller.GenreController = controller.NewGenre(genreService)
	)

	routes.Auth(server, authController, middleware)
	routes.Genre(server, genreController, middleware)
	return RestConfig{
		server: server,
	}
}

func (ap *RestConfig) Start() {
	port := os.Getenv("APP_PORT")
	host := os.Getenv("APP_HOST")
	if port == "" {
		port = "8090"
	}

	serve := fmt.Sprintf("%s:%s", host, port)
	if err := ap.server.Run(serve); err != nil {
		log.Panicf("failed to start server: %s", err)
	}
	log.Println("server start on port ", serve)
}
