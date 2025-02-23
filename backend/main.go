package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prudhvir3ddy/hiring-score/handlers"
	"github.com/prudhvir3ddy/hiring-score/services"
)

func main() {
	candidateService := services.NewCandidateService()
	if err := candidateService.LoadCandidates(); err != nil {
		log.Fatal(err)
	}

	candidateHandler := handlers.NewCandidateHandler(candidateService)

	r := gin.Default()

	r.GET("/api/candidates", candidateHandler.GetCandidates)

	r.Run(":8080")
}
