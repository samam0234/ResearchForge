package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/researchforge/backend/internal/model"
)

var (
	ErrNotFound = errors.New("resource not found")
)

type Repository interface {
	CreateProject(project *model.ResearchProject) error
	GetProjectByID(id string) (*model.ResearchProject, error)
	ListProjectsByUserID(userID string) ([]model.ResearchProject, error)
	UpdateProjectStatus(id string, status model.ProjectStatus, completedAt *time.Time) error

	CreateAgentExecution(exec *model.AgentExecution) error
	UpdateAgentExecution(exec *model.AgentExecution) error
	ListExecutionsByProjectID(projectID string) ([]model.AgentExecution, error)

	SaveReport(report *model.Report) error
	GetReportByProjectID(projectID string) (*model.Report, error)
}

// InMemoryRepository provides a robust fallback / development implementation
type InMemoryRepository struct {
	mu         sync.RWMutex
	projects   map[string]*model.ResearchProject
	executions map[string][]*model.AgentExecution
	reports    map[string]*model.Report
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		projects:   make(map[string]*model.ResearchProject),
		executions: make(map[string][]*model.AgentExecution),
		reports:    make(map[string]*model.Report),
	}
}

func (r *InMemoryRepository) CreateProject(project *model.ResearchProject) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if project.ID == "" {
		project.ID = uuid.New().String()
	}
	if project.CreatedAt.IsZero() {
		project.CreatedAt = time.Now()
	}
	r.projects[project.ID] = project
	return nil
}

func (r *InMemoryRepository) GetProjectByID(id string) (*model.ResearchProject, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.projects[id]
	if !ok {
		return nil, ErrNotFound
	}
	return p, nil
}

func (r *InMemoryRepository) ListProjectsByUserID(userID string) ([]model.ResearchProject, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var list []model.ResearchProject
	for _, p := range r.projects {
		if userID == "" || p.UserID == userID {
			list = append(list, *p)
		}
	}
	return list, nil
}

func (r *InMemoryRepository) UpdateProjectStatus(id string, status model.ProjectStatus, completedAt *time.Time) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	p, ok := r.projects[id]
	if !ok {
		return ErrNotFound
	}
	p.Status = status
	p.CompletedAt = completedAt
	return nil
}

func (r *InMemoryRepository) CreateAgentExecution(exec *model.AgentExecution) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if exec.ID == "" {
		exec.ID = uuid.New().String()
	}
	r.executions[exec.ProjectID] = append(r.executions[exec.ProjectID], exec)
	return nil
}

func (r *InMemoryRepository) UpdateAgentExecution(exec *model.AgentExecution) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	list, ok := r.executions[exec.ProjectID]
	if !ok {
		return ErrNotFound
	}
	for i, e := range list {
		if e.ID == exec.ID || e.AgentName == exec.AgentName {
			list[i] = exec
			return nil
		}
	}
	r.executions[exec.ProjectID] = append(r.executions[exec.ProjectID], exec)
	return nil
}

func (r *InMemoryRepository) ListExecutionsByProjectID(projectID string) ([]model.AgentExecution, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list, ok := r.executions[projectID]
	if !ok {
		return []model.AgentExecution{}, nil
	}
	var res []model.AgentExecution
	for _, e := range list {
		res = append(res, *e)
	}
	return res, nil
}

func (r *InMemoryRepository) SaveReport(report *model.Report) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if report.ID == "" {
		report.ID = uuid.New().String()
	}
	if report.CreatedAt.IsZero() {
		report.CreatedAt = time.Now()
	}
	r.reports[report.ProjectID] = report
	return nil
}

func (r *InMemoryRepository) GetReportByProjectID(projectID string) (*model.Report, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	rep, ok := r.reports[projectID]
	if !ok {
		return nil, ErrNotFound
	}
	return rep, nil
}
