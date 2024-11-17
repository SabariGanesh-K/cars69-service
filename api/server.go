package api

import (
	"fmt"
"time"
	db "github.com/SabariGanesh-K/cars69-service.git/db/sqlc"
	"github.com/SabariGanesh-K/cars69-service.git/token"
	"github.com/SabariGanesh-K/cars69-service.git/util"
	"github.com/SabariGanesh-K/cars69-service.git/cronjob"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	// "github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token: %v", err)
	}
	server := &Server{
		config: config,
		store: store,
		 tokenMaker: tokenMaker}
	server.setupRouter()
	return server, nil

	
}
func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
			AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
			AllowHeaders:  []string{"Content-Type", "Authorization", "Accept","Access-Control-Allow-Methods"},
			ExposeHeaders: []string{"Content-Length"},
			MaxAge:        12 * time.Hour,
			AllowCredentials: true,
    }))
	// router.Use(cors.Default())

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	// router.OPTIONS("/users/login", server.loginUser)

	router.POST("/tokens/renew_access",server.renewAccessToken)



	// router.POST("/cars/delete",server.updateCar)



	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	// authRoutes.POST("/accounts", server.createAccount)
	authRoutes.POST("/cars/ofuser",server.getUserOwnedCars)
	authRoutes.POST("/cars",server.createCar)
	
	authRoutes.POST("/cars/update",server.updateCar)
	authRoutes.POST("/cars/search",server.searchCarsFTS)
	authRoutes.POST("/cars/delete",server.deleteCar)
	// authRoutes.GET("/accounts/:id", server.getAccount)

	// authRoutes.POST("/transfers", server.createTransfer)

	server.router = router

}

func (server *Server) Start(address string) error {
	c := cronjob.InitCronScheduler()
	defer c.Stop()
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
