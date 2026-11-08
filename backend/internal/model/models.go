package model

import (
	"time"
)

type ProjectStatus string

const (
	StatusPending   ProjectStatus = "pending"
	StatusRunning   ProjectStatus = "running"
	StatusCompleted ProjectStatus = "completed"
	StatusFailed    ProjectStatus = "failed"
)

type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ResearchOptions struct {
	Depth       string   `json:"depth"`        // shallow, normal, deep
	MinSources  int      `json:"min_sources"`  // 최소 출처 수 (기본: 5)
	Language    string   `json:"language"`     // ko, en
	ReportFormat string  `json:"report_format"`// markdown, academic, summary
}

type ResearchProject struct {
	ID          string          `json:"id"`
	UserID      string          `json:"user_id"`
	Title       string          `json:"title"`
	Query       string          `json:"query"`
	Options     ResearchOptions `json:"options"`
	Status      ProjectStatus   `json:"status"`
	CreatedAt   time.Time       `json:"created_at"`
	CompletedAt *time.Time      `json:"completed_at,omitempty"`
}

type AgentExecution struct {
	ID         string                 `json:"id"`
	ProjectID  string                 `json:"project_id"`
	AgentName  string                 `json:"agent_name"` // Researcher, FactChecker, Analyst, Writer
	Role       string                 `json:"role"`
	Input      map[string]interface{} `json:"input"`
	Output     map[string]interface{} `json:"output"`
	Status     string                 `json:"status"` // running, completed, failed
	StartedAt  time.Time              `json:"started_at"`
	FinishedAt *time.Time             `json:"finished_at,omitempty"`
}

type SourceItem struct {
	Title       string  `json:"title"`
	URL         string  `json:"url"`
	Snippet     string  `json:"snippet"`
	Credibility float64 `json:"credibility"` // 신뢰도 점수 (0.0 ~ 1.0)
}

type Report struct {
	ID              string                 `json:"id"`
	ProjectID       string                 `json:"project_id"`
	ContentMarkdown string                 `json:"content_markdown"`
	Sources         []SourceItem           `json:"sources"`
	Metadata        map[string]interface{} `json:"metadata"`
	CreatedAt       time.Time              `json:"created_at"`
}

// DTOs
type CreateResearchRequest struct {
	Title   string          `json:"title"`
	Query   string          `json:"query"`
	Options ResearchOptions `json:"options"`
}

type ResearchDetailResponse struct {
	Project    ResearchProject  `json:"project"`
	Executions []AgentExecution `json:"executions"`
	Report     *Report          `json:"report,omitempty"`
}

type WebSocketEvent struct {
	Type      string      `json:"type"`       // STEP_START, STEP_PROGRESS, STEP_COMPLETE, PROJECT_DONE, ERROR
	ProjectID string      `json:"project_id"`
	AgentName string      `json:"agent_name,omitempty"`
	Progress  int         `json:"progress,omitempty"` // 0 ~ 100
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}
