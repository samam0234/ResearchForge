package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/researchforge/backend/internal/model"
	"github.com/researchforge/backend/internal/service"
)

type Handler struct {
	researchSvc service.ResearchService
}

func NewHandler(researchSvc service.ResearchService) *Handler {
	return &Handler{
		researchSvc: researchSvc,
	}
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// 헬스체크
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "researchforge-backend",
		})
	})

	// Auth Mock (JWT)
	api.Post("/auth/login", func(c *fiber.Ctx) error {
		var body struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
		}
		return c.JSON(fiber.Map{
			"token": "mock-jwt-token-for-scaffolding",
			"user": fiber.Map{
				"id":    "default-user-id",
				"email": body.Email,
				"name":  "Research User",
			},
		})
	})

	// Research 엔드포인트
	api.Post("/research", h.CreateResearch)
	api.Get("/research", h.ListResearches)
	api.Get("/research/:id", h.GetResearchDetail)
	api.Post("/research/:id/export", h.ExportReport)

	// WebSocket 라우트
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/research/:id", websocket.New(h.HandleWebSocket))
}

func (h *Handler) CreateResearch(c *fiber.Ctx) error {
	var req model.CreateResearchRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "잘못된 요청 본문입니다."})
	}
	if req.Query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "연구 질문이나 주제는 필수입니다."})
	}

	userID := "default-user-id" // MVP 임시 사용자
	project, err := h.researchSvc.CreateResearch(c.Context(), userID, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(project)
}

func (h *Handler) GetResearchDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "프로젝트 ID가 필요합니다."})
	}

	detail, err := h.researchSvc.GetResearchDetail(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "리서치 프로젝트를 찾을 수 없습니다."})
	}

	return c.JSON(detail)
}

func (h *Handler) ListResearches(c *fiber.Ctx) error {
	userID := "default-user-id"
	projects, err := h.researchSvc.ListUserResearches(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"items": projects})
}

func (h *Handler) ExportReport(c *fiber.Ctx) error {
	id := c.Params("id")
	detail, err := h.researchSvc.GetResearchDetail(c.Context(), id)
	if err != nil || detail.Report == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "내보낼 완성된 보고서가 없습니다."})
	}

	format := c.Query("format", "markdown")
	if format == "markdown" {
		c.Set("Content-Type", "text/markdown; charset=utf-8")
		c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"report-%s.md\"", id))
		return c.SendString(detail.Report.ContentMarkdown)
	}

	return c.JSON(detail.Report)
}

func (h *Handler) HandleWebSocket(c *websocket.Conn) {
	projectID := c.Params("id")
	eventCh, cleanup := h.researchSvc.SubscribeUpdates(projectID)
	defer cleanup()

	// 초기 연결 확인 메시지 전송
	initEvt := model.WebSocketEvent{
		Type:      "CONNECTED",
		ProjectID: projectID,
		Message:   "실시간 모니터링 세션이 연결되었습니다.",
	}
	initBytes, _ := json.Marshal(initEvt)
	if err := c.WriteMessage(websocket.TextMessage, initBytes); err != nil {
		return
	}

	// 클라이언트 수신 루프 (고루틴)
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			if _, _, err := c.ReadMessage(); err != nil {
				return
			}
		}
	}()

	for {
		select {
		case <-done:
			return
		case evt, ok := <-eventCh:
			if !ok {
				return
			}
			data, err := json.Marshal(evt)
			if err != nil {
				continue
			}
			if err := c.WriteMessage(websocket.TextMessage, data); err != nil {
				return
			}
			if evt.Type == "PROJECT_DONE" || evt.Type == "ERROR" {
				return
			}
		}
	}
}
