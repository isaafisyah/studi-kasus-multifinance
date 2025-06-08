package app

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/studi-kasus-multifinance/app/config"
	"github.com/isaafisyah/studi-kasus-multifinance/app/database"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/middleware"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/routes"
)

type Server struct {
	Gin *gin.Engine
}

func NewServer() *Server {
	return &Server{
		Gin: gin.Default(),
	}
}

func (s *Server) Run()  {
	cnf := config.Get()
	db, err := database.GetDatabaseConnection(cnf)
	if err != nil {
		panic(err)
	}
	
	s.Gin.Use(middleware.TimeoutMiddleware(5 * time.Second))
	
	//route konsumen
	routes.InitializeRoutesKonsumen(db, s.Gin)
	//route limit
	routes.InitializeRoutesLimit(db, s.Gin)
	//route record transaction
	routes.InitializeRoutesRecordTransaction(db, s.Gin)

	// Start the server
	if err := s.Gin.Run(":" + cnf.Server.Port); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}