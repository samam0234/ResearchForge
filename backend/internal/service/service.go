package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/researchforge/backend/internal/client"
	"github.com/researchforge/backend/internal/model"
	"github.com/researchforge/backend/internal/repository"
	"github.com/researchforge/backend/pkg/queue"
)

type ResearchService interface {
	CreateResearch(ctx context.Context, userID string, req model.CreateResearchRequest) (*model.ResearchProject, error)
	GetResearchDetail(ctx context.Context, projectID string) (*model.ResearchDetailResponse, error)
	ListUserResearches(ctx context.Context, userID string) ([]model.ResearchProject, error)
	SubscribeUpdates(projectID string) (<-chan model.WebSocketEvent, func())
}

type researchService struct {
	repo     repository.Repository
	aiClient client.AIClient
	eventBus queue.EventBus
}

func NewResearchService(repo repository.Repository, aiClient client.AIClient, eventBus queue.EventBus) ResearchService {
	return &researchService{
		repo:     repo,
		aiClient: aiClient,
		eventBus: eventBus,
	}
}

func (s *researchService) CreateResearch(ctx context.Context, userID string, req model.CreateResearchRequest) (*model.ResearchProject, error) {
	if req.Title == "" {
		req.Title = req.Query
		if len(req.Title) > 50 {
			req.Title = req.Title[:50] + "..."
		}
	}
	if req.Options.MinSources <= 0 {
		req.Options.MinSources = 5
	}
	if req.Options.Depth == "" {
		req.Options.Depth = "normal"
	}
	if req.Options.Language == "" {
		req.Options.Language = "ko"
	}

	project := &model.ResearchProject{
		ID:        uuid.New().String(),
		UserID:    userID,
		Title:     req.Title,
		Query:     req.Query,
		Options:   req.Options,
		Status:    model.StatusPending,
		CreatedAt: time.Now(),
	}

	if err := s.repo.CreateProject(project); err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}

	// 비동기 파이프라인 백그라운드 실행
	go s.executeResearchPipeline(project)

	return project, nil
}

func (s *researchService) executeResearchPipeline(project *model.ResearchProject) {
	bgCtx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	// 1. 상태를 running으로 변경
	_ = s.repo.UpdateProjectStatus(project.ID, model.StatusRunning, nil)
	_ = s.eventBus.PublishEvent(bgCtx, project.ID, model.WebSocketEvent{
		Type:      "PROJECT_STATUS",
		ProjectID: project.ID,
		Progress:  5,
		Message:   "리서치 작업이 시작되었습니다.",
		Timestamp: time.Now(),
	})

	// 단계별 에이전트 이벤트 시뮬레이션 및 전송
	steps := []struct {
		Agent    string
		Role     string
		StartMsg string
		DoneMsg  string
		Progress int
	}{
		{"Researcher", "웹 검색 및 자료 수집", "관련 웹 문서 및 최신 자료를 검색 중입니다...", "자료 수집 완료", 30},
		{"FactChecker", "출처 검증 및 팩트체크", "수집된 출처의 신뢰성 및 사실관계를 교차 검증 중입니다...", "출처 교차 검증 완료", 60},
		{"Analyst", "정보 분석 및 인사이트 도출", "핵심 데이터 분석 및 맥락별 인사이트를 합성 중입니다...", "분석 완료", 80},
		{"Writer", "구조화된 보고서 작성", "최종 마크다운 보고서와 출처 링크를 생성 중입니다...", "보고서 작성 완료", 95},
	}

	for _, step := range steps {
		stepStart := time.Now()
		exec := &model.AgentExecution{
			ID:        uuid.New().String(),
			ProjectID: project.ID,
			AgentName: step.Agent,
			Role:      step.Role,
			Status:    "running",
			StartedAt: stepStart,
		}
		_ = s.repo.CreateAgentExecution(exec)

		_ = s.eventBus.PublishEvent(bgCtx, project.ID, model.WebSocketEvent{
			Type:      "STEP_START",
			ProjectID: project.ID,
			AgentName: step.Agent,
			Progress:  step.Progress - 15,
			Message:   step.StartMsg,
			Timestamp: time.Now(),
		})

		time.Sleep(1 * time.Second) // 스텝 전환 인터벌

		stepDone := time.Now()
		exec.Status = "completed"
		exec.FinishedAt = &stepDone
		_ = s.repo.UpdateAgentExecution(exec)

		_ = s.eventBus.PublishEvent(bgCtx, project.ID, model.WebSocketEvent{
			Type:      "STEP_COMPLETE",
			ProjectID: project.ID,
			AgentName: step.Agent,
			Progress:  step.Progress,
			Message:   step.DoneMsg,
			Timestamp: time.Now(),
		})
	}

	// 2. AI Server 호출하여 최종 보고서 수령
	resp, err := s.aiClient.RunResearchCrew(bgCtx, project)
	now := time.Now()
	if err != nil || (resp != nil && resp.Status == "failed") {
		_ = s.repo.UpdateProjectStatus(project.ID, model.StatusFailed, &now)
		_ = s.eventBus.PublishEvent(bgCtx, project.ID, model.WebSocketEvent{
			Type:      "ERROR",
			ProjectID: project.ID,
			Message:   "리서치 파이프라인 처리 중 오류가 발생했습니다.",
			Timestamp: now,
		})
		return
	}

	// 3. 보고서 저장 및 최종 완료
	if resp != nil && resp.Report != nil {
		_ = s.repo.SaveReport(resp.Report)
	}
	_ = s.repo.UpdateProjectStatus(project.ID, model.StatusCompleted, &now)

	_ = s.eventBus.PublishEvent(bgCtx, project.ID, model.WebSocketEvent{
		Type:      "PROJECT_DONE",
		ProjectID: project.ID,
		Progress:  100,
		Message:   "리서치 보고서가 성공적으로 완성되었습니다.",
		Data:      resp.Report,
		Timestamp: now,
	})
}

func (s *researchService) GetResearchDetail(ctx context.Context, projectID string) (*model.ResearchDetailResponse, error) {
	p, err := s.repo.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	execs, _ := s.repo.ListExecutionsByProjectID(projectID)
	rep, _ := s.repo.GetReportByProjectID(projectID)

	return &model.ResearchDetailResponse{
		Project:    *p,
		Executions: execs,
		Report:     rep,
	}, nil
}

func (s *researchService) ListUserResearches(ctx context.Context, userID string) ([]model.ResearchProject, error) {
	return s.repo.ListProjectsByUserID(userID)
}

func (s *researchService) SubscribeUpdates(projectID string) (<-chan model.WebSocketEvent, func()) {
	return s.eventBus.Subscribe(projectID)
}
