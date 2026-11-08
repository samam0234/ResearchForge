package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/researchforge/backend/internal/client"
	"github.com/researchforge/backend/internal/config"
	"github.com/researchforge/backend/internal/handler"
	"github.com/researchforge/backend/internal/middleware"
	"github.com/researchforge/backend/internal/repository"
	"github.com/researchforge/backend/internal/service"
	"github.com/researchforge/backend/pkg/queue"
)

func main() {
	cfg := config.LoadConfig()

	app := fiber.New(fiber.Config{
		AppName:      "ResearchForge Backend API",
		ServerHeader: "Fiber",
	})

	// 미들웨어 설정
	middleware.SetupMiddlewares(app)

	// 계층별 의존성 주입
	repo := repository.NewInMemoryRepository()
	aiClient := client.NewAIClient(cfg.AIServerURL)
	eventBus := queue.NewMemoryEventBus()
	researchSvc := service.NewResearchService(repo, aiClient, eventBus)
	h := handler.NewHandler(researchSvc)

	h.RegisterRoutes(app)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("🚀 ResearchForge Backend Server starting on %s...", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Server terminated with error: %v", err)
	}
}
